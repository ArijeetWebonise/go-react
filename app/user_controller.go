package app

import (
	"errors"

	"github.com/Arijeet-webonise/go-react/pkg/framework"
)

// Login api for login
func (app *App) Login(w *framework.Response, r *framework.Request) {
	body := make(map[string]string)

	if err := r.Bind(&body); err != nil {
		app.Logger.Error(err)
		w.BadRequest(err)
		return
	}

	user := body["user"]
	pass := body["pass"]

	if user != "admin@admin" && pass != "admin" {
		err := errors.New("username or password are incorrect")
		app.Logger.Error(err)
		w.BadRequest(err)
		return
	}

	if _, err := app.FlashService.CreateSession(r.Request, w.ResponseWriter, 1); err != nil {
		app.Logger.Error(err)
		w.InternalError(err)
		return
	}
	w.Message("Login Successful")
}
