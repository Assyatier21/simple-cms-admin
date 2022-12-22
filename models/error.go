package models

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func SetError(Status int, Message string) (err Error) {
	err = Error{
		Status:  Status,
		Message: Message,
	}
	return
}
