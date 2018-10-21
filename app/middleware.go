package app

import (
	"net/http"

	"github.com/Arijeet-webonise/go-react/pkg/framework"
	"github.com/gorilla/csrf"
)

//RenderView renders a view
func (app *App) RenderView(viewHandler func(*framework.Response, *framework.Request)) http.Handler {
	return app.CSRF(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := framework.NewResponse(w)
		req := framework.Request{Request: r}
		viewHandler(&res, &req)
	}))
}

// Handle will be serving only those requests that dont need to be authed
func (app *App) Handle(handler func(*framework.Response, *framework.Request)) http.Handler {
	return app.CSRF(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := framework.NewResponse(w)
		req := framework.Request{Request: r}
		handler(&res, &req)
		res.Write()
	}))
}

// UnsafeHandle handles Request without CSRF
func (app *App) UnsafeHandle(handler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, csrf.UnsafeSkipCheck(r))
	})
}
