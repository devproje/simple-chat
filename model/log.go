package model

import "github.com/gocql/gocql"

type Log struct {
	ID      gocql.UUID `cql:"id"`
	Type    string     `cql:"type"`
	Author  string     `cql:"author"`
	Content string     `cql:"content"`
	Created string     `cql:"created"`
}
