package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/CoreumFoundation/iso20022-client/iso20022/processes"
)

type Server struct {
	httpServer http.Server
}

func CreateHandlers(parser processes.Parser, sendChannel chan<- []byte, receiveChannel <-chan []byte) http.Handler {
	r := gin.Default()
	r.Use(InjectDependencies(parser, sendChannel, receiveChannel))
	r.Use(CORSMiddleware())

	v1 := r.Group("/v1")

	v1.POST("/send", Send)
	v1.GET("/receive", Receive)
	return r
}

func New(addr string, handler http.Handler) *Server {
	return &Server{
		httpServer: http.Server{
			Addr:    addr,
			Handler: handler,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		},
	}
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		<-ctx.Done()

		// Create a 5s timeout context or waiting for app to shut down after 5 seconds
		ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelTimeout()

		s.httpServer.SetKeepAlivesEnabled(false)
		_ = s.httpServer.Shutdown(ctxTimeout)
	}()

	err := s.httpServer.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
