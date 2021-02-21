package service

import (
	"fmt"
	"github.com/AleksandrAkhapkin/testWLine/intenal/clients/postgres"
	"github.com/AleksandrAkhapkin/testWLine/pkg/logger"
	"github.com/pkg/errors"
	"time"
)

type Service struct {
	p        *postgres.Postgres //база
	id       int32              //индетификатор текущего сервера
	lastData string
}

//инициализация нового сервиса
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
		logger.LogFatal(errors.Wrap(err, "err in StartService in pq InsertNewServerID"))
	}
	fmt.Printf("Индефикатор запускаемого сервера: %d\n", s.id)

	//используется тайм тикер, а не тайм слип, так как слип накапливает время выполнения самого процесса
	tt := time.NewTicker(30 * time.Second)
	for {
		data, err := s.p.GetJSDataByID(s.id)
		if err != nil {
			logger.LogFatal(errors.Wrap(err, "err in service updateData in pq GetJSDataByID "))
		}
		s.lastData = data
		_ = <-tt.C
	}
}

func (s *Service) GetData() string {
	return s.lastData
}
