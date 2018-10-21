package app

import (
	"net/http"

	"github.com/Arijeet-webonise/go-react/pkg/framework"
)

// RenderIndex renders all HTML pages
func (app *App) RenderIndex(w *framework.Response, r *framework.Request) {
	tplList := []string{"./public/index.html"}

	res, err := app.TplParser.ParseTemplate(tplList, r.CSRFToken())
	if err != nil {
		app.Logger.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	w.RenderHTML(res)
}
