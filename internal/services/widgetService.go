package services

import (
	"errors"

	"github.com/brettman/sampleApi/internal/domain"
	"gopkg.in/mgo.v2"
)

// WidgetService - implementation of the IWidgetService interface
type WidgetService struct {
	DataContext *mgo.Database
}

// Widgets - get all widgets from whatever sources... currently only one database
func (w *WidgetService) Widgets() ([]domain.Widget, error) {
	widgets := []domain.Widget{}
	err := errors.New("not implemented")
	return widgets, err
}

// Widget - get a widget by id
func (w *WidgetService) Widget(id string) (*domain.Widget, error) {

	i := domain.NewWidget()
	i.ID = id
	i.Name = "manually created widget"
	i.Description = "this manual shit has to stop soon, dude"

	return i, nil
}
