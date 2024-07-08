package server

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/CoreumFoundation/iso20022-client/iso20022/processes"
)

func Send(c *gin.Context) {
	sendCh, ok := c.Get("sendChannel")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	sendChannel := sendCh.(chan<- []byte)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(c.Request.Body)

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	p, ok := c.Get("parser")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	parser := p.(processes.Parser)

	_, err = parser.ExtractIdentificationFromIsoMessage(body)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
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
		return
	}
	receiveChannel := recvCh.(<-chan []byte)

	wait, err := time.ParseDuration(c.Query("wait"))
	if err != nil {
		wait = time.Second
	}

	select {
	case message := <-receiveChannel:
		c.Data(http.StatusOK, "application/xml", message)
	case <-time.After(wait):
		c.Status(http.StatusNoContent)
	}
}
