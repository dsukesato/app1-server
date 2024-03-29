package database

import "context"

type DBConn interface {
	ExecContext(context.Context, string, ...interface{}) (Result, error)
	QueryContext(context.Context, string, ...interface{}) (Row, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
