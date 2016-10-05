package models

type RequestRegisteInfo struct {
	Tel    string `form:"tel"`
	Code   string `form:"code"`
	Pwd    string `form:"pwd"`
	RePwd  string `form:"repwd"`
	Gender int    `form:"gender"`
}
