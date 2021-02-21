package server

import (
	"github.com/AleksandrAkhapkin/testWLine/intenal/back/server/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(h *handlers.Handlers) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.Methods(http.MethodGet).Path("/ping").HandlerFunc(h.Ping)

	//получение данных
	router.Methods(http.MethodGet).Path("/get_data").HandlerFunc(h.GetData)

	return router
}
