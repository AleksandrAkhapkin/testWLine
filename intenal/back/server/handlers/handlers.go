package handlers

import (
	"github.com/AleksandrAkhapkin/testWLine/intenal/back/service"
	"net/http"
)

type Handlers struct {
	srv *service.Service
}

func NewHandlers(srv *service.Service) *Handlers {
	return &Handlers{
		srv: srv,
	}
}

func (h *Handlers) Ping(w http.ResponseWriter, _ *http.Request) {

	_, _ = w.Write([]byte("pong"))
}
