package helper

import (
	"encoding/json"
	"net/http"
)

// read dari JSON
func ReadFromRequestBody(r *http.Request, res interface{}) {
	// hal pertama yang perlu dilakukan adalah baca JSON-nya
	// tidak perlu konversi jadi string terlebih dahulu, bisa langsung
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(res)
	PanicIfError(err)
}

// write ke JSON
func WriteToResponseBody(w http.ResponseWriter, res interface{}) {
	// karena tipe datanya JSON, tambahkan writer header
	w.Header().Add("Content-Type", "application/json")

	// tulis data webResponse ke dalam JSON langsung ke writer
	// tidak perlu konversi dulu ke string
	encoder := json.NewEncoder(w)
	err := encoder.Encode(res)
	PanicIfError(err)
}
