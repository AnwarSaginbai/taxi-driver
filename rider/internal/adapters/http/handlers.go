package http

import (
	"context"
	"github.com/AnwarSaginbai/taxi-service/rider/internal/adapters"
	"github.com/AnwarSaginbai/taxi-service/rider/internal/application/core/domain"
	"log"
	"net/http"
	"time"
)

func (a *Adapter) PostOrderHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		PickUp  string `json:"pick_up"`
		DropOut string `json:"drop_out"`
		Price   int64  `json:"price"`
	}

	if err := adapters.ReadJSON(w, r, &input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output := &domain.Order{
		PickUp:    input.PickUp,
		DropOut:   input.DropOut,
		Price:     input.Price,
		CreatedAt: time.Now(),
	}

	data, err := a.service.PostOrder(context.Background(), *output)
	if err != nil {
		log.Fatalf("failled to insert: %v\n", err)
	}

	if err = adapters.WriteJSON(w, http.StatusOK, data, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
