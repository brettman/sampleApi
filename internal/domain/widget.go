package domain

import (
	"time"
)

// Widget - the primary domain object
type Widget struct {
	//ID           bson.ObjectId `bson:"_id,omitempty"`
	ID           string    `json:"id"`
	SerialNr     string    `json:"serialNr"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CurrentValue string    `json:"currentValue"`
	CreatedAt    time.Time `json:"createdAt"`
	LastUpdated  time.Time `json:"lastUpdated"`
}

// NewWidget - initialize a new widget with some values
func NewWidget() *Widget {
	return &Widget{CreatedAt: time.Now()}
}
