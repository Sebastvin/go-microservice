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
				Name:    "GTA",
				PriceID: "price_1RcDQsClTXDUG291KHTKj0RV",
			},
			"2": {
				ID:      "2",
				Name:    "Retro",
				PriceID: "price_1RX1NUClTXDUG29157neqin4",
			},
			"3": {
				ID:      "3",
				Name:    "Anime",
				PriceID: "price_1RVZ3hClTXDUG291P1wmsO9h",
			},
			"4": {
				ID:      "4",
				Name:    "Pixel Art",
				PriceID: "price_1RcDWaClTXDUG291ifiemCK9",
			},
			"5": {
				ID:      "5",
				Name:    "Watercolor",
				PriceID: "price_1RcDX5ClTXDUG291rrlWjZkH",
			},
			"6": {
				ID:      "6",
				Name:    "Studio Ghibli",
				PriceID: "price_1RcDXWClTXDUG291dFwFIOOp",
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

func (s *Store) Get(id string) (*StockItem, bool) {
	item, exists := s.stock[id]
	return item, exists
}
