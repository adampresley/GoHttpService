package GoHttpService

type BasicJsonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewBasicJsonResponse(success bool, message string) BasicJsonResponse {
	return BasicJsonResponse{
		Success: success,
		Message: message,
	}
}
