package web

// buat struct untuk standar responsenya
type WebResponse struct {
	Code   int
	Status string
	Data   interface{}
}
