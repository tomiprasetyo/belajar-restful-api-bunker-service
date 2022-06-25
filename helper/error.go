package helper

// membuat fungsi error agar tidak ada duplikasi kode
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
