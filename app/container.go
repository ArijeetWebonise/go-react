package app

import (
	"net/http"

	"github.com/Arijeet-webonise/go-react/app/config"
	"github.com/Arijeet-webonise/go-react/pkg/logger"
	"github.com/Arijeet-webonise/go-react/pkg/session"
	"github.com/Arijeet-webonise/go-react/pkg/templates"
	"github.com/go-zoo/bone"
)

// App wrapper for go application
type App struct {
	Router       *bone.Mux
	Config       config.Config
	Logger       logger.ILogger
	TplParser    templates.ITemplateParser
	CSRF         func(http.Handler) http.Handler
	FlashService session.ISessionService
}
