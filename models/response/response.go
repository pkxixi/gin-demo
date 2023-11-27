package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	Success = 10000
	Failure = 20000
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(data interface{}, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OKWithData(data interface{}, c *gin.Context) {
	Result(Success, data, "success", c)
}

func OKWithMsg(msg string, c *gin.Context) {
	Result(Success, map[string]interface{}{}, msg, c)
}

func Fail(data interface{}, msg string, c *gin.Context) {
	Result(Failure, data, msg, c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(Failure, map[string]interface{}{}, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[ErrorCode(code)]
	if ok {
		Result(int(code), map[string]interface{}{}, msg, c)
	}
	Result(Failure, map[string]interface{}{}, "unknown error", c)
}
