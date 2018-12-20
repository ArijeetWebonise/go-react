FROM golang

WORKDIR $GOPATH/src/github.com/ArijeetBaruah/MyBlog

RUN go get -v -u github.com/oxequa/realize
RUN go get github.com/lib/pq
RUN go get -v github.com/rubenv/sql-migrate/...
RUN go get -u github.com/gobuffalo/packr/...
RUN go get github.com/spf13/viper
RUN go get "github.com/Sirupsen/logrus"
RUN go get "github.com/go-zoo/bone"
RUN go get "github.com/knq/dburl"
RUN go get -u golang.org/x/tools/cmd/goimports
RUN go get -u  github.com/mitchellh/gox
RUN go get -u github.com/mitchellh/mapstructure
RUN go get -u github.com/xo/xo
RUN go get -u github.com/go-zoo/bone
RUN go get "github.com/robfig/cron"
RUN go get github.com/satori/go.uuid
RUN go get -u github.com/gorilla/sessions
RUN go get -u github.com/gorilla/csrf
RUN go get -u github.com/tdewolff/minify

ENV GR_PORT ":9999"
ENV GR_CSRF_AUTH_KEY "8d0e5181ce96016b027ac2a792ade2fe5f22a5811583d85b0c35ef2b0aa498ea"
ENV GR_SESSION_AUTH_KEY "8d0e5181ce96016b027ac2a792ade2fe5f1fa5811583d85b0c35ef2b0aa498ea"
ENV GR_DB_USERNAME "local"
ENV GR_DB_PASSWORD "local"
ENV GR_DB_HOST "db"
ENV GR_DB_NAME "MyBlog"
ENV GR_DB_DRIVER_NAME "postgres"
ENV GR_DB_DATA_SOURCE "host=db user=local password=toor dbname=MyBlog sslmode=disable"
ENV GR_GRAPHIQL false
ENV GR_GRAPHQL_PLAYGROUND true
ENV GR_GRAPHQL_PRETTY true

CMD ["realize", "start"]

RUN go get github.com/graphql-go/graphql
RUN go get github.com/graphql-go/handler
