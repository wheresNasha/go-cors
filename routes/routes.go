package routes

// routes.go : wires all HTTP routes to their handlers.
//  WHAT URLs exist?
//  WHAT HTTP methods are allowed for each URL?
//  WHAT handler functions are called for each URL?
//  WHAT data is sent to and received from each URL?
//  WHAT HTTP status codes are returned for each URL?
//  WHAT HTTP headers are sent to and received from each URL?
//  WHAT HTTP body is sent to and received from each URL?
//  WHAT HTTP cookies are sent to and received from each URL?
//  WHAT HTTP authentication is required for each URL?

import (
	"go-crud/handler"

	"github.com/gin-gonic/gin"
)

// Register wires all HTTP routes to their handlers.
func Register(router *gin.Engine) {

	// when someone hits /retrieve, execute the handler function
	router.GET("/retrieve", handler.Retrieve)

	// create a new resource
	router.POST("/create", handler.Create)

	// update the entire resource
	router.PUT("/update", handler.Update)

	// delete the entire resource
	router.DELETE("/delete", handler.Delete)

	// update a particular field of a resource
	router.PATCH("/patch", handler.Patch)

	// HEAD is similar to GET, but without the response body.
	// It's useful for checking if a resource exists without downloading the content.
	// Postman shows the status (200) but no message because HTTP forbids a body on HEAD responses. Gin honors that even if you call c.JSON().
	//
	// Think of HEAD as a quick "ping" to check if a file exists without downloading it.
	// Purpose is to read metadata like size, last modified date, etc. without downloading the content.
	// Health checks on large objects without downloading the content.
	// "Does this file exist? How big is it?"
	router.HEAD("/head", handler.Head)
}
