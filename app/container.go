package app

import (
	"database/sql"
	"net/http"

	"github.com/Arijeet-webonise/go-react/app/config"
	"github.com/Arijeet-webonise/go-react/app/models"
	"github.com/Arijeet-webonise/go-react/pkg/logger"
	"github.com/Arijeet-webonise/go-react/pkg/session"
	"github.com/Arijeet-webonise/go-react/pkg/templates"
	"github.com/go-zoo/bone"
	"github.com/graphql-go/handler"
)

// App wrapper for go application
type App struct {
	Router            *bone.Mux
	Config            config.Config
	Logger            logger.ILogger
	TplParser         templates.ITemplateParser
	DB                *sql.DB
	APIHandler        *handler.Handler
	CSRF              func(http.Handler) http.Handler
	FlashService      session.ISessionService
	UserService       models.UserService
	CustomUserService models.CustomUserService
	GraphQlService    GraphQlService
}
