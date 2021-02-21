package main

import (
	"flag"
	"fmt"
	"github.com/AleksandrAkhapkin/testWLine/intenal/back/server"
	"github.com/AleksandrAkhapkin/testWLine/intenal/back/server/handlers"
	"github.com/AleksandrAkhapkin/testWLine/intenal/back/service"
	"github.com/AleksandrAkhapkin/testWLine/intenal/clients/postgres"
	"github.com/AleksandrAkhapkin/testWLine/intenal/types/config"
	"github.com/AleksandrAkhapkin/testWLine/logger"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var startMessage = `Перед запуском программы необходимо указать подключение к базе в config/config.yaml
Помимо этого, вы можете использовать флаг --port что бы указать порт для запуска в формате ":8080"
`

func main() {
	fmt.Println(startMessage)

	//флаг - путь до конфига
	configPath := new(string)
	//флаг - порт для запуска
	port := new(string)
	flag.StringVar(configPath, "config-path", "config/config.yaml", "path to yaml config")
	flag.StringVar(port, "port", ":8080", "change port")
	flag.Parse()

	cnfFile, err := os.Open(*configPath)
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with os.Open"))
	}

	cnf := config.Config{}
	if err := yaml.NewDecoder(cnfFile).Decode(&cnf); err != nil {
		logger.LogFatal(errors.Wrap(err, "err with Decode config"))
	}

	pq, err := postgres.NewPostgres(cnf.PostgresDsn)
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with NewPostgres"))
	}

	srv := service.NewService(pq)

	handls := handlers.NewHandlers(srv)

	//Корректное закрытие базы при получении сигнала
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		_ = <-sigChan
		log.Println("Finish service")
		if err := pq.Close(); err != nil {
			logger.LogFatal(err)
		}
		os.Exit(0)
	}()

	//go srv.StartRewrite()

	server.StartServer(handls, *port)
}
