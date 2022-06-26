package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/helper"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/model/domain"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/model/web"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/repository"
)

// implementasi kontrak service dari BunkerServiceServe
// membutuhkan koneksi ke database
type BunkerServiceServeImpl struct {
	// membutuhkan repository karena manipulasi datanya menggunakan repository
	BunkerServiceRepository repository.BunkerServiceRepository

	// koneksi ke database
	DB *sql.DB

	// attribut validate
	Validate *validator.Validate
}

// implementasi kontrak
func (service BunkerServiceServeImpl) Create(ctx context.Context, request web.BunkerServiceCreateRequest) web.BunkerServiceResponse {
	// sebelum melakukan create data / transaksinya dimulai
	// kita lakukan validasi terlebih dahulu
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// karena kita menggunakan database transaksional menggunakan mysql
	// berarti kita butuh requestnya dalam bentuk transaksional

	// begin transaction
	// hal yang pertama dilakukan adalah me-wrap method ini dalam bentuk transaksional
	// lalu kita juga harus detect, jika terjadi error maka harus di rollback transaksinya
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	// jika error rollback, jika tidak commit transaksinya
	// lakukan hal ini di semua method kontrak
	defer helper.CommitOrRollback(tx)

	bunkerService := domain.BunkerService{
		NoSO:            request.NoSO,
		NamaPerusahaan:  request.NamaPerusahaan,
		NamaKapal:       request.NamaKapal,
		NamaProduk:      request.NamaProduk,
		JumlahPengisian: request.JumlahPengisian,
		Pelabuhan:       request.Pelabuhan,
		NopolTruk:       request.NopolTruk,
		NamaOperator:    request.NamaOperator,
		Description:     request.Description,
	}

	bunkerService = service.BunkerServiceRepository.Save(ctx, tx, bunkerService)

	// konversi data
	return helper.ToBunkerServiceResponse(bunkerService)

}

func (service BunkerServiceServeImpl) Update(ctx context.Context, request web.BunkerServiceUpdateRequest) web.BunkerServiceResponse {
	// sebelum melakukan update data / transaksinya dimulai
	// kita lakukan validasi terlebih dahulu
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// validasi data terlebih dahulu
	bunkerService, err := service.BunkerServiceRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	// jika sudah ketemua datanya
	bunkerService.NoSO = request.NoSO
	bunkerService.NamaPerusahaan = request.NamaPerusahaan
	bunkerService.NamaKapal = request.NamaKapal
	bunkerService.NamaProduk = request.NamaProduk
	bunkerService.JumlahPengisian = request.JumlahPengisian
	bunkerService.Pelabuhan = request.Pelabuhan
	bunkerService.NopolTruk = request.NopolTruk
	bunkerService.NamaOperator = request.NamaOperator
	bunkerService.Description = request.Description

	bunkerService = service.BunkerServiceRepository.Update(ctx, tx, bunkerService)

	// konversi data
	return helper.ToBunkerServiceResponse(bunkerService)
}

func (service BunkerServiceServeImpl) Delete(ctx context.Context, bunkerServiceId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// validasi data terlebih dahulu
	bunkerService, err := service.BunkerServiceRepository.FindById(ctx, tx, bunkerServiceId)
	helper.PanicIfError(err)

	service.BunkerServiceRepository.Delete(ctx, tx, bunkerService)
}

func (service BunkerServiceServeImpl) FindById(ctx context.Context, bunkerServiceId int) web.BunkerServiceResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	bunkerService, err := service.BunkerServiceRepository.FindById(ctx, tx, bunkerServiceId)
	helper.PanicIfError(err)

	return helper.ToBunkerServiceResponse(bunkerService)

}

func (service BunkerServiceServeImpl) FindAll(ctx context.Context) []web.BunkerServiceResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	bunkerServices := service.BunkerServiceRepository.FindAll(ctx, tx)

	return helper.ToBunkerServiceResponses(bunkerServices)
}
