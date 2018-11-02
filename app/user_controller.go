package app

import (
	"errors"

	"github.com/ArijeetBaruah/go-react/pkg/framework"
	"golang.org/x/crypto/bcrypt"
)

// Login api for login
func (app *App) Login(w *framework.Response, r *framework.Request) {
	body := make(map[string]string)

	if err := r.Bind(&body); err != nil {
		app.Logger.Error(err)
		w.BadRequest(err)
		return
	}

	email := body["user"]
	pass := body["pass"]

	user, err := app.CustomUserService.GetUser(email)
	if err != nil {
		app.Logger.Error(err)
		err := errors.New("username or password are incorrect")
		w.InternalError(err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		app.Logger.Error(err)
		err = errors.New("username or password are incorrect")
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
