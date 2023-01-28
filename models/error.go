package models

type Error struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SetError(code int, status string, message string) Error {
	return Error{
		Code:    code,
		Status:  status,
		Message: message,
	}
}
