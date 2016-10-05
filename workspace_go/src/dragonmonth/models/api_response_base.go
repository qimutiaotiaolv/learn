package models

import (
	"encoding/json"
	"time"
)

type ApiResponseBase struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	SystemTime int64       `json:"system_time"`
	Datas      interface{} `json:"datas"`
}

func NewResponseModel(code int, message string, datasModel interface{}) *ApiResponseBase {
	model := &ApiResponseBase{
		Code:       code,
		Message:    message,
		Datas:      datasModel,
		SystemTime: time.Now().Unix(),
	}
	return model
}

func (this *ApiResponseBase) Json() (string, error) {
	buf, err := json.Marshal(this)
	if err != nil {
		return "", err
	}
	return string(buf), err
}
