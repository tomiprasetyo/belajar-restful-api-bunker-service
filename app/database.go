package app

import (
	"database/sql"
	"time"

	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/helper"
)

// method untuk koneksi ke database
func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_restful_api_bunker_service")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
