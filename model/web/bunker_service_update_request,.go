package web

import "time"

// menambahkan validasi di create dan update
// data model update request
type BunkerServiceUpdateRequest struct {
	Id              int    `validate:"required" json:"id"`
	NoSO            string `validate:"required" json:"noSO"`
	NamaPerusahaan  string `validate:"required" json:"namaPerusahaan"`
	NamaKapal       string `validate:"required" json:"namaKapal"`
	NamaProduk      string `validate:"required" json:"namaProduk"`
	JumlahPengisian int    `validate:"required" json:"jumlahPengisian"`
	Pelabuhan       string `validate:"required" json:"pelabuhan"`
	NopolTruk       string `validate:"required" json:"nopolTruk"`
	NamaOperator    string `validate:"required" json:"namaOperator"`
	Description     string `validate:"required" json:"description"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
