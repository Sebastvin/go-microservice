package main

import (
	"context"

	pb "github.com/sebastvin/commons/api"
	"github.com/sebastvin/omsv-payments/gateway"
	"github.com/sebastvin/omsv-payments/processor"
)

type service struct {
	processor processor.PaymentProcessor
	gateway   gateway.OrdersGateway
}

func NewService(processor processor.PaymentProcessor, gateway gateway.OrdersGateway) *service {
	return &service{processor, gateway}
}

func (s *service) CreatePayment(ctx context.Context, o *pb.Order) (string, error) {
	link, err := s.processor.CreatePaymentLink(o)
	if err != nil {
		return "", err
	}

	err = s.gateway.UpdateOrderAfterPaymentLink(ctx, o.ID, link)
	if err != nil {
		return "", err
	}

	return link, nil
}
