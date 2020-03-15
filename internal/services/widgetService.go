package services

import (
	"errors"

	"github.com/brettman/sampleApi/internal/domain"
)

// IWidgetService - defines the basic Widget operations
type IWidgetService interface {

	// Basic CRUD
	Widgets() ([]domain.Widget, error)
	Widget(id string) (domain.Widget, error)
	AddWidget(widget domain.Widget) (domain.Widget, error)
	UpdateWidget(widget domain.Widget) (domain.Widget, error)
	DeleteWidget(id string) error

	// just and excuse to do something else with data layer
	WidgetsByCategory(category string) ([]domain.Widget, error)

	// some kind of a long running calculation
	UpdateCurrentValue(id string) (domain.Widget, error)
}

// WidgetService - implementation of the IWidgetService interface
type WidgetService struct {
	DataContext string
}

// Widgets - get all widgets from whatever sources... currently only one database
func (w WidgetService) Widgets() ([]domain.Widget, error) {
	widgets := []domain.Widget{}
	err := errors.New("not implemented")
	return widgets, err
}
