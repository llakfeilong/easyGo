package EasyGo

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type Validator struct {
	ErrorResults []ExceptionResult `json:"results"`
	ValidResult  bool
}

type ExceptionResult struct {
	FiledName      string      `json:"filedName"`
	FailureMessage string      `json:"failureMessage"`
	FiledValue     interface{} `json:"filedValue"`
}

func NewValid() *Validator {
	return &Validator{
		ErrorResults: make([]ExceptionResult, 0),
	}
}

/**
 * 解析参数
 */
func (valid *Validator) VaildSturct(parameter interface{}) *Validator {
	v := reflect.ValueOf(parameter)
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name     //参数名称
		fieldType := v.Field(i).Type()          //参数类型
		fieldValue := v.Field(i).Interface()    //参数值
		fvalue := fmt.Sprintf("%v", fieldValue) //格式化
		var result ExceptionResult
		for _, key := range ValidType {
			fieldTag := v.Type().Field(i).Tag.Get(key)
			if fieldTag == "" {
				continue
			}

			switch key {
			case VALIDLENGTH, VALIDNOTNULL:
				fieldTagslice := strings.Split(fieldTag, ";")
				var validvalue, defualtmsg string
				for _, value := range fieldTagslice {
					index := strings.Index(value, ":")
					fieldTagkey := value[:index]
					fieldTagvalue := value[index+1:]
					if fieldTagkey == VALIDVALUE {
						//key等于校验值标志
						validvalue = fieldTagvalue
					} else if fieldTagkey == VALIDDEFAULTMSG {
						//KEY等于默认错误消息
						defualtmsg = fieldTagvalue
					}
				}
				if key == VALIDLENGTH {
					b, field, msg, filedvalue := valid.validStructLength(validvalue, fieldType.String(), fvalue, fieldName, defualtmsg)
					if !b {
						result.FiledName = field
						result.FailureMessage = msg
						result.FiledValue = filedvalue
						valid.ErrorResults = append(valid.ErrorResults, result)
						valid.ValidResult = false
					} else {
						valid.ValidResult = true
					}
				} else if key == VALIDNOTNULL {
					b, field, msg, filedvalue := valid.validStructNotNull(fieldType.String(), fvalue, fieldName, defualtmsg)
					if !b {
						result.FiledName = field
						result.FailureMessage = msg
						result.FiledValue = filedvalue
						valid.ErrorResults = append(valid.ErrorResults, result)
						valid.ValidResult = false
					} else {
						valid.ValidResult = true
					}
				}

			}

		}
	}
	return valid
}

//校验单个参数是否为空,只支持int string
func (valid *Validator) ValidNotNull(parameter interface{}) (bool, string) {
	v := reflect.ValueOf(parameter)
	value := v.Field(0).Interface().(string) //强制转换类型
	switch t := parameter.(type) {
	case int:
		if value == "" {
			return false, "the parameter is Empty"
		} else {
			return true, "the parameter is not Empty"
		}
	case string:
		if value == "" {
			return false, "the parameter is Empty"
		} else {
			return true, "the parameter is not Empty"
		}
	default:
		_ = t
		return false, "This function only supports int and string types"
	}
}

//校验长度
//func (valid *Validator) ValidLength(parameter interface{})(bool,string){
//
//}

//校验结构体参数是否空值
func (valid *Validator) validStructNotNull(fieldType, fieldValue, fieldName, defaultmsg string) (b bool, field, message, fieldvalue string) {
	if fieldValue == "" {
		return false, fieldName, defaultmsg, fieldValue
	} else if fieldType != "int" && fieldType != "string" {
		return false, fieldName, "参数类型不支持校验长度", fieldValue
	}
	return true, fieldName, "校验成功", fieldValue
}

/**
 *  校验结构体参数长度
 */
func (valid *Validator) validStructLength(fieldTagValue, fieldType, fieldValue, fieldName, defaultmsg string) (b bool, field, message, fieldvalue string) {
	if fieldTagValue == "" {
		return false, fieldName, "参数不能为空", fieldValue
	} else if fieldType != "int" && fieldType != "string" {
		return false, fieldName, "参数类型不支持校验长度", fieldValue
	}
	index := strings.Index(fieldTagValue, "-")
	length := len(fieldValue)
	log.Println("字符串长度:", length)
	if index != -1 {
		//找到范围符号
		minstr := fieldTagValue[:index]
		maxstr := fieldTagValue[index+1:]
		min, err := strconv.Atoi(minstr)
		if err != nil {
			return false, fieldName, "最小范围长度必须是int类型", fieldValue
		}
		max, err := strconv.Atoi(maxstr)
		if err != nil {
			return false, fieldName, "最大范围长度必须是int类型", fieldValue
		}
		if length >= min && length <= max {
			return true, fieldName, "校验成功", fieldValue
		} else {
			return false, fieldName, "超过范围长度", fieldValue
		}

	} else {
		max, err := strconv.Atoi(fieldTagValue)
		log.Println("字符串限制长度", max)
		if err != nil {
			return false, fieldName, "最大长度必须是int类型", fieldValue
		}
		if length > max {
			return false, fieldName, defaultmsg, fieldValue
		} else {
			return true, fieldName, "校验成功", fieldValue
		}
	}
}
