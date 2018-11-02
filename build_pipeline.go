package pipeline

//go:generate echo "fetching dependencies"

//go:generate echo "running migrations"
//go:generate sql-migrate up
//go:generate echo "generating models using xo"
//go:generate xo pgsql://$GR_DB_USERNAME:$GR_DB_PASSWORD@$GR_DB_HOST/$GR_DB_NAME?sslmode=disable -o app/models --suffix=.xo.go --template-path templates/

//go:generate echo "performing packr clean"
//go:generate packr clean
//go:generate echo "running `packr` command"
//go:generate packr
//go:generate echo "done!"
