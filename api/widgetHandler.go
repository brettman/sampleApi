package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/brettman/sampleApi/internal/services"
)

// Widget - get all widgets
func Widget(service *services.WidgetService) func(c *gin.Context) {
	return func(c *gin.Context) {
		//widget := domain.NewWidget("idd", "widgetName", "widget desc")
		widgets, _ := service.Widgets()
		c.JSON(http.StatusOK, widgets)
	}
}
