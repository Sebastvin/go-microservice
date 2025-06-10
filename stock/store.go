package main

import (
	"context"
	"fmt"

	pb "github.com/sebastvin/commons/api"
)

type StockItem struct {
	ID      string
	Name    string
	PriceID string
}

type Store struct {
	stock map[string]*StockItem
}

func NewStore() *Store {
	return &Store{
		stock: map[string]*StockItem{
			"1": {
				ID:      "1",
				Name:    "Onion",
				PriceID: "price_1RVZ3hClTXDUG291P1wmsO9h",
			},
			"2": {
				ID:      "2",
				Name:    "Pepper",
				PriceID: "price_1RX1NUClTXDUG29157neqin4",
			},
		},
	}
}

func (s *Store) GetItem(ctx context.Context, id string) (*pb.Item, error) {
	for _, item := range s.stock {
		if item.ID == id {
			return &pb.Item{
				ID:      item.ID,
				Name:    item.Name,
				PriceID: item.PriceID,
			}, nil
		}
	}

	return nil, fmt.Errorf("item not found")
}

func (s *Store) GetItems(ctx context.Context, ids []string) ([]*pb.Item, error) {
	var res []*pb.Item
	for _, id := range ids {
		if i, ok := s.stock[id]; ok {
			res = append(res, &pb.Item{
				ID:      i.ID,
				Name:    i.Name,
				PriceID: i.PriceID,
			})
		}
	}

	return res, nil
}
