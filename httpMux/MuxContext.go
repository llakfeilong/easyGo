package httpMux

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type MuxContext struct {
	Writer    http.ResponseWriter
	Request   *http.Request
	Validator *Validator
}

func NewContext(w http.ResponseWriter, r *http.Request) *MuxContext {
	return &MuxContext{
		Writer:    w,
		Request:   r,
		Validator: NewValid(),
	}
}

/**
 * 获取表单参数值
 */
func (c *MuxContext) Query(parameter string) string {
	return c.Request.FormValue(parameter)
}

/**
 * 获取表单参数值
 */
func (c *MuxContext) PostFormValue(parameter string) string {
	return c.Request.PostFormValue(parameter)
}

/**
 * 获取表单参数值
 */
func (c *MuxContext) PostFormGet(parameter string) string {
	return c.Request.PostForm.Get(parameter)
}

/**
 * 获取表单参数值
 */
func (c *MuxContext) Get(key string) string {
	return c.Request.Form.Get(key)
}

/**
 * 获取post提交json
 */
func (c *MuxContext) GetPostJson() []byte {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	return body
}

/**
 * 获取文件表单参数
 */
func (c *MuxContext) FormValue(key string) string {
	return c.Request.FormValue(key)
}

/**
 * 解析表单文件上传
 */
func (c *MuxContext) ParseMultipartForm(maxMemory int64) error {
	return c.Request.ParseMultipartForm(maxMemory)
}

/**
 * 获取表单文件
 */
func (c *MuxContext) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return c.Request.FormFile(key)
}

/**
 * 输出string
 */
func (c *MuxContext) WriteString(data interface{}) {
	bytejson, _ := json.Marshal(data)
	io.WriteString(c.Writer, string(bytejson))
}

/**
 * 使用默认返回格式输出
 */
func (c *MuxContext) WritedString(sucess bool, code string, msg string, data interface{}) {
	result := NewResultInstance()
	if sucess {
		result.Sucess(msg, code, data)
	} else {
		result.Fail(msg, code, data)
	}
	bytejson, _ := json.Marshal(result)
	io.WriteString(c.Writer, string(bytejson))
}
