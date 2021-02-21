package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(dsn string) (*Postgres, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "err with Open DB")
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "err with ping DB")
	}

	return &Postgres{db}, nil
}

func (p *Postgres) Close() error {
	return p.db.Close()
}

//Задает порядковый ID запуска сервера
func (p *Postgres) InsertNewServerID() (int32, error) {

	var id int32
	err := p.db.QueryRow("INSERT INTO server_data (js_data) VALUES ($1) RETURNING id", "").Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

//получить поле js_data по id запуска сервера
func (p *Postgres) GetJSDataByID(id int32) (string, error) {

	var data string
	err := p.db.QueryRow("SELECT js_data FROM server_data WHERE id = $1", id).Scan(&data)
	if err != nil {
		return "", err
	}

	return data, nil
}
