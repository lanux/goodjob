package handler

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/i18n"
)

type ResponseBasic struct {
	Code    int         `json:"code" xml:"code"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
}

func Success(data interface{}, msg string) *ResponseBasic {
	return &ResponseBasic{
		Code:    200,
		Message: msg,
		Data:    data,
	}
}

func Error(code int, msg string) *ResponseBasic {
	return &ResponseBasic{
		Code:    code,
		Message: msg,
	}
}

func ErrorWithLocale(code int, format string, ctx iris.Context) *ResponseBasic {
	return &ResponseBasic{
		Code:    code,
		Message: i18n.Translate(ctx, format),
	}
}
