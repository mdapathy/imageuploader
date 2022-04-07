package httpapi

import (
	"github.com/gorilla/mux"
)

func configureRoutes(resources *resources) *mux.Router {
	handler := imagesHandler{resources}
	router := mux.NewRouter()

	private := router.PathPrefix("/api/v1/private/gallery/").Subrouter()
	private.Use(UserMiddleware)
	private.Handle("/images", handler.List()).Methods("GET")
	private.Handle("/images/{id}", handler.Details()).Methods("GET")

	private.Handle("/images", handler.Create()).Methods("POST")
	private.Handle("/images/{id}", handler.Delete()).Methods("DELETE")

	return router
}
