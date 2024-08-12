package server

import "github.com/CoreumFoundation/iso20022-client/iso20022/queue"

type MessageStatusResponse struct {
	MessageID      string       `json:"message_id"`
	DeliveryStatus queue.Status `json:"delivery_status"`
}
