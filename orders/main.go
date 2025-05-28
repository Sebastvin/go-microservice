package main

import (
	"context"
	"log"
	"net"

	"github.com/sebastvin/commons"
	"google.golang.org/grpc"
)

var (
	grcpAddr = common.EnvString("GRCP_ADDR", "localhost:2000")
)

func main() {

	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grcpAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)
	NewGRPCHandler(grpcServer, svc)

	svc.CreateOrder(context.Background())

	log.Println("GRPC Server Started at ", grcpAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
