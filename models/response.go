package models

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}

func SetResponse(Status int, Message string, Data []interface{}) (res Response) {
	res = Response{
		Status:  Status,
		Message: Message,
		Data:    Data,
	}
	return
}
