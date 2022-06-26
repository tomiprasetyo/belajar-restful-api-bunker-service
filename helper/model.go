package helper

import (
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/model/domain"
	"github.com/tomiprasetyo/belajar-restful-api-bunker-service/model/web"
)

// konversi data BunkerService ke BunkerServiceResponse
func ToBunkerServiceResponse(bunkerService domain.BunkerService) web.BunkerServiceResponse {
	return web.BunkerServiceResponse{
		Id:              bunkerService.Id,
		NoSO:            bunkerService.NoSO,
		NamaPerusahaan:  bunkerService.NamaPerusahaan,
		NamaKapal:       bunkerService.NamaKapal,
		NamaProduk:      bunkerService.NamaProduk,
		JumlahPengisian: bunkerService.JumlahPengisian,
		Pelabuhan:       bunkerService.Pelabuhan,
		NopolTruk:       bunkerService.NopolTruk,
		NamaOperator:    bunkerService.NamaOperator,
		Description:     bunkerService.Description,
		CreatedAt:       bunkerService.CreatedAt,
		UpdatedAt:       bunkerService.UpdatedAt,
	}
}

// konversi data BunkerService ke BunkerServiceResponses
func ToBunkerServiceResponses(bunkerServices []domain.BunkerService) []web.BunkerServiceResponse {
	var bunkerServiceResponses []web.BunkerServiceResponse
	for _, bunkerService := range bunkerServices {
		bunkerServiceResponses = append(bunkerServiceResponses, ToBunkerServiceResponse(bunkerService))
	}

	return bunkerServiceResponses
}
