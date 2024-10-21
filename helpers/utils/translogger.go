package utils

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type TransLogger struct {
	ClientToWebhook  string
	WebhookToBackend string
	BackendToWebhook string
	WebhookToClient  string
	Status           string
	ErrDesc          string
}

func (tl *TransLogger) SaveLog() {
	text := fmt.Sprintf("Client->Webhook:%s, Webhook->Backend:%s, Webhook<-Backend:%s, Client<-Webhook:%s, Status:%s, ErrDesc:%s",
		tl.ClientToWebhook, tl.WebhookToBackend, tl.BackendToWebhook, tl.WebhookToClient, tl.Status, tl.ErrDesc)
	log.Info(text)
}
