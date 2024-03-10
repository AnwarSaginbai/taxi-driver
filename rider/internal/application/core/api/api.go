package api

import (
	"context"
	"fmt"
	"github.com/AnwarSaginbai/taxi-service/rider/internal/application/core/domain"
	"github.com/AnwarSaginbai/taxi-service/rider/internal/ports"
)

type Service struct {
	db ports.DBPort
}

func NewService(db ports.DBPort) *Service {
	return &Service{db: db}
}

func (svc *Service) GetOrder(ctx context.Context, id int64) (domain.Order, error) {
	return domain.Order{}, nil
}
func (svc *Service) PostOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	order, err := svc.db.Insert(ctx, order)
	if err != nil {
		return domain.Order{}, fmt.Errorf("failed to create order: %w\n", err)
	}
	return order, nil
}
