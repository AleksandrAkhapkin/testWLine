package server

import (
	"github.com/AleksandrAkhapkin/testWLine/intenal/back/server/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(h *handlers.Handlers) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.Methods(http.MethodGet).Path("/ping").HandlerFunc(h.Ping)

	return router
}
