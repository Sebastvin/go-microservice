package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/sebastvin/commons/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DbName   = "orders"
	CollName = "orders"
)

type store struct {
	db *mongo.Client
}

func NewStore(db *mongo.Client) *store {
	return &store{db}
}

func (s *store) Create(ctx context.Context, o Order) (primitive.ObjectID, error) {
	col := s.db.Database(DbName).Collection(CollName)

	newOrder, err := col.InsertOne(ctx, o)

	id := newOrder.InsertedID.(primitive.ObjectID)
	return id, err
}

func (s *store) Get(ctx context.Context, id, customerID string) (*Order, error) {
	col := s.db.Database(DbName).Collection(CollName)

	oID, _ := primitive.ObjectIDFromHex(id)

	var o Order
	err := col.FindOne(ctx, bson.M{
		"_id":        oID,
		"customerID": customerID,
	}).Decode(&o)

	return &o, err
}

func (s *store) Update(ctx context.Context, id string, newOrder *pb.Order) error {
	col := s.db.Database(DbName).Collection(CollName)

	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Store Update: Error converting ID %s to ObjectID: %v", id, err)
		return fmt.Errorf("invalid order ID: %w", err)
	}

	log.Printf("Store Update: Attempting to update order with ObjectID: %v. New fields: Status=%s, PaymentLink=%s, Image=%s, ResultsBase64Count=%d",
		oID, newOrder.Status, newOrder.PaymentLink, newOrder.Image, len(newOrder.ResultsBase64))

	res, err := col.UpdateOne(ctx,
		bson.M{"_id": oID},
		bson.M{"$set": bson.M{
			"paymentLink":   newOrder.PaymentLink,
			"status":        newOrder.Status,
			"resultsBase64": newOrder.ResultsBase64,
		}},
	)

	if err != nil {
		log.Printf("Store Update: Error during MongoDB UpdateOne for ID %v: %v", oID, err)
		return fmt.Errorf("failed to update order in db: %w", err)
	}

	log.Printf("Store Update: MongoDB UpdateOne result for ID %v: MatchedCount=%d, ModifiedCount=%d, UpsertedCount=%d",
		oID, res.MatchedCount, res.ModifiedCount, res.UpsertedCount)

	return err
}

func (s *store) GetByID(ctx context.Context, id string) (*Order, error) {
	col := s.db.Database(DbName).Collection(CollName)

	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Store GetByID: Error converting ID %s to ObjectID: %v", id, err)
		return nil, fmt.Errorf("invalid order ID: %w", err)
	}

	var o Order
	err = col.FindOne(ctx, bson.M{"_id": oID}).Decode(&o)

	if err != nil {
		log.Printf("Store GetByID: Failed to find order with ID %v: %v", oID, err)
		return nil, fmt.Errorf("failed to find order by ID: %w", err)
	}

	log.Printf("Store GetByID: Successfully retrieved order with ID %v. Order struct: %+v", oID, o)

	return &o, nil
}
