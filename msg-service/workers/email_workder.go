package workers

import (
	"encoding/json"
	"log"
)

type EmailEvent struct {
	To      string `json:"to"`
	OtcCode string `json:"otcCode"`
}

type EmailSender struct{}

func init() {
	RegisterWorker("EmailSender", &EmailSender{})
}

func (e *EmailSender) handle(msg []byte) error {
	log.Printf("Msg: %v", string(msg))
	event := new(EmailEvent)
	json.Unmarshal(msg, event)
	log.Printf("Event %v, %v", event.To, event.OtcCode)

	return nil
	/*
		switch event.Type {
		case 1:
			fmt.Printf("Type1: To=%s, Template=%s, Message=%s", event.To, event.Template, event.Body)
			break
		case 2:
			fmt.Printf("Type2: To=%s, Template=%s, Message=%s", event.To, event.Template, event.Body)
			break
		case 3:
			fmt.Printf("Type3: To=%s, Template=%s, Message=%s", event.To, event.Template, event.Body)
			break
		default:
			fmt.Printf("Invalid Queue Message.")
		}
	*/
}
