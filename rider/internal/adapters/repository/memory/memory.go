package memory

import (
	"context"
	"errors"
	"github.com/AnwarSaginbai/taxi-service/rider/internal/application/core/domain"
	"log"
	"sync"
	"time"
)

type Repository struct {
	sync.RWMutex
	data map[int64]*domain.Order
}

func New() *Repository {
	return &Repository{
		data: make(map[int64]*domain.Order),
	}
}

func (r *Repository) Get(ctx context.Context, id int64) (domain.Order, error) {
	r.RLock()
	defer r.RUnlock()

	data, ok := r.data[id]
	if !ok {
		return domain.Order{}, errors.New("order not found")
	}
	return *data, nil
}

func (r *Repository) Insert(ctx context.Context, order domain.Order) (domain.Order, error) {
	r.Lock()
	defer r.Unlock()

	id := generateID()

	order.ID = id

	r.data[id] = &order
	log.Println(order)
	return order, nil
}

func generateID() int64 {
	// Реализация зависит от требований вашего приложения.
	// Можно использовать UUID, случайные числа или другие методы генерации.
	// Возвращаем временную метку как пример.
	return time.Now().UnixNano()
}
