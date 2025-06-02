package processor

import (
	pb "github.com/sebastvin/commons/api"
)

type PaymentProcessor interface {
	CreatePaymentLink(*pb.Order) (string, error)
}
