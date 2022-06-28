package middleware

import (
	"net/http"

	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/helper"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/model/web"
)

// sesuai dengan kontak api yang sudah dibuat, kita perlu menambahkan
// authentication menggunakan API-key
// dibuat dalam bentuk middleware
//
// buat struct dengan kontrak handler
// middleware harus dalam bentuk handler
type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(h http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: h}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// cek value di header
	if "SECRETKEY" == r.Header.Get("X-API-Key") {
		// ok
		m.Handler.ServeHTTP(w, r)
	} else {
		// error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(w, webResponse)
	}

}
