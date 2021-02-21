package service

import (
	"fmt"
	"github.com/AleksandrAkhapkin/testWLine/intenal/clients/postgres"
	"github.com/AleksandrAkhapkin/testWLine/logger"
	"github.com/pkg/errors"
	"time"
)

type Service struct {
	p        *postgres.Postgres
	id       int32
	lastData string
}

func NewService(pq *postgres.Postgres) *Service {

	return &Service{
		p: pq,
	}
}

//Создание айди сервиса и запуск получения js_data раз в 30 секунд
func (s *Service) StartService() {

	var err error

	//Инициализируем в базе новую строчку которая отвечает за текущий сервер
	s.id, err = s.p.InsertNewServerID()
	if err != nil {
		logger.LogError(errors.Wrap(err, "err in StartService in pq InsertNewServerID"))
		return
	}
	fmt.Printf("Индефикатор запускаемого сервера: %d", s.id)

	ch := make(chan struct{})
	go s.updateData(ch)
	go func() {
		for {
			// в паралелль запускаем бесконечный цикл вычитывания канала
			_ = <-ch
			logger.LogInfo("restart checkLog")
			go s.updateData(ch)
			//если мы вычитали канал - значит была ошибка и надо запустить функцию заново
		}
	}()
}

//раз в 30 секунд получает данные из таблицы и обновляет данные в s.data
func (s *Service) updateData(ch chan struct{}) {

	//используется тайм тикер, а не тайм слип, так как слип накапливает время выполнения самого процесса
	tt := time.NewTicker(30 * time.Second)
	for {
		data, err := s.p.GetJSDataByID(s.id)
		if err != nil {
			logger.LogError(errors.Wrap(err, "err in service updateData in pq GetJSDataByID "))
			ch <- struct{}{}
			return
		}
		s.lastData = data
		_ = <-tt.C
	}
}
