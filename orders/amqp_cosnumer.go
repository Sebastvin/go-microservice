package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	pb "github.com/sebastvin/commons/api"
	"github.com/sebastvin/commons/broker"
	"go.opentelemetry.io/otel"
)

type consumer struct {
	service OrdersService
}

func NewConsumer(service OrdersService) *consumer {
	return &consumer{service}
}

func (c *consumer) Listen(ch *amqp.Channel) {
	q, err := ch.QueueDeclare("", true, false, true, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(q.Name, "", broker.OrderPaidEvent, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received message: %s", d.Body)

			// Extract the headers
			ctx := broker.ExtractAMQPHeader(context.Background(), d.Headers)

			tr := otel.Tracer("amqp")
			_, messageSpan := tr.Start(ctx, fmt.Sprintf("AMQP - consume - %s", q.Name))

			o := &pb.Order{}
			if err := json.Unmarshal(d.Body, o); err != nil {
				d.Nack(false, false)
				log.Printf("failed to unmarshal order: %v", err)
				continue
			}

			_, err := c.service.UpdateOrder(context.Background(), o)
			if err != nil {
				log.Printf("failed to update order: %v", err)

				if err := broker.HandleRetry(ch, &d); err != nil {
					log.Printf("Error handling retry: %v", err)
				}
				continue
			}

			messageSpan.AddEvent("order.status.updated.paid")
			messageSpan.End()
			log.Printf("AMQP Consumer: Order %s status updated to paid.", o.ID)
			log.Printf("AMQP Consumer: Triggering image generation for order %s", o.ID)

			err = c.service.GenerateAndSaveImages(ctx, o.ID)
			if err != nil {
				log.Printf("AMQP Consumer: Failed to generate and save images for order %s: %v", o.ID, err)
			} else {
				log.Printf("AMQP Consumer: Image generation process initiated (or completed) for order %s", o.ID)
				messageSpan.AddEvent("image.generation.completed")
			}

			log.Println("AMQP Consumer: Order processing complete. Acknowledging message.")
			d.Ack(false)
		}
	}()

	<-forever
}
