package biz

type Resp struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Total int         `json:"total,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func RespSuccess(data interface{}, total ...int) *Resp {
	resp := &Resp{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
	if len(total) > 0 {
		resp.Total = total[0]
	}
	return resp
}

func RespFail(bizCode int, errMsg ...string) *Resp {
	var msg string
	if len(errMsg) > 0 {
		msg = errMsg[0]
	} else {
		msg = CodeText(bizCode)
	}
	return &Resp{
		Code: bizCode,
		Msg:  msg,
	}
}
