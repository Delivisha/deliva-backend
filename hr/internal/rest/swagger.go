package rest

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"net/http"
)

//go:embed index.html
//go:embed api.swagger.json
var swaggerUI embed.FS

func RegisterSwagger(mux *chi.Mux) error {
	const specRoot = "/hr-spec/"

	// mount the swagger specification
	mux.Mount(specRoot, http.StripPrefix(specRoot, http.FileServer(http.FS(swaggerUI))))

	return nil
}
