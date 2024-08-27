package middleware

import (
	custom_error "github.com/Nahom-Derese/Loan-Tracking-API/domain/errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}

		// status := custom_error.MapErrorToStatusCode(err)
		status := 1

		c.Status(status)
		error_message := custom_error.ErrMessage(err)
		msg := render.JSON{Data: error_message}
		msg.Render(c.Writer)
	}

}
