package controller

import (
	"github.com/devproje/plog/log"
	"github.com/devproje/simple-chat/database"
	"github.com/devproje/simple-chat/model"
	"github.com/gocql/gocql"
	"time"
)

func Logging(data *model.Log) {
	var conn *gocql.Session
	var stmt string
	var err error

	conn, err = database.Open()
	if err != nil {
		log.Errorln(err)
		return
	}
	defer database.Close(conn)

	stmt = `insert into log(id, type, author, content, created) values (?, ?, ?, ?, ?)`
	query := conn.Query(stmt, gocql.TimeUUID(), data.Type, data.Author, data.Content, time.Now().Format(time.RFC3339))
	err = query.Exec()
	if err != nil {
		return
	}
}
