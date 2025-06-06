package main

import (
	"context"
	"fmt"

	pb "github.com/sebastvin/commons/api"
)

type Store struct {
	stock map[string]*pb.Item
}

func NewStore() *Store {
	return &Store{
		stock: map[string]*pb.Item{
			"2": {
				ID:       "2",
				Name:     "Onion",
				PriceID:  "price_1RVZ3hClTXDUG291P1wmsO9h",
				Quantity: 10,
			},
			"1": {
				ID:       "1",
				Name:     "Pepper",
				PriceID:  "price_1RX1NUClTXDUG29157neqin4",
				Quantity: 20,
			},
		},
	}
}

func (s *Store) GetItem(ctx context.Context, id string) (*pb.Item, error) {
	for _, item := range s.stock {
		if item.ID == id {
			return item, nil
		}
	}

	return nil, fmt.Errorf("item not found")
}

func (s *Store) GetItems(ctx context.Context, ids []string) ([]*pb.Item, error) {
	var res []*pb.Item
	for _, id := range ids {
		if i, ok := s.stock[id]; ok {
			res = append(res, i)
		}
	}

	return res, nil
}
