package web

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func HandeFunc(mux *http.ServeMux, pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	handler := http.HandlerFunc(decorateCollectorRoute(pattern, handlerFunc))

	mux.Handle(pattern, handler)
}

func DecorateWithMiddlewares(mux http.Handler) http.Handler {
	return otelhttp.NewHandler(NewErrorHandler(mux), "app")
}

func decorateCollectorRoute(pattern string, handleFunc func(http.ResponseWriter, *http.Request)) func(
	http.ResponseWriter,
	*http.Request,
) {
	if pattern == "/" {
		return NewRootRouteHandlerMiddleware(handleFunc)
	}

	return handleFunc
}
