package ports

import (
	"context"
	"github.com/AnwarSaginbai/taxi-service/rider/internal/application/core/domain"
)

type DBPort interface {
	Insert(ctx context.Context, order domain.Order) (domain.Order, error)
	Get(ctx context.Context, id int64) (domain.Order, error)
}
