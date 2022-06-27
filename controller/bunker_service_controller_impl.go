package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/helper"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/model/web"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/service"
)

// implementasi detail dari BunkerService controller
// hal pertama yang dibutuhkan adalah controller ini cukuplah service
type BunkerServiceControllerImpl struct {
	BunkerServiceServe service.BunkerServiceServe
}

func (c *BunkerServiceControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// // decode data dari http request
	bunkerServiceCreateRequest := web.BunkerServiceCreateRequest{}

	// dereference pointer karena data berupa struct
	helper.ReadFromRequestBody(r, &bunkerServiceCreateRequest)

	// dari controller kita panggil servicenya
	// lalu kita  kirim contextnya
	bunkerServiceResponse := c.BunkerServiceServe.Create(r.Context(), bunkerServiceCreateRequest)

	// selanjutnya buat responsenya
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bunkerServiceResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (c *BunkerServiceControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// sama seperti create, yang membedakan di path ada parameter id
	// setelah kita lakukan parsing update requestnya, kita set id-nya
	bunkerServiceUpdateRequest := web.BunkerServiceUpdateRequest{}
	// dereference pointer karena data berupa struct
	helper.ReadFromRequestBody(r, &bunkerServiceUpdateRequest)

	// id-nya kita ambil dari params
	// dikonversi dulu dari string menjadi int
	bunkerServiceId := p.ByName("bunkerServiceId")
	id, err := strconv.Atoi(bunkerServiceId)
	helper.PanicIfError(err)

	bunkerServiceUpdateRequest.Id = id

	bunkerServiceResponse := c.BunkerServiceServe.Update(r.Context(), bunkerServiceUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bunkerServiceResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (c *BunkerServiceControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// untuk delete tidak butuh body dan response, cukup id-nya saja
	bunkerServiceId := p.ByName("bunkerServiceId")
	id, err := strconv.Atoi(bunkerServiceId)
	helper.PanicIfError(err)

	c.BunkerServiceServe.Delete(r.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (c *BunkerServiceControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// hampir sama dengan delete, bedanya ada responsenya
	bunkerServiceId := p.ByName("bunkerServiceId")
	id, err := strconv.Atoi(bunkerServiceId)
	helper.PanicIfError(err)

	bunkerServiceResponse := c.BunkerServiceServe.FindById(r.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bunkerServiceResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *BunkerServiceControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// untuk FindAll cukup butuh context saja dan returnnya slice
	bunkerServiceResponses := c.BunkerServiceServe.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bunkerServiceResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
