package EasyGo

import "time"

type Result struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Time int64       `json:"time"`
	Data interface{} `json:"data"`
}

func NewResultInstance() *Result {
	return new(Result)
}

func (r *Result) SetCode(code string) {
	r.Code = code
}

func (r *Result) SucessDefault(data interface{}) {
	r.Code = "0"
	r.Msg = "处理成功"
	r.Time = time.Now().Unix()
	r.Data = data
}

func (r *Result) Sucess(msg string, code string, data interface{}) {
	r.Code = "0"
	r.Msg = "处理成功"
	r.Time = time.Now().Unix()
	r.Data = data
}

func (r *Result) Fail(msg string, code string, data interface{}) {
	r.Code = code
	r.Msg = msg
	r.Time = time.Now().Unix()
	r.Data = data
}

func (r *Result) GetCode() string {
	return r.Code
}

func (r *Result) GetMsg() string {
	return r.Msg
}

func (r *Result) SetMsg(msg string) {
	r.Msg = msg
}

func (r *Result) SetTime(time int64) {
	r.Time = time
}

func (r *Result) SetData(data interface{}) {
	r.Data = data
}
