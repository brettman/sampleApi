package services

import "github.com/brettman/sampleApi/internal/domain"

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
