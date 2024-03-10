package http

import (
	"fmt"
	"github.com/AnwarSaginbai/taxi-service/rider/internal/application/core/api"
	"log"
	"net/http"
	"time"
)

type Adapter struct {
	service *api.Service
}

func NewAdapter(service *api.Service) *Adapter {
	return &Adapter{service: service}
}
func (a *Adapter) Run(addr int) {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", addr),
		Handler:           mux,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       10 * time.Second,
	}
	mux.HandleFunc("POST /order", a.PostOrderHandler)

	log.Fatal(srv.ListenAndServe())
}
