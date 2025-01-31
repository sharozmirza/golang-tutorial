package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // GET request handler
	server.GET("/events/:id", getEvent) // /events/1, /events/5
	server.POST("/events", createEvent) // POST request handler
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
}
