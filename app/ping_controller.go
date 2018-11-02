package app

import (
	"github.com/ArijeetBaruah/go-react/pkg/framework"
)

//Ping will indicate the health
func (a App) ping(w *framework.Response, r *framework.Request) {
	//go:generate echo hello form ping
	a.Logger.Info("hello from the log side")
	w.Message("pong")
}
