package main

import (
	dapi "github.com/brettman/sampleApi/api"
	"github.com/brettman/sampleApi/internal/services"
	"github.com/gin-gonic/gin"
)

// investigate DI from google:
//  https://github.com/google/wire
func main() {
	router := gin.Default()
	widgetService := services.WidgetService{
		DataContext: "for now this datacontext is just a string",
	}

	api := router.Group("/api")
	{
		api.GET("/widget", dapi.Widget(&widgetService))
	}

	router.Run()
}
