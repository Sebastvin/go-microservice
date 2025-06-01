package main

import (
	"context"

	pb "github.com/sebastvin/commons/api"
)

type PaymentService interface {
	CreatePayment(context.Context, *pb.Order) (string, error)
}
