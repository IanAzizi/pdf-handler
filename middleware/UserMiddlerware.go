// middleware/usermiddleware
package middleware

import (
	"fmt"
	"net/http"

	"pdf-APP/services"

	"github.com/gin-gonic/gin"
)

func UserMiddleware(c *gin.Context) {
	fmt.Println("QUERY Id =", c.Query("Id"))
	fmt.Println("QUERY password =", c.Query("password"))

	id := c.Query("Id")
	password := c.Query("password")

	if id == "" || password == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "missing credentials",
		})
		return
	}

	user, ok := services.Authenticate(id, password)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	c.Set("user", user)
	c.Next()
}
