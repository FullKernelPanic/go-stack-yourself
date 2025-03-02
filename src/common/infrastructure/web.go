package infrastructure

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

func NewHTTPServer(ctx context.Context, port int, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      handler,
	}
}
