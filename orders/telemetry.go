package main

import (
	"context"
	"fmt"

	pb "github.com/sebastvin/commons/api"
	"go.opentelemetry.io/otel"
	otelCodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type TelemetryMiddleware struct {
	next OrdersService
}

func NewTelemetryMiddleware(next OrdersService) OrdersService {
	return &TelemetryMiddleware{next}
}

func (s *TelemetryMiddleware) GetOrder(ctx context.Context, p *pb.GetOrderRequest) (*pb.Order, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(fmt.Sprintf("GetOrder: %v", p))
	return s.next.GetOrder(ctx, p)
}

func (s *TelemetryMiddleware) UpdateOrder(ctx context.Context, o *pb.Order) (*pb.Order, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(fmt.Sprintf("UpdateOrder: %v", o))
	return s.next.UpdateOrder(ctx, o)
}

func (s *TelemetryMiddleware) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest, items []*pb.Item) (*pb.Order, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(fmt.Sprintf("CreateOrder: %v, items: %v", p, items))
	return s.next.CreateOrder(ctx, p, items)
}

func (s *TelemetryMiddleware) ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) ([]*pb.Item, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(fmt.Sprintf("ValidateOrder: %v", p))
	return s.next.ValidateOrder(ctx, p)
}

func (s *TelemetryMiddleware) GenerateAndSaveImages(ctx context.Context, orderID string) (err error) {

	tr := otel.Tracer("orders-service")
	ctx, span := tr.Start(ctx, fmt.Sprintf("GenerateAndSaveImages for order ID: %s", orderID))
	defer span.End()

	span.AddEvent(fmt.Sprintf("Starting image generation for order ID: %s", orderID))

	err = s.next.GenerateAndSaveImages(ctx, orderID)

	if err != nil {
		span.SetStatus(otelCodes.Error, "Image generation failed")
		span.RecordError(err)
	} else {
		span.SetStatus(otelCodes.Ok, "Image generation completed")
	}

	return err
}
