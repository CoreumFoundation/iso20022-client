package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CoreumFoundation/iso20022-client/iso20022/processes"
)

func Send(c *gin.Context) {
	mq, ok := c.Get("messageQueue")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	messageQueue := mq.(processes.MessageQueue)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(c.Request.Body)

	message, err := io.ReadAll(c.Request.Body)
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

	messageId, _, err := parser.ExtractMetadataFromIsoMessage(message)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// TODO: Check for duplicate messages by ID
	fmt.Printf("Sending message with ID : %s\n", messageId)

	c.Status(http.StatusCreated)

	go messageQueue.PushToSend(messageId, message)
}

func Receive(c *gin.Context) {
	mq, ok := c.Get("messageQueue")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	messageQueue := mq.(processes.MessageQueue)

	message, ok := messageQueue.PopFromReceive()
	if ok {
		c.Data(http.StatusOK, "application/xml", message)
	} else {
		c.Status(http.StatusNoContent)
	}
}
