package main

import (
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)    // GET request handler
	server.GET("/events/:id", getEvent) // /events/1, /events/5
	server.POST("/events", createEvent) // POST request handler

	server.Run(":8080") // localhost:8080
}

// handler for GET /events route
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

// handler for GET /events:id route
func getEvent(context *gin.Context) {
	// context.Param("id") -> returns "id" as string
	// ParseInt() interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value.
	// Since we are interested in a base 10 int64, we passed 10 and 64
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

// handler for POST /events route
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

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
