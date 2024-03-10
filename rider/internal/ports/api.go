package ports

import (
	"context"
	"github.com/AnwarSaginbai/taxi-service/rider/internal/application/core/domain"
)

//Контракт API

type APIPort interface {
	GetOrder(ctx context.Context, id int64) (domain.Order, error)
	PostOrder(ctx context.Context, order domain.Order) (domain.Order, error)
}
