package ports

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-stack-yourself/src/common/ports/web"
)

type HandlerMock struct {
	handler func(http.ResponseWriter, *http.Request)
}

func (h *HandlerMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler(w, r)
}

func TestNewErrorHandler(t *testing.T) {
	a := assert.New(t)

	result := []byte("lorem")

	eh := web.NewErrorHandler(
		&HandlerMock{
			handler: func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write(result)
			},
		},
	)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	eh.ServeHTTP(w, r)

	a.Equal([]byte("lorem"), w.Body.Bytes())
}

func TestNewErrorHandlerInternalServerError(t *testing.T) {
	a := assert.New(t)
	errorMsg := "test error"
	eh := web.NewErrorHandler(
		&HandlerMock{
			handler: func(w http.ResponseWriter, r *http.Request) {
				panic(errorMsg)
			},
		},
	)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	eh.ServeHTTP(w, r)

	a.Equal(
		[]byte(fmt.Sprintf("Internal server error <b>%s</b>", errorMsg)),
		w.Body.Bytes(),
	)
	a.Equal(http.StatusInternalServerError, w.Code)
}

func TestNewErrorHandlerCustomNotFound(t *testing.T) {
	a := assert.New(t)

	eh := web.NewErrorHandler(
		&HandlerMock{
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
		},
	)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	eh.ServeHTTP(w, r)

	a.Equal(
		[]byte("<html><body><h1>Page not found</h1><p>The page you're looking for does not exist.</p></body></html>"),
		w.Body.Bytes(),
	)
	a.Equal(http.StatusNotFound, w.Code)
}
