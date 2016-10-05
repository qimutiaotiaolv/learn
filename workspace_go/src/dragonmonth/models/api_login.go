package models

import (
	"labix.org/v2/mgo/bson"
)

type ApiLoginResponse struct {
	Tocken string
	UserID string
}

func NewApiLoginResponse(userID string) *ApiLoginResponse {
	bean := &ApiLoginResponse{
		UserID: userID,
		Tocken: bson.NewObjectId().Hex(),
	}
	return bean
}
