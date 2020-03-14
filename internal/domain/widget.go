package domain

// Widget - the primary domain object
type Widget struct {
	ID           string
	Name         string
	Description  string
	CurrentValue string
}

// WidgetService - defines the basic Widget operations
type WidgetService interface {

	// Basic CRUD
	Widgets() ([]Widget, error)
	Widget(id string) (Widget, error)
	AddWidget(widget Widget) (Widget, error)
	UpdateWidget(widget Widget) (Widget, error)
	DeleteWidget(id string) error

	// just and excuse to do something else with data layer
	WidgetsByCategory(category string) ([]Widget, error)

	// some kind of a long running calculation
	UpdateCurrentValue(id string) (Widget, error)
}
