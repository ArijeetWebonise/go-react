package app

//InitRoute initilize Route
func (app *App) InitRoute() {
	app.Router.Handle("/graphql", app.APIHandler)

	app.Router.Get("/:any", app.RenderView(app.RenderIndex))
}
