package service

import (
	"github.com/AleksandrAkhapkin/testWLine/intenal/clients/postgres"
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

//func( s *Service) StartRewrite() {
//
//	ch := make(chan struct{})
//
//	go s.updateData(ch)
//	go func() {
//		for {
//			// в паралелль запускаем бесконечный цикл вычитывания канала
//			_ = <-ch
//			log.Println("restart checkLog")
//			go s.updateData(ch)
//			//если мы вычитали канал - значит была ошибка и надо запустить функцию заново
//		}
//	}()
//}
