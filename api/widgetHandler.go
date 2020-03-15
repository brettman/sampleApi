package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/brettman/sampleApi/internal/services"
)

// Widgets - get all widgets
func Widgets(service *services.WidgetService) func(c *gin.Context) {
	return func(c *gin.Context) {
		widgets, _ := service.Widgets()
		c.JSON(http.StatusOK, widgets)
	}
}

// Widget - get one widget by id
func Widget(service *services.WidgetService) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		i, err := service.Widget(id)
		if err != nil {
			log.Panic(err)
		}
		c.JSON(http.StatusOK, i)
	}
}

// AddWidget - add a new widget to db
func AddWidget(service *services.WidgetService) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusTooEarly, "Not implemented yet")
	}
}

// UpdateWidget - update a new widget to db
func UpdateWidget(service *services.WidgetService) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusTooEarly, "Not implemented yet")
	}
}

// DeleteWidget - delete a widget
func DeleteWidget(service *services.WidgetService) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusTooEarly, "Not implemented yet")
	}
}

// WidgetsByCategory - get all the widgets with a particular category
func WidgetsByCategory(service *services.WidgetService) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusTooEarly, "Not implemented yet")
	}
}

// UpdateCurrentValue - This should be a long running, async process
func UpdateCurrentValue(service *services.WidgetService) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusTooEarly, "Not implemented yet")
	}
}
