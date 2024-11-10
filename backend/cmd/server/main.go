package main

import (
	"fmt"

	"github.com/Lucas-Linhar3s/Rubo/di"
	middleware "github.com/Lucas-Linhar3s/Rubo/middlewares"
	"github.com/Lucas-Linhar3s/Rubo/pkg/graphql"
)

func main() {
	app, cleanUp, err := di.InitializeApp()
	if err != nil {
		panic(fmt.Errorf("failed to initialize app: %w", err))
	}
	defer cleanUp()

	app.Server.Router.Use(
		middleware.CORSMiddleware(),
		middleware.GinContextToContextMiddleware(),
	)

	graphql.InitGraphqlServer(app.Server.Router, app.HandlerServer)

	if err := app.Server.Run(app.Config); err != nil {
		app.Logger.Error(err.Error())
	}
}
