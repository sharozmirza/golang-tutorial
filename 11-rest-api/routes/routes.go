package routes

import (
	"example.com/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // GET request handler
	server.GET("/events/:id", getEvent) // /events/1, /events/5

	/*
		One way of registering the Authentication middleware to a request handler
		is to use it like below:

		server.POST("/events", middleware.Authenticate, createEvent)

		In this case the functions will be executed in left to right order.
		So, if the user is not authenticated, the handle will abort the request and
		the createEvent function will not be reached. But this needs to added to
		all the individual request handler that will be using this middleware.
	*/

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent) // POST request handler
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
