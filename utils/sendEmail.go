package utils

import (
	"commons/api/communication"
	"commons/broker"
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

func SendMessage(ctx context.Context, channel *amqp.Channel, to []string, cc []string, subject, message string) error {

	q, err := channel.QueueDeclare(broker.SendEmailEvent, true, false, false, false, nil)
	if err != nil {
		log.Printf("Broker error: %v", err)
		return err
	}
	body, err := json.Marshal(&communication.EmailRequest{
		To:      to,
		Cc:      cc,
		Subject: subject,
		Message: message,
	})
	if err != nil {
		log.Printf("Json marshal error: %v", err)
		return err
	}

	tr := otel.Tracer("amqp")
	amqpContext, messageSpan := tr.Start(ctx, fmt.Sprintf("AMQP - publish - %s", q.Name))
	defer messageSpan.End()

	headers := broker.InjectAMQPHeaders(amqpContext)

	channel.PublishWithContext(amqpContext, "", q.Name, false, false, amqp.Publishing{
		ContentType:  "application/json",
		Body:         body,
		DeliveryMode: amqp.Persistent,
		Headers:      headers,
	})

	return nil
}
