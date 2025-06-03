package main

import (
	"context"
	"testing"

	"github.com/sebastvin/commons/api"
	"github.com/sebastvin/omsv-payments/processor/inmem"
)

func TestService(t *testing.T) {
	processor := inmem.NewInmem()
	svc := NewService(processor)

	t.Run("should create a payment link", func(t *testing.T) {
		link, err := svc.CreatePayment(context.Background(), &api.Order{})
		if err != nil {
			t.Errorf("CreatePayment() error = %v want nil", err)
		}

		if link == "" {
			t.Errorf("CreatePayment() link in empty")
		}
	})
}
