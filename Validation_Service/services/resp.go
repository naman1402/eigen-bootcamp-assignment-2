package services

// CustomResponse struct to represent a validation response
// Matches Othentic Validation Service spec
// data: true/false/null, error: bool, message: string or null

type CustomResponse struct {
	Data    interface{} `json:"data"`
	Error   bool        `json:"error"`
	Message interface{} `json:"message"`
}

// NewCustomResponse creates a new instance of CustomResponse for success
func NewCustomResponse(data interface{}, message interface{}) CustomResponse {
	return CustomResponse{
		Data:    data,
		Error:   false,
		Message: message,
	}
}

// NewCustomError creates a new instance of CustomResponse for error
func NewCustomError(message interface{}, data interface{}) CustomResponse {
	return CustomResponse{
		Data:    data,
		Error:   true,
		Message: message,
	}
}
