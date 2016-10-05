package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

/**
 * 用户基本信息，存在mysql中
 */
type User struct {
	UserId         string `xorm:"notnull unique pk"`
	Emial          string `xorm:"index"`
	Telephone      string `xorm:"notnull unique index"`
	Password       string
	InvitationCode string
	Gender         int       //0:男 1:女
	RegistrationAt time.Time `xorm:"created"`
	UpdateAt       time.Time `xorm:"updated"`
}

func NewUser(email, tel, pwd, invc string, gender int) *User {
	u := &User{
		UserId:         bson.NewObjectId().Hex(),
		Emial:          email,
		Telephone:      tel,
		Password:       pwd,
		InvitationCode: invc,
		Gender:         gender,
	}
	return u
}
