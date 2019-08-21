package service

import "log"

type PhoneService struct {
	Service
}

func init() {
	RegisterServiceMapping("PhoneService", PhoneService{})
}

func (w *PhoneService) RegisterConsumeEvent(body string){
	log.Printf("Begin handle phone message:")
	log.Printf("Send Message to: %s", body)
	log.Printf("Complete handler phone message")
}

func (w *PhoneService) UninstallConsumeEvent(){
}
