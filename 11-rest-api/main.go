package main

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)    // GET request handler
	server.POST("/events", createEvent) // POST request handler

	server.Run(":8080") // localhost:8080
}

// handler for GET /events endpoint
func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

// handler for POST /events endpoint
func createEvent(context *gin.Context) {
	var event models.Event

	// ShouldBindJSON() is used to bind incoming JSON data from an HTTP request body to a Go struct
	// In this case it binds the incoming JSON data to Event model
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
