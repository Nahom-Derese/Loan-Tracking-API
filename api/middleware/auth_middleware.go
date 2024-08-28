package middleware

import (
	"net/http"
	"strings"

	custom_error "github.com/Nahom-Derese/Loan-Tracking-API/domain/errors"
	tokenutil "github.com/Nahom-Derese/Loan-Tracking-API/internal/auth"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, custom_error.ErrMessage(custom_error.ErrInvalidToken))
			return
		}
		authToken := t[1]
		authorized, err := tokenutil.IsAuthorized(authToken, secret)

		if err != nil || !authorized {
			c.AbortWithStatusJSON(http.StatusUnauthorized, custom_error.ErrMessage(err))
			return
		}

		claims, err := tokenutil.ExtractUserClaimsFromToken(authToken, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, custom_error.ErrMessage(err))
			return
		}
		c.Set("x-user-id", claims["id"])
		c.Set("x-user-role", claims["role"])
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.MustGet("x-user-role")
		if role != "admin" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, custom_error.ErrorMessage{Message: "unauthorized"})
			ctx.Abort()
		}
		ctx.Next()
	}
}
