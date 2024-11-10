package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// InitGraphqlServer initializes the graphql server
func InitGraphqlServer(r *gin.Engine, h *handler.Server) {
	group := r.Group("/graphql")

	{
		group.POST("/query", func(c *gin.Context) {
			h.ServeHTTP(c.Writer, c.Request)
		})
		group.GET("/", playgroundHandler())
	}
}
