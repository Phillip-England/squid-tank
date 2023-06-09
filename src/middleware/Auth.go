package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	sessionToken := c.GetHeader("Authorization")[7:]
	fmt.Println(sessionToken)
	fmt.Println("hit the middleware")
	c.Next()
}