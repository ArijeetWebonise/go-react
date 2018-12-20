package app

import (
	"database/sql"
	"net/http"

	"github.com/ArijeetBaruah/MyBlog/app/config"
	"github.com/ArijeetBaruah/MyBlog/app/models"
	"github.com/ArijeetBaruah/MyBlog/pkg/logger"
	"github.com/ArijeetBaruah/MyBlog/pkg/session"
	"github.com/ArijeetBaruah/MyBlog/pkg/templates"
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
