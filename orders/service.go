package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	common "github.com/sebastvin/commons"
	pb "github.com/sebastvin/commons/api"
	"github.com/sebastvin/omsv-orders/gateway"
	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	store        OrdersStore
	gateway      gateway.StockGateway
	openaiClient *openai.Client
}

func NewService(store OrdersStore, gateway gateway.StockGateway, openaiApiKey string) *service {
	client := openai.NewClient(
		option.WithAPIKey(openaiApiKey),
	)
	return &service{
		store:        store,
		gateway:      gateway,
		openaiClient: &client,
	}
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

	inStock, items, err := s.gateway.CheckIfItemIsInStock(ctx, p.CustomerID, p.Items)
	if err != nil {
		return nil, err
	}

	if !inStock {
		return items, common.ErrNoStock
	}

	return items, nil
}

func (s *service) GenerateAndSaveImages(ctx context.Context, orderID string) error {
	log.Printf("Starting image generation for order ID: %s", orderID)

	orderStruct, err := s.store.GetByID(ctx, orderID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("GenerateAndSaveImages: Order %s not found in store.", orderID)
			return fmt.Errorf("order not found: %w", err)
		}
		log.Printf("GenerateAndSaveImages: Failed to get order %s from store: %v", orderID, err)
		return fmt.Errorf("failed to get order from store: %w", err)
	}

	log.Printf("GenerateAndSaveImages: Retrieved order %s from store. Image present: %t, Items count: %d",
		orderID, orderStruct.Image != "", len(orderStruct.Items))

	if orderStruct.Image == "" || len(orderStruct.Items) == 0 {
		log.Printf("GenerateAndSaveImages: Order ID %s is missing image or items for generation in store. Image empty: %t, Items empty: %t",
			orderID, orderStruct.Image == "", len(orderStruct.Items) == 0)
		return nil
	}

	imageData, err := base64.StdEncoding.DecodeString(orderStruct.Image)
	if err != nil {
		log.Printf("GenerateAndSaveImages: Failed to decode base64 image for order %s: %v", orderID, err)
		return fmt.Errorf("failed to decode base64 image: %w", err)
	}

	log.Printf("GenerateAndSaveImages: Successfully decoded image base64 (%d bytes) for order %s.", len(imageData), orderID)

	generatedImagesBase64 := []string{}

	for _, item := range orderStruct.Items {
		styleName := item.StyleReference
		if styleName == "" {
			log.Printf("GenerateAndSaveImages: Item with empty ID found in order %s, skipping.", orderID)
			continue
		}

		prompt := fmt.Sprintf("Apply the style %s to this image.", styleName)

		log.Printf("GenerateAndSaveImages: Calling OpenAI images.edit for order %s, style '%s' with prompt: '%s'", orderID, styleName, prompt)

		imageReader := bytes.NewReader(imageData)
		imageFile := openai.File(imageReader, "test.png", "image/png")

		resp, err := s.openaiClient.Images.Edit(ctx, openai.ImageEditParams{
			Prompt: prompt,
			Image:  openai.ImageEditParamsImageUnion{OfFile: imageFile},
			N:      openai.Int(1),
			Model:  openai.ImageModelGPTImage1,
			Size:   openai.ImageEditParamsSize1024x1024,
		})

		if err != nil {
			log.Printf("GenerateAndSaveImages: OpenAI images.edit API call failed for order %s, style '%s': %v", orderID, styleName, err)
			return fmt.Errorf("openai image edit failed for style '%s': %w", styleName, err)
		}

		log.Printf("GenerateAndSaveImages: Received %d generated image(s) for order %s, style '%s'.", len(resp.Data), orderID, styleName)

		for _, resultItem := range resp.Data {
			if resultItem.B64JSON != "" {
				generatedImagesBase64 = append(generatedImagesBase64, resultItem.B64JSON)
			} else {
				log.Printf("GenerateAndSaveImages: OpenAI API did not return B64JSON for an image result for order %s, style '%s'. Response item: %+v", orderID, styleName, resultItem)
			}
		}
	}

	if len(generatedImagesBase64) > 0 {
		log.Printf("Service GenerateAndSaveImages: Generated %d images for order %s. Saving to DB.", len(generatedImagesBase64), orderID)
		update := &pb.Order{
			ID:            orderID,
			Status:        "ready",
			ResultsBase64: generatedImagesBase64,
		}
		err = s.store.Update(ctx, orderID, update)
		if err != nil {
			log.Printf("GenerateAndSaveImages: Failed to update order %s with generated images: %v", orderID, err)
			return fmt.Errorf("failed to update order with results: %w", err)
		}
	}

	log.Printf("GenerateAndSaveImages: Successfully updated order %s with generated images.", orderID)

	return nil
}
