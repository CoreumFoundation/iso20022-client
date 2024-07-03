package server

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Send(c *gin.Context) {
	sendCh, ok := c.Get("sendChannel")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	sendChannel := sendCh.(chan<- []byte)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(c.Request.Body)

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.Status(http.StatusCreated)

	go func() {
		sendChannel <- body
	}()
}

func Receive(c *gin.Context) {
	recvCh, ok := c.Get("receiveChannel")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	receiveChannel := recvCh.(<-chan []byte)

	select {
	case message := <-receiveChannel:
		c.Data(http.StatusOK, "application/xml", message)
	case <-time.After(time.Second):
		c.Status(http.StatusNoContent)
	}
}
