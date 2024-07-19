package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CoreumFoundation/iso20022-client/iso20022/processes"
)

type Handler struct {
	Parser       processes.Parser
	MessageQueue processes.MessageQueue
}

func (h *Handler) Send(c *gin.Context) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(c.Request.Body)

	message, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	messageId, _, err := h.Parser.ExtractMetadataFromIsoMessage(message)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// TODO: Check for duplicate messages by ID
	fmt.Printf("Sending message with ID : %s\n", messageId)

	c.Status(http.StatusCreated)

	go h.MessageQueue.PushToSend(messageId, message)
}

func (h *Handler) Receive(c *gin.Context) {
	message, ok := h.MessageQueue.PopFromReceive()
	if ok {
		c.Data(http.StatusOK, "application/xml", message)
	} else {
		c.Status(http.StatusNoContent)
	}
}
