package web

import "time"

// data model response
type BunkerServiceResponse struct {
	Id              int    `json:"id"`
	NoSO            string `json:"noSO"`
	NamaPerusahaan  string `json:"namaPerusahaan"`
	NamaKapal       string `json:"namaKapal"`
	NamaProduk      string `json:"namaProduk"`
	JumlahPengisian int    `json:"jumlahPengisian"`
	Pelabuhan       string `json:"pelabuhan"`
	NopolTruk       string `json:"nopolTruk"`
	NamaOperator    string `json:"namaOperator"`
	Description     string `json:"description"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
