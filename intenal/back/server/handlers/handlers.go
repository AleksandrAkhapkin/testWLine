package handlers

import (
	"encoding/json"
	"github.com/AleksandrAkhapkin/testWLine/intenal/back/service"
	"github.com/AleksandrAkhapkin/testWLine/pkg/logger"
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

func (h *Handlers) GetData(w http.ResponseWriter, r *http.Request) {

	data := h.srv.GetData()
	resp := struct {
		JsData string `json:"js_data"`
	}{JsData: data}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.LogError(err)
	}

}
