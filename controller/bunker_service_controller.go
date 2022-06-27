package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// membuat BunkerService controller / Web API
// dibuat dalam bentuk service agar terprediksi apa yang mau dibuat
// buat method tiap API-nya
// methodnya sudah standar karena harus mengikuti handler dari HTTP
type BunkerServiceController interface {
	// parameternya mengikuti dari HTTP handler
	// karena kita menggunakan library httpRouter kita tambahkan Params sebagai parameter
	// karena di kontraknya harus ada Params
	// semua sama kalo untuk controller
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
