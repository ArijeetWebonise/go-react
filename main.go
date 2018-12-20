package main

import (
	"net/http"

	"github.com/ArijeetBaruah/MyBlog/app"
	"github.com/ArijeetBaruah/MyBlog/app/config"
	"github.com/ArijeetBaruah/MyBlog/app/models"
	"github.com/ArijeetBaruah/MyBlog/pkg/database"
	"github.com/ArijeetBaruah/MyBlog/pkg/logger"
	"github.com/ArijeetBaruah/MyBlog/pkg/session"
	"github.com/ArijeetBaruah/MyBlog/pkg/templates"
	"github.com/go-zoo/bone"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/graphql-go/handler"
)

func main() {
	logger := &logger.RealLogger{}
	logger.Initialise()

	cfg := &config.AppConfig{
		Logger: logger,
	}
	cfg = cfg.ConstructAppConfig()

	db := &database.DatabaseWrapper{}
	dbConn, dbErr := db.Initialise(&cfg.DB)
	if dbErr != nil {
		logger.Panic(dbErr)
		return
	}

	CSRF := csrf.Protect([]byte(cfg.CSRFAuthkey))

	a := app.App{
		Router:    bone.New(),
		Config:    cfg,
		Logger:    logger,
		DB:        dbConn,
		TplParser: &templates.TemplateParser{},
		FlashService: &session.CookieStoreServiceImpl{
			Store:  sessions.NewCookieStore([]byte(cfg.SessionAuthkey)),
			Secure: false,
		},
		CSRF:              CSRF,
		UserService:       &models.UserServiceImpl{DB: dbConn},
		CustomUserService: &models.CustomUserServiceImpl{DB: dbConn},
		GraphQlService:    &app.GraphQlServiceImpl{DB: dbConn},
	}

	schema, err := a.GraphQlService.GetSchema()
	if err != nil {
		panic(err)
	}

	a.APIHandler = handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     cfg.GraphQL.Pretty,
		GraphiQL:   cfg.GraphQL.GraphiQL,
		Playground: cfg.GraphQL.Playground,
	})

	a.GraphQlService.GetApp(&a)

	a.InitRoute()

	if err := http.ListenAndServe(cfg.Port, a.Router); err != nil {
		panic(err)
	}
}
