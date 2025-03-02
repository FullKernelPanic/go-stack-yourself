package web

import (
	"bytes"
	"fmt"
	"net/http"
)

func NewErrorHandler(h http.Handler) *ErrorHandler {
	return &ErrorHandler{Handler: h}
}

type ErrorHandler struct {
	Handler http.Handler
}

func (h *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec := NewResponseRecorder(w)

	defer func() {
		if err := recover(); err != nil {
			customInternalServerErrorHandler(w, r, err)
		}
	}()

	h.Handler.ServeHTTP(rec, r)

	switch rec.StatusCode {
	case http.StatusNotFound:
		customNotFoundHandler(w, r)

	default:
		rec.Persist()
	}
}

func customInternalServerErrorHandler(w http.ResponseWriter, r *http.Request, err any) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(fmt.Sprintf("Internal server error <b>%s</b>", err)))
}

type ResponseRecorder struct {
	Writer     http.ResponseWriter
	StatusCode int
	Headers    http.Header
	Body       *bytes.Buffer
}

func NewResponseRecorder(w http.ResponseWriter) *ResponseRecorder {
	return &ResponseRecorder{
		Writer:     w,
		Headers:    http.Header{},
		Body:       new(bytes.Buffer),
		StatusCode: http.StatusOK,
	}
}

func (rr *ResponseRecorder) Header() http.Header {
	return rr.Headers
}

func (rr *ResponseRecorder) Write(data []byte) (int, error) {
	return rr.Body.Write(data)
}

func (rr *ResponseRecorder) WriteHeader(statusCode int) {
	rr.StatusCode = statusCode
}

func (rr *ResponseRecorder) Persist() {
	for key, values := range rr.Headers {
		for _, value := range values {
			rr.Writer.Header().Add(key, value)
		}
	}

	rr.Writer.WriteHeader(rr.StatusCode)

	_, _ = rr.Writer.Write(rr.Body.Bytes())
}

func customNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("<html><body><h1>Page not found</h1><p>The page you're looking for does not exist.</p></body></html>"))
}

type RootRouteHandlerMiddleware struct {
	handlerFunc func(http.ResponseWriter, *http.Request)
}

func NewRootRouteHandlerMiddleware(handlerFunc func(http.ResponseWriter, *http.Request)) func(
	http.ResponseWriter,
	*http.Request,
) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handlerFunc(w, r)
			return
		}

		http.NotFound(w, r)
	}
}
