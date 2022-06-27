package exception

// membuat struct untuk not found
// mengikuti kontrak interface error
type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}
