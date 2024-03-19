package database

import (
	"github.com/devproje/simple-chat/config"
	"github.com/gocql/gocql"
	"github.com/lithammer/dedent"
)

func Init() error {
	conn, err := Open()
	if err != nil {
		return err
	}

	stmt := dedent.Dedent(`create table if not exists log(
    	id	    uuid 	primary key,
    	type    varchar,
    	author  varchar,
    	content varchar,
    	created varchar
	);`)

	err = conn.Query(stmt).Exec()
	if err != nil {
		return err
	}

	return nil
}

func Open() (*gocql.Session, error) {
	cluster := gocql.NewCluster(config.Load().Database.Clusters...)
	cluster.Keyspace = config.Load().Database.KeySpace
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: config.Load().Database.Credential.Username,
		Password: config.Load().Database.Credential.Password,
	}

	session, err := cluster.CreateSession()
	return session, err
}

func Close(session *gocql.Session) {
	if session.Closed() {
		return
	}

	session.Close()
}
