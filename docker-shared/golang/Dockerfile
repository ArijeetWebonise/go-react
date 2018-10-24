FROM golang

WORKDIR $GOPATH/src/github.com/Arijeet-webonise/go-react

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