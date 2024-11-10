package utils

import (
	"context"
	"errors"

	middleware "github.com/Lucas-Linhar3s/Rubo/middlewares"
	"github.com/gin-gonic/gin"
)

// GinContextFromContext is a helper function to extract gin.Context from context.Context
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(middleware.GinContextKeyInstance)
	if ginContext == nil {
		return nil, errors.New("Gin Context not found")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, errors.New("Gin Context not found")
	}

	return gc, nil
}
