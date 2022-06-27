package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/helper"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/model/domain"
)

// mengimplementasi kontrak dari BunkerServiceRepository
type BunkerServiceRepositoryImpl struct {
}

// buat method repository baru yang returnnya adalah BunkerServiceRepositoryImpl
func NewBunkerServiceRepository() BunkerServiceRepository {
	return &BunkerServiceRepositoryImpl{}
}

// membuat implementasi bunker service repostory mengikuti kontrak dari bunker service repository
// method Save() digunakan untuk membuat data baru
func (repository *BunkerServiceRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, bunkerService domain.BunkerService) domain.BunkerService {
	// buat query sql untuk membuat data yang baru
	SQL := "INSERT INTO bunker_services(no_so, nama_perusahaan, nama_kapal, nama_produk, jumlah_pengisian, pelabuhan, nopol_truk, nama_operator, description) VALUES (?,?,?,?,?,?,?,?,?)"

	// eksekusi query sql
	result, err := tx.ExecContext(ctx, SQL, bunkerService.NoSO, bunkerService.NamaPerusahaan, bunkerService.NamaKapal, bunkerService.NamaProduk, bunkerService.JumlahPengisian, bunkerService.Pelabuhan, bunkerService.NopolTruk, bunkerService.NamaOperator, bunkerService.Description)
	helper.PanicIfError(err)

	// mendapatkan id yang terakhir
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	// set id
	bunkerService.Id = int(id) // convert int64 ke integer

	return bunkerService
}

// method Update() digunakan untuk mengupdate data yang sudah ada sebelumnya
func (repository *BunkerServiceRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, bunkerService domain.BunkerService) domain.BunkerService {
	// buat query sql untuk mengupdate data
	SQL := "UPDATE bunker_services SET no_so = ?, nama_perusahaan = ?, nama_kapal = ?, nama_produk = ?, jumlah_pengisian = ?, pelabuhan = ?, nopol_truk = ?, nama_operator = ?, description = ? WHERE id = ?"

	// eksekusi query sql
	_, err := tx.ExecContext(ctx, SQL, bunkerService.NoSO, bunkerService.NamaPerusahaan, bunkerService.NamaKapal, bunkerService.NamaProduk, bunkerService.JumlahPengisian, bunkerService.Pelabuhan, bunkerService.NopolTruk, bunkerService.NamaOperator, bunkerService.Description, bunkerService.Id)
	helper.PanicIfError(err)

	return bunkerService
}

// method Delete() digunakan untuk menghapus data yang sudah ada sebelumnya
func (repository *BunkerServiceRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, bunkerService domain.BunkerService) {
	// buat query sql untuk menghapus data
	SQL := "DELETE FROM bunker_services WHERE id = ?"

	// eksekusi query sql
	_, err := tx.ExecContext(ctx, SQL, bunkerService.Id)
	helper.PanicIfError(err)
}

func (repository *BunkerServiceRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bunkerServiceId int) (domain.BunkerService, error) {
	// buat query sql untuk mencari data berdasarkan id
	SQL := "SELECT id, no_so, nama_perusahaan, nama_kapal, nama_produk, jumlah_pengisian, pelabuhan, nopol_truk, nama_operator, description FROM bunker_services WHERE id = ?"

	// eksekusi query sql
	rows, err := tx.QueryContext(ctx, SQL, bunkerServiceId)
	helper.PanicIfError(err)
	defer rows.Close()

	// data bunker service
	bunkerService := domain.BunkerService{}

	// lakukan pengechekan data jika ada dan data kosong
	if rows.Next() {
		// ambil data hasil query
		err := rows.Scan(&bunkerService.Id, &bunkerService.NoSO, &bunkerService.NamaPerusahaan, &bunkerService.NamaKapal, &bunkerService.NamaProduk, &bunkerService.JumlahPengisian, &bunkerService.Pelabuhan, &bunkerService.NopolTruk, &bunkerService.NamaOperator, &bunkerService.Description)
		helper.PanicIfError(err)

		return bunkerService, nil

	} else {

		return bunkerService, errors.New("data bunker service tidak ditemukan")
	}
}

func (repository *BunkerServiceRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.BunkerService {
	// buat query sql untuk mencari semua data
	SQL := "SELECT id, no_so, nama_perusahaan, nama_kapal, nama_produk, jumlah_pengisian, pelabuhan, nopol_truk, nama_operator, description FROM bunker_services"

	// eksekusi query sql
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	// slice data bunker service
	var bunkerServices []domain.BunkerService

	// hasil query jika datanya ada
	for rows.Next() {
		bunkerService := domain.BunkerService{}

		// ambil data hasil query
		err := rows.Scan(&bunkerService.Id, &bunkerService.NoSO, &bunkerService.NamaPerusahaan, &bunkerService.NamaKapal, &bunkerService.NamaProduk, &bunkerService.JumlahPengisian, &bunkerService.Pelabuhan, &bunkerService.NopolTruk, &bunkerService.NamaOperator, &bunkerService.Description)
		helper.PanicIfError(err)

		// menambah data
		bunkerServices = append(bunkerServices, bunkerService)
	}

	// return slice bunker services
	return bunkerServices
}
