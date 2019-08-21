package service

import "log"

type EmailService struct {
	Service
}

func init(){
	RegisterServiceMapping("EmailService", EmailService{})
}

func (w *EmailService) RegisterConsumeEvent(body string){
	log.Printf("Begin handle message:")
	log.Printf("Send Email to: %s", body)
	log.Printf("Complete handler message")
}

func (w *EmailService) UninstallConsumeEvent(){
}
