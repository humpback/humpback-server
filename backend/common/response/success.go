package response

import (
	"net/http"
)

type SucceedMsg struct {
	StatusCode int    `json:"statusCode"`
	Msg        string `json:"msg"`
}

func NewRespSucceed() *SucceedMsg {
	return &SucceedMsg{
		StatusCode: http.StatusOK,
		Msg:        "Succeed",
	}
}
