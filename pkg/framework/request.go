package framework

import (
	"database/sql"
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/csrf"
)

// this struct basically adds a context to the http.Request so that
// authenticator or any other middleward could push out the data
// to main request handler
type Request struct {
	*http.Request
	context map[string]interface{}
}

type JsonNullInt64 struct {
	sql.NullInt64
}

func (v JsonNullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullInt64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

func (r *Request) Push(key string, value interface{}) {
	if r.context == nil {
		r.context = map[string]interface{}{}
	}
	r.context[key] = value
}

func (r *Request) Value(key string) interface{} {
	return r.context[key]
}

func (r *Request) QueryParam(key string) string {
	return r.URL.Query().Get(key)
}

func (r *Request) ReadBody() (map[string]interface{}, error) {
	return ReadBody(r.Request)
}

func ReadBody(r *http.Request) (map[string]interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	bodyMap := make(map[string]interface{})
	err := decoder.Decode(&bodyMap)
	if err != nil {
		return bodyMap, err
	}
	return bodyMap, nil
}

func (r *Request) Bind(v interface{}) error {
	return Bind(r.Request.Body, v)
}

func Bind(body io.ReadCloser, v interface{}) error {
	defer body.Close()
	err := json.NewDecoder(body).Decode(v)
	return err
}

func GetPublicIPFromRequest(r *http.Request) (string, error) {
	pubIp := strings.Split(r.RemoteAddr, ":")[0]
	fIps := r.Header["X-Forwarded-For"]
	if len(fIps) < 1 {
		return "", errors.New("no ip in X-Forwarded-For header")
	}
	pubIp = strings.TrimSpace(strings.Split(fIps[0], ",")[0])
	return pubIp, nil
}

// CSRFToken returns csrf token for a request
func (r *Request) CSRFToken() template.HTML {
	return csrf.TemplateField(r.Request)
}