package service

import (
	"context"

	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/model/web"
)

type GUID struct {
	GUID string `uri:"guid" binding:"required,uuid4"`
}

// implementasi bunker service serve atau bisnis logic
// buat kontrak service terlebih dahulu dalam bentuk interface
// biasanya service function itu jumlahnya mengikuti APInya
// biasanya dalam satu API cuma manggil satu function disebuah service
// methodnya sama dengan yang ada di repository, yang beda ada di logic method create
// biasakan gunakan context di golang
// buat model baru khusus untuk representasi dari request
type BunkerServiceServe interface {
	// tidak direkomendasikan return domain secara langsung
	// karena domain adalah representasi dari tabel
	// misal saja kita tidak mau mengekspos semua data yang ada di tabel.
	// jika kita return-kan domain secara langsung bahayanya,
	// jika nanti ada satu atau dua kolom yang sensitif datanya akan terekspos di API
	// jika tidak kita filter maka akan sangat berbahaya
	// disarankan untuk buat data model request dan data model response-nya
	Create(ctx context.Context, request web.BunkerServiceCreateRequest) web.BunkerServiceResponse
	Update(ctx context.Context, request web.BunkerServiceUpdateRequest) web.BunkerServiceResponse
	Delete(ctx context.Context, guid GUID)
	FindById(ctx context.Context, guid GUID) web.BunkerServiceResponse
	FindAll(ctx context.Context) []web.BunkerServiceResponse
}
