package routes

import (
	"computerextra/elaerning-go/handler"
	"computerextra/elaerning-go/templates"

	"github.com/gorilla/mux"
)

func GetRoutes(router *mux.Router) {

	// Main Route
	router.Handle("/", handler.Component(templates.Index()))

}
