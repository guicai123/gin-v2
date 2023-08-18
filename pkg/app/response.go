package app

import (
	"github.com/gin-gonic/gin"

	"github.com/guicai123/gin-v2/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

// Response setting gin.JSON
func (g *Gin) Nresponse(Code int, Msg string, data interface{}) {
	g.C.JSON(200, Response{
		Code: Code,
		Msg:  Msg,
		Data: data,
	})
	return
}
