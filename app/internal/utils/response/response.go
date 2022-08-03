package response

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func Error(message string, err string, data interface{}) Response {
	spError := strings.Split(err, "\n")
	return Response{
		Status:  false,
		Message: message,
		Errors:  spError,
		Data:    data,
	}
}
