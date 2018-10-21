package app

//InitRoute initilize Route
func (app *App) InitRoute() {
	app.Router.Get("/api/v1/ping", app.Handle(app.ping))
	app.Router.Post("/api/v1/login", app.UnsafeHandle(app.Handle(app.Login)))

	app.Router.Get("/:any", app.RenderView(app.RenderIndex))
}
