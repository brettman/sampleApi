package domain

import "time"

// Widget - the primary domain object
type Widget struct {
	ID           string
	Name         string
	Description  string
	CurrentValue string
	TimeStamp    time.Time
}

// NewWidget - initialize a new widget with some values
func NewWidget() *Widget {
	return &Widget{TimeStamp: time.Now()}
}
