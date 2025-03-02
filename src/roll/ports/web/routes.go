package web

import (
	"net/http"

	"go-stack-yourself/src/common/ports/web"
)

func NewRouting() *http.ServeMux {
	mux := http.NewServeMux()

	web.HandeFunc(mux, "GET /", RolldiceHandler)

	return mux
}
