package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/brettman/sampleApi/internal/domain"
	"github.com/brettman/sampleApi/internal/services"
	"github.com/brettman/sampleApi/internal/utils"
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
		log.Printf("id is: %s", id)
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
		var w domain.Widget
		err := c.ShouldBindJSON(&w)
		if err != nil {
			utils.Log(err)
		}

		if w.ID == "" {
			log.Printf("the id for the widget is empty")
			c.JSON(http.StatusBadRequest, w)
			return
		}

		widget, err1 := service.AddWidget(w)
		utils.Log(err1)

		c.JSON(http.StatusAccepted, widget)
	}
}

// UpdateWidget - update a new widget to db
func UpdateWidget(service *services.WidgetService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var w domain.Widget
		err := c.ShouldBindJSON(&w)
		utils.Log(err)

		widget, err1 := service.UpdateWidget(w)
		utils.Log(err1)

		c.JSON(http.StatusOK, widget)
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
