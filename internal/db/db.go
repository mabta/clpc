package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Init 初始化数据库连接
func Init(dsn string) (err error) {
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	return nil
}

// IsNoRows 判断是否属于没有数据行的错误
func IsNoRows(err error) bool {
	return err == sql.ErrNoRows
}
