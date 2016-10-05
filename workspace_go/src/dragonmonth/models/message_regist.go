package models

type MessageInsertUser struct {
	MessageNSQ
}

func NewMessageInsertUser(opt int, sql string, bean interface{}) *MessageInsertUser {
	return &MessageInsertUser{
		MessageNSQ: MessageNSQ{
			Option: opt,
			Sql:    sql,
			Bean:   bean,
		},
	}
}
