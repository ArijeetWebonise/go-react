package framework

import (
	"errors"
	"io"
	"net/http"
)

var (
	DefaultErrorCode = http.StatusConflict
)

type Response struct {
	ResponseWriter http.ResponseWriter

	msg        string
	data       JSONResponse
	statusCode int
	err        error
	written    bool
	success    bool
}

func NewResponse(w http.ResponseWriter) Response {
	r := Response{}
	r.ResponseWriter = w
	r.data = JSONResponse{}
	r.statusCode = -1
	r.msg = "success"
	r.written = false
	r.success = true

	return r
}

func (r *Response) Data(data map[string]interface{}) {
	r.data = data
}

func (r *Response) PutInData(k string, v interface{}) {
	r.data[k] = v
}

func (r *Response) Written() {
	r.written = true
}

func (r *Response) StatusCode(code int) {
	r.statusCode = code
}

// use this if you want to send success false
// but http status as 200
func (r *Response) SetSuccess(flag bool) {
	r.success = flag
}

func (r *Response) Error(err error) {
	r.err = err
	//logger.E(r.err)
}

func (r *Response) BadRequest(err ...error) {
	r.statusCode = http.StatusBadRequest
	// currently it supports only one err
	if len(err) > 0 {
		r.err = err[0]
		//logger.E(r.err)
	}
}

func (r *Response) UnProcessableEntity(err ...error) {
	r.statusCode = http.StatusUnprocessableEntity
	// currently it supports only one err
	if len(err) > 0 {
		r.err = err[0]
		//logger.E(r.err)
	}
}

func (r *Response) NotFound(err ...error) {
	r.statusCode = http.StatusNotFound
	if len(err) > 0 {
		r.err = err[0]
		//logger.E(r.err)
	}

}

func (r *Response) Unauthorised(err ...error) {
	r.statusCode = http.StatusUnauthorized
	if len(err) > 0 {
		r.err = err[0]
		//logger.E(r.err)
	}
	return

}

func (r *Response) InternalError(err ...error) {
	r.statusCode = http.StatusInternalServerError
	if len(err) > 0 {
		r.err = err[0]
		//logger.E(r.err)
	}
}

func (r *Response) Conflict(err ...error) {
	r.statusCode = DefaultErrorCode
	if len(err) > 0 {
		r.err = err[0]
		//logger.E(r.err)
	}
}

func (r *Response) Message(msg string) {
	r.msg = msg

}

func (r *Response) Write() {
	if r.written {
		return
	}
	if r.err != nil || r.statusCode > 399 {
		r.writeErrorResponse()
		return
	}
	if r.statusCode == http.StatusFound { // redirect do not write anything
		return
	}
	r.writeResponse()
}

func (r *Response) writeResponse() {
	r.ResponseWriter.Header().Add("Content-Type", "application/json")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Auth-Key, Session-Key")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
	if r.statusCode == -1 {
		r.ResponseWriter.WriteHeader(http.StatusOK)
	} else {
		r.ResponseWriter.WriteHeader(r.statusCode)
	}
	res := JSONResponse{
		"message": r.msg,
		"success": r.success,
		"data":    r.data,
	}
	r.ResponseWriter.Write(res.ByteArray())
}

func (r *Response) writeErrorResponse() {
	r.ResponseWriter.Header().Add("Content-Type", "application/json")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Auth-Key")
	if r.statusCode == -1 {
		r.ResponseWriter.WriteHeader(DefaultErrorCode)
	} else {
		r.ResponseWriter.WriteHeader(r.statusCode)
	}
	if r.err == nil {
		switch r.statusCode {
		case http.StatusUnauthorized:
			r.err = errors.New("Unauthorized access")
		default:
			r.err = errors.New("Illegal request")
		}

	}

	res := JSONResponse{
		"message": r.err.Error(),
		"success": false,
		"data":    r.data,
	}
	r.ResponseWriter.Write(res.ByteArray())

}

func (r *Response) Redirect(url string, req *http.Request) {
	r.StatusCode(http.StatusFound)
	r.ResponseWriter.Header().Add("Content-Type", "application/json")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Auth-Key, Session-Key")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
	http.Redirect(r.ResponseWriter, req, url, r.statusCode)
}

// RenderHTML renders HTML pages
func (r *Response) RenderHTML(res string) {
	io.WriteString(r.ResponseWriter, res)
}
