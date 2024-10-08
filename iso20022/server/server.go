package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/processes"
)

type Server struct {
	messageQueue processes.MessageQueue
	httpServer   http.Server
}

func createHandlers(logger logger.Logger, parser processes.Parser, messageQueue processes.MessageQueue) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	h := Handler{
		Logger:       logger,
		Parser:       parser,
		MessageQueue: messageQueue,
	}

	r.GET("", h.Doc)
	r.GET("/swagger.json", h.Swagger)

	v1 := r.Group("/v1")

	v1.POST("/send", h.Send)
	v1.GET("/receive", h.Receive)
	v1.GET("/status/:id", h.Status)
	return r
}

func New(logger logger.Logger, parser processes.Parser, messageQueue processes.MessageQueue, addr string) *Server {
	return &Server{
		messageQueue: messageQueue,
		httpServer: http.Server{
			Addr:    addr,
			Handler: createHandlers(logger, parser, messageQueue),
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
		s.messageQueue.Close()
	}()

	err := s.httpServer.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
