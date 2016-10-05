package models

// {
//     "code": 0,
//     "msg": "OK",
//     "result": {
//         "count": 1,   //成功发送的短信个数
//         "fee": 0.05,     //扣费金额，70个字按一条计费，超出70个字时按每67字一条计费，类型：双精度浮点型/double
//         "sid": 1097   //短信id；多个号码时以该id+各手机号尾号后8位作为短信id,
//                       //（数据类型：64位整型，对应Java和C#的long，不可用int解析)
//     }
// }
type PYSecurityCodeMsgSendedReturn struct {
	Count int     `json:"count"`
	Fee   float64 `json:"fee"`
	Sid   int64   `json:"sid"`
}

/**
 *片云验证码发送后的返回数据
 *保存在mongodb中
 */
type PYSecurityCodeMsgSendedResult struct {
	Result PYSecurityCodeMsgSendedReturn `json:"result"`
	Code   int                           `json:"code"`
	Msg    string                        `json:"msg"`
}
