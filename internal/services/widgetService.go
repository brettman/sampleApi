package services

import (
	"log"

	"github.com/brettman/sampleApi/internal/utils"

	"github.com/brettman/sampleApi/internal/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// WidgetService - implementation of the IWidgetService interface
type WidgetService struct {
	DataContext *mgo.Database
}

// Widgets - get all widgets from whatever sources... currently only one database
func (w *WidgetService) Widgets() ([]domain.Widget, error) {
	widgets := []domain.Widget{}
	err := w.DataContext.C("widgets").Find(bson.M{}).All(&widgets)
	return widgets, err
}

// Widget - get a widget by id
func (w *WidgetService) Widget(id string) (*domain.Widget, error) {

	log.Printf("The id passed to the widget service is: %s", id)
	widget := domain.Widget{}
	err := w.DataContext.C("widgets").Find(bson.M{"id": id}).One(&widget)
	if err != nil {
		utils.Log(err)
	}
	return &widget, nil
}

// AddWidget - add a widget to the db
func (w *WidgetService) AddWidget(widget domain.Widget) (domain.Widget, error) {

	err := w.DataContext.C("widgets").Insert(&widget)
	if err != nil {
		utils.Log(err)
	}

	return widget, err
}

// UpdateWidget - add a widget to the db
func (w *WidgetService) UpdateWidget(widget domain.Widget) (domain.Widget, error) {

	//err := w.DataContext.C("widgets").UpdateId(widget.ID, &widget)
	selector := bson.M{"id": widget.ID}
	_, err := w.DataContext.C("widgets").Upsert(selector, &widget)
	if err != nil {
		utils.Log(err)
	}

	return widget, err
}
