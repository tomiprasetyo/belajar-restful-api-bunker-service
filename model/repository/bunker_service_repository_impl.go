package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/helper"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/model/domain"
)

// binding uuid agar bisa dicari, dirubah/update
type GUIDBinding struct {
	GUID string `uri:"guid" binding:"required,uuid4"`
}

// mengimplementasi kontrak dari bunker service repository
type BunkerServiceRepositoryImpl struct {
}

// membuat implementasi bunker service repostory mengikuti kontrak dari bunker service repository
// method Save() digunakan untuk membuat data baru
func (repository *BunkerServiceRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, bunkerService domain.BunkerService) domain.BunkerService {
	// global unique identifier
	// tujuannya agar tidak mengekspos id asli dari database
	var guid = uuid.New().String()

	// buat query sql untuk membuat data yang baru
	SQL := "INSERT INTO bunker_services(guid, no_so, nama_perusahaan, nama_kapal, nama_produk, jumlah_pengisian, pelabuhan, nopol_truk, nama_operator, description) VALUES (?,?,?,?,?,?,?,?,?,?)"

	// eksekusi query sql
	result, err := tx.ExecContext(ctx, SQL, guid, bunkerService.NoSO, bunkerService.NamaPerusahaan, bunkerService.NamaKapal, bunkerService.NamaProduk, bunkerService.JumlahPengisian, bunkerService.Pelabuhan, bunkerService.NopolTruk, bunkerService.NamaOperator, bunkerService.Description)
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
	var binding GUIDBinding

	// buat query sql untuk mengupdate data
	SQL := "UPDATE bunker_services SET no_so = ?, nama_perusahaan = ?, nama_kapal = ?, nama_produk = ?, jumlah_pengisian = ?, pelabuhan = ?, nopol_truk = ?, nama_operator = ?, description = ? WHERE guid = ?"

	// eksekusi query sql
	_, err := tx.ExecContext(ctx, SQL, bunkerService.NoSO, bunkerService.NamaPerusahaan, bunkerService.NamaKapal, bunkerService.NamaProduk, bunkerService.JumlahPengisian, bunkerService.Pelabuhan, bunkerService.NopolTruk, bunkerService.NamaOperator, bunkerService.Description, binding.GUID)
	helper.PanicIfError(err)

	return bunkerService
}

// method Delete() digunakan untuk menghapus data yang sudah ada sebelumnya
func (repository *BunkerServiceRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, bunkerService domain.BunkerService) {
	var binding GUIDBinding

	// buat query sql untuk menghapus data
	SQL := "DELETE FROM bunker_services WHERE guid = ?"

	// eksekusi query sql
	_, err := tx.ExecContext(ctx, SQL, binding.GUID)
	helper.PanicIfError(err)
}

func (repository *BunkerServiceRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bunkerService domain.BunkerService) (domain.BunkerService, error) {
	var binding GUIDBinding

	// buat query sql untuk mencari data berdasarkan guid
	SQL := "SELECT guid, no_so, nama_perusahaan, nama_kapal, nama_produk, jumlah_pengisian, pelabuhan, nopol_truk, nama_operator, description FROM bunker_services WHERE guid = ?"

	// eksekusi query sql
	rows, err := tx.QueryContext(ctx, SQL, binding.GUID)
	helper.PanicIfError(err)

	// data bunker service
	dataBunkerService := domain.BunkerService{}

	// lakukan pengechekan data jika ada dan data kosong
	if rows.Next() {
		// ambil data hasil query
		err := rows.Scan(&bunkerService.GUID, &bunkerService.NoSO, &bunkerService.NamaPerusahaan, &bunkerService.NamaKapal, &bunkerService.NamaProduk, &bunkerService.JumlahPengisian, &bunkerService.Pelabuhan, &bunkerService.NopolTruk, &bunkerService.NamaOperator, &bunkerService.Description)
		helper.PanicIfError(err)

		return dataBunkerService, nil

	} else {

		return dataBunkerService, errors.New("data bunker service tidak ditemukan")
	}
}

func (repository *BunkerServiceRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, bunkerService domain.BunkerService) []domain.BunkerService {
	// buat query sql untuk mencari semua data
	SQL := "SELECT guid, no_so, nama_perusahaan, nama_kapal, nama_produk, jumlah_pengisian, pelabuhan, nopol_truk, nama_operator, description FROM bunker_services"

	// eksekusi query sql
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	// slice data bunker service
	var bunkerServices []domain.BunkerService

	// hasil query jika datanya ada
	for rows.Next() {
		bunkerService := domain.BunkerService{}

		// ambil data hasil query
		err := rows.Scan(&bunkerService.GUID, &bunkerService.NoSO, &bunkerService.NamaPerusahaan, &bunkerService.NamaKapal, &bunkerService.NamaProduk, &bunkerService.JumlahPengisian, &bunkerService.Pelabuhan, &bunkerService.NopolTruk, &bunkerService.NamaOperator, &bunkerService.Description)
		helper.PanicIfError(err)

		// menambah data
		bunkerServices = append(bunkerServices, bunkerService)
	}

	// return slice bunker services
	return bunkerServices
}
