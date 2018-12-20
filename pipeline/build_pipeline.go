package pipeline

//go:generate echo "fetching dependencies"
//go:generate go get github.com/lib/pq
//go:generate go get -v github.com/rubenv/sql-migrate/...
//go:generate go get -u github.com/gobuffalo/packr/...
//go:generate go get github.com/spf13/viper
//go:generate go get "github.com/Sirupsen/logrus"
//go:generate go get "github.com/go-zoo/bone"
//go:generate go get github.com/gorilla/csrf
//go:generate go get -u golang.org/x/tools/cmd/goimports
//go:generate go get golang.org/x/crypto/bcrypt
//go:generate go get -u  github.com/mitchellh/gox
//go:generate go get -u github.com/mitchellh/mapstructure
//go:generate go get -u github.com/xo/xo
//go:generate go get github.com/satori/go.uuid
//go:generate go get -u github.com/gorilla/sessions

//go:generate echo "running migrations"
//go:generate sql-migrate up
//go:generate echo "generating models using xo"
//go:generate xo pgsql://$GR_DB_USERNAME:$GR_DB_PASSWORD@$GR_DB_HOST/$GR_DB_NAME?sslmode=disable -o app/models --suffix=.xo.go --template-path templates/

//go:generate echo "performing packr clean"
//go:generate packr clean
//go:generate echo "running `packr` command"
//go:generate packr
//go:generate echo "done!"
