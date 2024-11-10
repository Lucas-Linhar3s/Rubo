package di

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Lucas-Linhar3s/Rubo/database"
	"github.com/Lucas-Linhar3s/Rubo/modules/auth/application"
	"github.com/Lucas-Linhar3s/Rubo/modules/graphql/generated"
	"github.com/Lucas-Linhar3s/Rubo/modules/graphql/resolvers"
	"github.com/Lucas-Linhar3s/Rubo/pkg/config"
	"github.com/Lucas-Linhar3s/Rubo/pkg/http/server"
	"github.com/Lucas-Linhar3s/Rubo/pkg/jwt"
	"github.com/Lucas-Linhar3s/Rubo/pkg/log"
)

// App is the main application
type App struct {
	Config        *config.Config
	Server        *server.Server
	HandlerServer *handler.Server
	Logger        *log.Logger
	Jwt           *jwt.JWT
}

type modules struct {
	authApp *application.AuthApp
}

func initializeModules(logger *log.Logger, database *database.Database, config *config.Config, jwt *jwt.JWT) *modules {
	// Initialize modules here
	authApp := application.NewAuthApp(logger, database, config, jwt)
	return &modules{
		authApp: authApp,
	}
}

// InitializeApp initializes the application
func InitializeApp() (*App, func(), error) {
	viper := config.NewViper()
	config := config.LoadAttributes(viper)
	logger := log.NewLog(config)
	jwt := jwt.NewJwt(config)
	database := database.NewDatabase(config, logger)
	modules := initializeModules(logger, database, config, jwt)
	handlerServer := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{
			AuthApp: modules.authApp,
		},
	}))

	server := server.NewServer()

	return &App{
		Config:        config,
		Server:        server,
		HandlerServer: handlerServer,
		Logger:        logger,
		Jwt:           jwt,
	}, func() {}, nil
}
