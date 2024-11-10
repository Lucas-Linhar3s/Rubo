package infrastructure

import "time"

// AuthModel is a struct that represent the model of auth
type AuthModel struct {
	ID           *string     `copier:"ID"`
	Email        *string     `copier:"Email"`
	Password     *string     `copier:"Password"`
	PasswordHash *string     `copier:"PasswordHash"`
	Name         *string     `copier:"Name"`
	Picture      *string     `copier:"Picture"`
	Role         *string     `copier:"Role"`
	CreatedAt    *time.Timer `copier:"CreatedAt"`
	UpdatedAt    *time.Timer `copier:"UpdatedAt"`
}
