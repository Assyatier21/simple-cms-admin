package models

type Response struct {
	Code    int           `json:"code"`
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}
type ResponseLogin struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SetResponse(code int, status string, message string, data []interface{}) Response {
	return Response{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func SetResponseLogin(code int, status string, message string, data interface{}) ResponseLogin {
	return ResponseLogin{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}
}
