package main

import (
	"example/rest-api/configs"
	"example/rest-api/internal/app/rest_api/database"
	"example/rest-api/internal/app/rest_api/repositories"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	_ "github.com/lib/pq"

	"github.com/rs/zerolog/log"
	 "example/rest-api/internal/app/rest_api/services"
 "example/rest-api/internal/app/rest_api/handlers"
	routes "example/rest-api/api/router"
	serve "example/rest-api/api/server"
	"time"

)

func main(){
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	config:=configs.NewConfig()
	
	client,err:=database.NewSQLClient(database.Config{
		DBDriver: config.Database.DatabaseDriver,
		DBSource: config.Database.DatabaseSource,
		MaxOpenConns: 25,
		MaxIdleConns: 25,
		ConnMaxIdleTime: 5*time.Minute,
		ConnectionTimeout: 5*time.Second,
	})

	if err!=nil{
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	defer func(){
		if err:=client.DB.Close(); err!=nil{
			log.Error().Msg("Failed to close database connection")
		}	
	}()

	userRepo:=repositories.NewUserRepository(client.DB)

	userService:=services.NewUserService(userRepo)

	userHandlers:=handlers.NewUserHandler(userService)

	cors:=config.CorsNew()

	router:=gin.Default()
	router.Use(cors)

	routes.RegisterPublicEndPoints(router,userHandlers)
	
	server:=serve.NewServer(log.Logger,router,config)
	server.Serve()

}


