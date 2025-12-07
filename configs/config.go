package configs

import (
	"example/rest-api/internal/app/rest_api/constants"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"

	"os"
)

type Config struct{
	Server serverConfig
	Database databaseConfig
}

type serverConfig struct {
	Address string
}

type databaseConfig struct{
	DatabaseDriver string
	DatabaseSource string	
}

func NewConfig() *Config{	
	

	c:=&Config{
		Server: serverConfig{
			Address:GetEnvOrPanic(constants.EnvKeys.ServerAddress),
		},
		Database: databaseConfig{
			DatabaseDriver:GetEnvOrPanic(constants.EnvKeys.DBDriver),
			DatabaseSource:GetEnvOrPanic(constants.EnvKeys.DBSource),
		},
	}
	return c
}

func GetEnvOrPanic(key string) string {
	value:=os.Getenv(key)
	if value==""{
		panic(fmt.Sprintf("Environment variable %s not set",key))
	}
	return value
}


func(conf *Config) CorsNew() gin.HandlerFunc{
	allowedOrigin:=GetEnvOrPanic(constants.EnvKeys.CorsAllowedOrigin)

	return cors.New(cors.Config{
		AllowMethods: []string{http.MethodGet,http.MethodPost,http.MethodPut,http.MethodDelete},
		AllowHeaders: []string{constants.Headers.Origin},
		ExposeHeaders: []string{constants.Headers.ContentLength},
		AllowCredentials:true,
		AllowOriginFunc: func(origin string) bool{
			return origin==allowedOrigin
		},
		MaxAge: constants.MaxAge,

	})
}
