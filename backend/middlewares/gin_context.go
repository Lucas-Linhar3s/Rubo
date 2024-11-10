package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ginContextKey string

const GinContextKeyInstance ginContextKey = "GinContextKey"

// GinContextToContextMiddleware is a middleware that adds the gin.Context to the request context
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContextKeyInstance, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
