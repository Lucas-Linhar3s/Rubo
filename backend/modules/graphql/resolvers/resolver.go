package resolvers

import (
	"github.com/Lucas-Linhar3s/Rubo/modules/auth/application"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthApp *application.AuthApp
}
