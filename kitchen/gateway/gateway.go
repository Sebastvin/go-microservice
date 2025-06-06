package gateway

import (
	"context"

	pb "github.com/sebastvin/commons/api"
)

type KitchenGateway interface {
	UpdateOrder(context.Context, *pb.Order) error
}
