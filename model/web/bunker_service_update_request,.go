package web

import "time"

// menambahkan validasi di create dan update
// data model update request
type BunkerServiceUpdateRequest struct {
	Id              int    `validate:"required"`
	NoSO            string `validate:"required,min=1,max=100" json:"noSO"`
	NamaPerusahaan  string `validate:"required,min=1,max=100" json:"namaPerusahaan"`
	NamaKapal       string `validate:"required,min=1,max=100" json:"namaKapal"`
	NamaProduk      string `validate:"required,min=1,max=100" json:"namaProduk"`
	JumlahPengisian int    `validate:"required,min=1,max=100" json:"jumlahPengisian"`
	Pelabuhan       string `validate:"required,min=1,max=100" json:"pelabuhan"`
	NopolTruk       string `validate:"required,min=1,max=100" json:"nopolTruk"`
	NamaOperator    string `validate:"required,min=1,max=100" json:"namaOperator"`
	Description     string `validate:"required,min=1,max=100" json:"description"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
