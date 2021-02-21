package server

import (
	"github.com/AleksandrAkhapkin/testWLine/intenal/back/server/handlers"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func StartServer(handlers *handlers.Handlers, port string) {

	router := NewRouter(handlers)
	log.Println("Start service in port " + port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(errors.Wrap(err, "err with NewRouter"))
	}
}
