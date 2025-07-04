package main

import (
	"context"

	pb "github.com/sebastvin/commons/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrdersService interface {
	CreateOrder(context.Context, *pb.CreateOrderRequest, []*pb.Item) (*pb.Order, error)
	ValidateOrder(context.Context, *pb.CreateOrderRequest) ([]*pb.Item, error)
	GetOrder(context.Context, *pb.GetOrderRequest) (*pb.Order, error)
	UpdateOrder(context.Context, *pb.Order) (*pb.Order, error)
	GenerateAndSaveImages(ctx context.Context, orderID string) error
}

type OrdersStore interface {
	Create(context.Context, Order) (primitive.ObjectID, error)
	Get(ctx context.Context, id, customerID string) (*Order, error)
	Update(ctx context.Context, id string, o *pb.Order) error
	GetByID(ctx context.Context, id string) (*Order, error)
}

type Order struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID    string             `bson:"customerID,omitempty"`
	Status        string             `bson:"status,omitempty"`
	PaymentLink   string             `bson:"paymentLink,omitempty"`
	Items         []*pb.Item         `bson:"items,omitempty"`
	Image         string             `bson:"image,omitempty"`
	ResultsBase64 []string           `bson:"resultsBase64,omitempty"`
}

func (o *Order) ToProto() *pb.Order {
	return &pb.Order{
		ID:            o.ID.Hex(),
		CustomerID:    o.CustomerID,
		Status:        o.Status,
		PaymentLink:   o.PaymentLink,
		Image:         o.Image,
		ResultsBase64: o.ResultsBase64,
	}
}
