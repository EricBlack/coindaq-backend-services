package utils

import (
	"reflect"
	"strings"
	"errors"
	"strconv"
	"regexp"
)

const TagName = "validate"

var MailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

type Validator struct {}

var Validate = newValidator()

func newValidator() *Validator {
	return &Validator{}
}

func(this *Validator) ValidateData(data interface{}) error {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	//普通值类型不解析，直接返回
	if v.Kind() != reflect.Slice && v.Kind() !=reflect.Ptr && v.Kind() != reflect.Struct {
		return nil
	}

	err := this.parseParam(t, v, reflect.StructTag(""))
	return err
}

func (this *Validator) parseParam(t reflect.Type, v reflect.Value, tag reflect.StructTag) error {
	//判断是否是数组
	if v.Kind() == reflect.Slice {
		for i, n:=0, v.Len(); i<n; i++ {
			t1 := reflect.TypeOf(v.Index(i).Interface())
			err := this.parseParam(t1, v.Index(i), tag)
			if err != nil {
				return err
			}
		}
		//是否是指针
	} else if v.Kind() ==reflect.Ptr {
		t = t.Elem()
		v =v.Elem()
		err := this.parseParam(t, v, tag)
		if err != nil {
			return err
		}
		//是否是结构体
	} else if v.Kind() == reflect.Struct{
		for i :=0; i<t.NumField(); i++ {
			f := t.Field(i)
			t.FieldByName(f.Name)

			err := this.parseParam(f.Type, v.FieldByName(f.Name), f.Tag)
			if err != nil {
				return err
			}
		}
		//解析普通类型
	} else {
		invoker := reflect.ValueOf(this)
		methodsName := strings.Split(tag.Get("valid"), ",")
		for i, n :=0, len(methodsName); i<n; i++ {
			methodName := methodsName[i]
			if len(methodName) !=0 {
				method := invoker.MethodByName("Check" + methodName)
				inVal := []reflect.Value{v, reflect.ValueOf(tag)}
				outVal := method.Call(inVal)
				outValLen := len(outVal)
				if outValLen != 1 {
					v.Set(outVal[0])
				}
				if !outVal[outValLen-1].IsNil() {
					return outVal[outValLen-1].Elem().Interface().(error) //返回错误信息
				}
			}
		}
	}

	return nil
}

//检查字符串
func (this *Validator) CheckString(data string, tag reflect.StructTag) (string, error){
	data = strings.Replace(data, "'", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data = strings.Replace(data, "\\", "", -1)
	data = strings.Replace(data, "\"", "", -1)

	var minLen, maxLen int = 0, 0

	var err error = nil

	lenStr := strings.Split(tag.Get("len"), ",")

	if lenStr[0] != "" {
		minLen, err = strconv.Atoi(lenStr[0])
		if err != nil {
			return "", errors.New(tag.Get("name") + ":验证字符串的最小长度参数输入有误!")
		}
	}
	if len(lenStr) == 2 && lenStr[1] != "" {
		maxLen, err = strconv.Atoi(lenStr[1])
		if err != nil {
			return "", errors.New(tag.Get("name") + ":验证字符串的最大长度参数输入有误!")
		}
	}

	if minLen > maxLen {
		return "", errors.New(tag.Get("name") + ":最小长度和最大长度冲突!")
	}
	dataLen := len(data)
	if dataLen < minLen {
		return data, errors.New(tag.Get("name") + ":字符串长度过短!")
	}

	if maxLen != 0 && dataLen > maxLen {
		return data, errors.New(tag.Get("name") + ":字符串过长!")
	}

	return data, nil
}

//检查正负数
func (this *Validator) CheckPosNO(data int, tag reflect.StructTag) error {
	if data <0 {
		return errors.New(tag.Get("name") + ":不能为负数")
	}

	return nil
}

//邮件检测
func (this *Validator) CheckEmail (data string, tag reflect.StructTag) error {
	if !MailRe.MatchString(data) {
		return errors.New(tag.Get("name") + ":邮件格式不正确！")
	}
	return nil
}


