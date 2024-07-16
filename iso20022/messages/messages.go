package messages

import (
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/messages"
)

var (
	messageConstructor         = map[string]messages.ConstructorFunc{}
	extendedMessageConstructor = make(map[string][]messages.ConstructorWithUrn)
)

func init() {
	messageConstructor, extendedMessageConstructor = messages.GetMessageConstructors()
}
