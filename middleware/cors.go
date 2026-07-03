package middleware

// cors.go : handles cross-origin requests from the frontend dev server.
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS handles cross-origin requests from the frontend dev server.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")

		// c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,HEAD,OPTIONS")
		// as delete is not in the list, it will not work : cors error on hitting button of delete on UI

		// GET, POST, HEAD are already allowed by default
		// both buttons : CREATE and DELETE are working
		// c.Header("Access-Control-Allow-Methods", "PUT,DELETE,PATCH,OPTIONS")

		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,HEAD,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		// Preflight: answer OPTIONS and stop (don't hit your POST handler)
		if c.Request.Method == http.MethodOptions {
			// if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
