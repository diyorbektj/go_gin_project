package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, Message string, Data interface{}) Response {
	res := Response{
		Status:  status,
		Message: Message,
		Error:   nil,
		Data:    Data,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")

	res := Response{
		Status:  false,
		Message: message,
		Error:   splittedError,
		Data:    data,
	}

	return res
}
