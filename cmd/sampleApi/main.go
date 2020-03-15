package main

import (
	"log"

	dapi "github.com/brettman/sampleApi/api"
	"github.com/brettman/sampleApi/internal/services"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

// todo: investigate DI from google:
//  https://github.com/google/wire
func main() {
	router := gin.Default()

	widgetService := services.WidgetService{
		DataContext: getDbSession(),
	}

	api := router.Group("/api")
	{
		// find
		api.GET("/widgets", dapi.Widgets(&widgetService))
		api.GET("/widget/:id", dapi.Widget(&widgetService))
		api.GET("/widgets/:category", dapi.WidgetsByCategory(&widgetService))

		// create
		api.POST("/widget", dapi.AddWidget(&widgetService))

		// update
		api.PATCH("/widget", dapi.UpdateWidget(&widgetService))
		api.PATCH("/widget/:id/current", dapi.UpdateCurrentValue(&widgetService))

		// remove
		api.DELETE("/widget/:id", dapi.DeleteWidget(&widgetService))
	}

	router.Run()
}

func getDbSession() *mgo.Database {

	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Panic(err)
	}
	db := session.DB("WidgetDb")
	return db
}
