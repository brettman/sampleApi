package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// investigate DI from google:
//  https://github.com/google/wire
func main() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/widget", widgetHandler)
	}

	router.Run()
}

func widgetHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "this is a message that all is ok")
}
