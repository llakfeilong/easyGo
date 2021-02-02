package errorType

const (
	UPLOADFILE_NULL_ERROR      = "23301" //上传文件为空
	UPLOADFILE_SIZE_ERROR      = "23302" //文件大小超过限制
	UNSUPPORTED_FILETYPE_ERROR = "23303" //不支持的文件类型
	UNFIND_FILEKEY_ERROR       = "23304" //找不到文件秘钥
	UNFIND_FILE_ERROR          = "23305" //找不到该文件
	UNFIND_FOLODER             = "23306" //未创建该文件夹
	NOTVALID_ERROR             = "10001" //接口参数校验错误
	DUPLICATENAME              = "23307" //重复的文件夹名称
	SAVEFILEFAIL               = "23308" //保存文件失败系统异常
	REQUESTMETHODFAIL          = "23309" //请求方法错误
	JSON_PARSE_ERROR           = "10004" //json解析错误
	SYSTEM_ERR                 = "-1"    //系统异常
	UNFIND_FILETYPE_ID         = "23310" //不存在的文件类型ID
)

var BussinessErrorType = map[string]string{
	UPLOADFILE_NULL_ERROR:      "上传文件为空",
	UPLOADFILE_SIZE_ERROR:      "文件大小超过限制",
	UNSUPPORTED_FILETYPE_ERROR: "不支持的文件类型",
	UNFIND_FILEKEY_ERROR:       "找不到文件秘钥",
	UNFIND_FILE_ERROR:          "找不到该文件",
	UNFIND_FOLODER:             "未创建该文件夹",
	NOTVALID_ERROR:             "接口参数校验错误",
	DUPLICATENAME:              "重复的文件夹名称",
	SAVEFILEFAIL:               "保存文件失败，系统异常",
	REQUESTMETHODFAIL:          "请求方法不正确",
	JSON_PARSE_ERROR:           "json解析错误",
	SYSTEM_ERR:                 "系统异常",
	UNFIND_FILETYPE_ID:         "不存在的文件类型ID",
}
