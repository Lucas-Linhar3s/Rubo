package application

import (
	"fmt"
	"time"

	"github.com/Lucas-Linhar3s/Rubo/database"
	"github.com/Lucas-Linhar3s/Rubo/modules/auth/domain"
	"github.com/Lucas-Linhar3s/Rubo/modules/graphql/model"
	"github.com/Lucas-Linhar3s/Rubo/pkg/config"
	v1 "github.com/Lucas-Linhar3s/Rubo/pkg/http/response/v1"
	"github.com/Lucas-Linhar3s/Rubo/pkg/jwt"
	"github.com/Lucas-Linhar3s/Rubo/pkg/log"
	"github.com/Lucas-Linhar3s/Rubo/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// AuthApp is the application for auth
type AuthApp struct {
	logger *log.Logger
	db     *database.Database
	config *config.Config
	jwt    *jwt.JWT
}

// NewAuthApp is a function that returns a new auth application
func NewAuthApp(
	logger *log.Logger,
	db *database.Database,
	config *config.Config,
	jwt *jwt.JWT,
) *AuthApp {
	return &AuthApp{
		logger: logger,
		db:     db,
		config: config,
		jwt:    jwt,
	}
}

// LoginWithEmailAndPassword is a function that returns a session out
func (app *AuthApp) LoginWithEmailAndPassword(ctx *gin.Context, req *model.LoginUserInput) (*model.SessionOut, error) {
	const msg = "Error on LoginWithEmailAndPassword"

	var (
		service = domain.GetService(domain.GetRepository(app.db))
		data    = new(domain.AuthModel)
		res     = new(model.SessionOut)
		err     error
	)

	if data, err = utils.ConvertRequestToModel[domain.AuthModel](req); err != nil {
		app.logger.Error(msg, zap.Error(err))
		return nil, err
	}
	if exist, err := service.VerifyEmail(*data.Email); err != nil {
		app.logger.Error(msg, zap.Error(err))
		return nil, err
	} else if !exist {
		return nil, v1.ErrEmailNotExists
	}

	if data, err = service.LoginWithEmailAndPassword(data); err != nil {
		app.logger.Error(msg, zap.Error(err))
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*data.Password), []byte(req.Password)); err != nil {
		return nil, v1.ErrInvalidPassword
	}

	if data.Role != nil {
		domain.RoleDefault = *data.Role
	}

	now := time.Now()
	expiresAt := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute()+app.config.Security.Jwt.ExpiresAt, now.Second(), now.Nanosecond(), now.Location())

	res.UserID = utils.GetString(data.ID)
	res.DataExpiracao = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", expiresAt.Year(), expiresAt.Month(), expiresAt.Day(), expiresAt.Hour(), expiresAt.Minute(), expiresAt.Second())

	token, err := app.jwt.GenToken("", *data.ID, *data.Email, domain.RoleDefault, expiresAt)
	if err != nil {
		return nil, err
	}
	res.AccessToken = utils.GetString(token)

	return res, nil
}
