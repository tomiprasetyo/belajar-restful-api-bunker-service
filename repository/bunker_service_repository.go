package repository

import (
	"context"
	"database/sql"

	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/model/domain"
)

// data access layer ke domain bunker service menggunakan repository pattern
// membuat kontrak menggunakan interface
// yang nanti akan diimplementasi menggunakan struct
type BunkerServiceRepository interface {
	// buat operasi CRUD
	// di golang biasakan menggunakan context sebagai parameter pertama
	// saat menggunakan database, terutama database relational ada baiknya membuat function yang mendukung yang namanya database transactional
	// digunakan diparameter kedua
	// contoh di golang menggunakan sql.Tx
	// parameter ketiga menggunakan data aslinya
	// return data bunker service
	Save(ctx context.Context, tx *sql.Tx, bunkerService domain.BunkerService) domain.BunkerService
	Update(ctx context.Context, tx *sql.Tx, bunkerService domain.BunkerService) domain.BunkerService
	Delete(ctx context.Context, tx *sql.Tx, bunkerService domain.BunkerService)                  // delete tidak perlu return data bunker service
	FindById(ctx context.Context, tx *sql.Tx, bunkerServiceId int) (domain.BunkerService, error) // return error jika data tidak ditemukan
	FindAll(ctx context.Context, tx *sql.Tx) []domain.BunkerService                              // return slice data bunker service
}
