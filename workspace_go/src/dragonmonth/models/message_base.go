package models

const (
	MQ_INSERT = iota
	MQ_UPDATE
	MQ_REMOVE
)

type MessageNSQ struct {
	Option int         `json:"option"`
	Sql    string      `json:"sql"`
	Bean   interface{} `json:"bean"`
}
