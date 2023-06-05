package e

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type HttpError struct {
	Message string
	Code int
}

func NewHttpError(message string, code int) *HttpError {
	return &HttpError{
		Message: message,
		Code: code,
	}
}

func (v *HttpError) Error() string {
	return fmt.Sprintf(v.Message)
}

func (v *HttpError) HttpExit(c *gin.Context) {
	c.JSON(v.Code, gin.H{
		"msg": v.Message,
	})
}