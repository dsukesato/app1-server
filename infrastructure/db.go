package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dsukesato/go13/pbl/app1-server/interfaces/database"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBConn struct {
	conn *sql.DB
}

func Init() database.DBConn {
	// Cloud SQL環境での開発
	//var (
	//	connectionName = mustGetenv("CLOUDSQL_CONNECTION_NAME")
	//	user           = mustGetenv("CLOUDSQL_USER")
	//	dbName         = os.Getenv("CLOUDSQL_DATABASE_NAME") // NOTE: dbName may be empty
	//	password       = os.Getenv("CLOUDSQL_PASSWORD")      // NOTE: password may be empty
	//	socket         = os.Getenv("CLOUDSQL_SOCKET_PREFIX")
	//)
	//
	//if socket == "" {
	//	socket = "/cloudsql"
	//}
	//
	//dbURI := fmt.Sprintf("%s:%s@unix(%s/%s)/%s?parseTime=true", user, password, socket, connectionName, dbName)

	// ローカル環境での開発
	var user = "root"
	var password = "lookin"
	var dbName = "pbl_app1"

	dbURI := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", user, password, dbName)
	conn, err := sql.Open("mysql", dbURI)

	if err != nil {
		panic(fmt.Sprintf("DB: %v", err))
	}

	dbc := new(DBConn)
	dbc.conn = conn

	return dbc

	// connection string format: user=USER password=PASSWORD host=/cloudsql/PROJECT_ID:REGION_ID:INSTANCE_ID/[ dbname=DB_NAME]
	// dbURI := fmt.Sprintf("user=%s password=%s host=/cloudsql/%s dbname=%s", user, password, connectionName, dbName)
	// conn, err := sql.Open("postgres", dbURI)
}

// Cloud SQL環境での開発
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}

func (db *DBConn) ExecContext(ctx context.Context, query string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := db.conn.Exec(query, args...)
	if err != nil {
		return res, err
	}
	res.Result = result

	return res, nil
}

func (db *DBConn) QueryContext(ctx context.Context, query string, args ...interface{}) (database.Row, error) {
	rows, err := db.conn.Query(query, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows

	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
