package domain

import "time"

type Order struct {
	ID          int64     `json:"id"`
	PickUp      string    `json:"pick_up"`
	DropOut     string    `json:"drop_out"`
	Price       int64     `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}
