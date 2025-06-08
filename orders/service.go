package main

import (
	"context"
	"log"

	common "github.com/sebastvin/commons"
	pb "github.com/sebastvin/commons/api"
	"github.com/sebastvin/omsv-orders/gateway"
)

type service struct {
	store   OrdersStore
	gateway gateway.StockGateway
}

func NewService(store OrdersStore, gateway gateway.StockGateway) *service {
	return &service{store, gateway}
}

func (s *service) GetOrder(ctx context.Context, p *pb.GetOrderRequest) (*pb.Order, error) {
	o, err := s.store.Get(ctx, p.OrderID, p.CustomerID)

	if err != nil {
		return nil, err
	}

	return o.ToProto(), nil
}

func (s *service) UpdateOrder(ctx context.Context, o *pb.Order) (*pb.Order, error) {
	err := s.store.Update(ctx, o.ID, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (s *service) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest, items []*pb.Item) (*pb.Order, error) {
	orderToSave := Order{
		CustomerID:  p.CustomerID,
		Status:      "pending",
		Items:       items,
		PaymentLink: "",
		Image:       p.Image,
	}

	log.Printf("Service CreateOrder: Attempting to save internal Order struct - CustomerID: %s, Image: %s, Status: %s",
		orderToSave.CustomerID, orderToSave.Image, orderToSave.Status)

	id, err := s.store.Create(ctx, orderToSave)
	if err != nil {
		log.Printf("Service CreateOrder: Failed to save order to store: %v", err)
		return nil, err
	}

	o := &pb.Order{
		ID:         id.Hex(),
		CustomerID: p.CustomerID,
		Status:     "pending",
		Items:      items,
		Image:      p.Image,
	}

	log.Printf("Service CreateOrder: Order created successfully - ID: %s, Image: %s", o.ID, o.Image)

	return o, nil
}

func (s *service) ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) ([]*pb.Item, error) {
	if len(p.Items) == 0 {
		return nil, common.ErrNoItems
	}

	mergedItems := mergeItemsQuantities(p.Items)

	inStock, items, err := s.gateway.CheckIfItemIsInStock(ctx, p.CustomerID, mergedItems)
	if err != nil {
		return nil, err
	}

	if !inStock {
		return items, common.ErrNoStock
	}

	return items, nil
}

func mergeItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
	merged := make([]*pb.ItemsWithQuantity, 0)

	for _, item := range items {
		found := false
		for _, finalItem := range merged {
			if finalItem.ID == item.ID {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}
		}

		if !found {
			merged = append(merged, item)
		}
	}

	return merged
}
