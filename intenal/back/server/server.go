package server

import (
	"github.com/AleksandrAkhapkin/testWLine/intenal/back/server/handlers"
	"github.com/AleksandrAkhapkin/testWLine/pkg/logger"
	"github.com/pkg/errors"
	"net/http"
)

//запускаем сервер
func StartServer(handlers *handlers.Handlers, port string) {

	router := NewRouter(handlers)
	logger.LogInfo("Start service in port " + port)
	if err := http.ListenAndServe(port, router); err != nil {
		logger.LogFatal(errors.Wrap(err, "err with NewRouter"))
	}
}
