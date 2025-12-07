package serve

import (
	"example/rest-api/configs"
	"net/http"
	"syscall"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"errors"
	"context"
	"os/signal"
	"time"
)

type Server struct{
	l zerolog.Logger
	router *gin.Engine
	config *configs.Config
}

func NewServer(l zerolog.Logger, router *gin.Engine, config *configs.Config) *Server{
	return &Server{
		l:l,
		router:router,
		config:config,
	}
}

func (s *Server) Serve() {
 srv := &http.Server{
  Addr:    s.config.Server.Address,
  Handler: s.router.Handler(),
 }

 go func() {
  // service connections
  if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
   s.l.Fatal().Err(err).Msg("listen")
  }
 }()

 quit := make(chan os.Signal, 1)

 signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
 <-quit
 s.l.Info().Msg("Shutting down server...")

 ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
 defer cancel()
 if err := srv.Shutdown(ctx); err != nil {
  s.l.Fatal().Err(err).Msg("Server Shutdown")
 }

 <-ctx.Done()
 s.l.Info().Msg("Server shutdown timeout of 30 seconds")
 s.l.Info().Msg("Server exiting")
}