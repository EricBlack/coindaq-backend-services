package workers

import (
	"log"
	"encoding/json"
)

type PhoneEvent struct {
	Phone      string `json:"phone"`
	OtcCode string `json:"otc_code"`
}

type PhoneSender struct{}

func init() {
	RegisterWorker("PhoneSender", &PhoneSender{})
}

func (e *PhoneSender) handle(msg []byte) error {
	event := new(PhoneEvent)
	json.Unmarshal(msg, event)

	log.Printf("Event %v, %v", event.Phone, event.OtcCode)

	return nil
}
