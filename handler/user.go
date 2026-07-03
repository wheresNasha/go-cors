package handler

// handler/user.go : handles user-related HTTP requests.
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 	Think of Context as a tray passed to the handler function that contains two things:
// The Input (Request): Headers, cookies, URL parameters, and JSON body sent by the user.
// The Output (Response): Tools to write back data to the user.

// Your function uses c.JSON() to write data into that context, which Gin then flushes out to the user as a final HTTP response.

func Retrieve(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World from GET (RETRIEVE)!",
	})
}

func Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Hello, World from POST (CREATE)!",
	})
}

func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World from PUT (UPDATE)!",
	})
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World from DELETE (DELETE)!",
	})
}

func Patch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World from PATCH!",
	})
}

func Head(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World from HEAD!",
	})
}

// What HTTP methods are allowed for this resource?
// Browser CORS preflight
// router.OPTIONS("/options", func(c *gin.Context) { // when someone hits /, execute the handler function
// 	c.Header("Allow", "GET,POST,PUT,DELETE,PATCH,HEAD,OPTIONS")
// 	//Without Access-Control-Allow-Origin, the browser blocks the request — even if status is 200.
// 	c.Status(http.StatusNoContent)
// })

// gin.H{"message": "Hello!"} is one line map no trailing comma
// gin.H{ is a multi-line map with a trailing comma: its a go syntax rule
// 	"message": "Hello!",
// }
