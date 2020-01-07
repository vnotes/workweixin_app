package dbs

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	DB = sqlx.MustConnect("mysql", "notes:notes@tcp(mysql_db:3306)/weixin")
}

func Cli() *sqlx.DB {
	return DB
}