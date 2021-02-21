package server

import (
	"github.com/AleksandrAkhapkin/testWLine/intenal/back/server/handlers"
	"github.com/AleksandrAkhapkin/testWLine/logger"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

//запускаем сервер
func StartServer(handlers *handlers.Handlers, port string) {

	router := NewRouter(handlers)
	log.Println("Start service in port " + port)
	if err := http.ListenAndServe(port, router); err != nil {
		logger.LogFatal(errors.Wrap(err, "err with NewRouter"))
	}
}
