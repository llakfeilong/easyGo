package httpMux

//校验枚举
const (
	VALID           = "Valid"      //校验标志
	VALIDLENGTH     = "@Length"    //长度
	VALIDNOTNULL    = "@NotNull"   //不为空
	VALIDDEFAULTMSG = "DefaultMsg" //默认消息
	VALIDVALUE      = "Value"      //校验值
	VALIDSIZE       = "@Size"
)

var ValidType = []string{
	VALIDLENGTH, VALIDNOTNULL, VALIDSIZE,
}
