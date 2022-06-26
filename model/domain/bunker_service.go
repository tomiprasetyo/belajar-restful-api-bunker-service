package domain

import "time"

// data model domain bunker service
type BunkerService struct {
	Id              int       `json:"id"`
	NoSO            string    `json:"noSO"`
	NamaPerusahaan  string    `json:"namaPerusahaan"`
	NamaKapal       string    `json:"namaKapal"`
	NamaProduk      string    `json:"namaProduk"`
	JumlahPengisian int       `json:"jumlahPengisian"`
	Pelabuhan       string    `json:"pelabuhan"`
	NopolTruk       string    `json:"nopolTruk"`
	NamaOperator    string    `json:"namaOperator"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
