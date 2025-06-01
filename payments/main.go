package main

import (
	"context"
	"log"
	"net"
	"time"

	common "github.com/sebastvin/commons"
	"github.com/sebastvin/commons/broker"
	"github.com/sebastvin/commons/discovery"
	"github.com/sebastvin/commons/discovery/consul"
	"google.golang.org/grpc"
)

var (
	serviceName = "payment"
	amqpUser    = common.EnvString("RABBITMQ_USER", "guest")
	amqpPass    = common.EnvString("RABBITMQ_PASS", "guest")
	amqpHost    = common.EnvString("RABBITMQ_HOST", "localhost")
	amqpPort    = common.EnvString("RABBITMQ_PORT", "5672")
	grcpAddr    = common.EnvString("GRCP_ADDR", "localhost:2001")
	consulAddr  = common.EnvString("CONSUL_ADDR", "localhost:8500")
)

func main() {
	registry, err := consul.NewRegistry(consulAddr, serviceName)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)

	if err := registry.Register(ctx, instanceID, serviceName, grcpAddr); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				log.Fatal("failed to health check")
			}
			time.Sleep(time.Second * 1)
		}
	}()

	defer registry.Deregister(ctx, instanceID, serviceName)

	ch, close := broker.Connect(amqpUser, amqpPass, amqpHost, amqpPort)

	defer func() {
		close()
		ch.Close()
	}()

	svc := NewService()

	amqpConsumer := NewConsumer(svc)
	go amqpConsumer.Listen(ch)

	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grcpAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer l.Close()

	log.Println("GRPC Server Started at ", grcpAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
