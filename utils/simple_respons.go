package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespJson struct {
	Code    int         `json:"code"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

func JSONR(c *gin.Context, arg ...interface{}) (werror error) {
	var (
		wcode int
		msg   interface{}
		data  interface{}
	)
	if len(arg) == 3 {
		wcode = arg[0].(int)
		msg = arg[1]
		data = arg[2]
	} else {
		wcode = arg[0].(int)
		msg = arg[1]
	}
	var body interface{}
	if wcode == 200 {
		switch msg.(type) {
		case string:
			body = RespJson{Code: wcode, Message: msg.(string), Data: data}
			c.JSON(http.StatusOK, body)
		default:
			c.JSON(http.StatusOK, msg)
			body = msg
		}
	} else {
		switch msg.(type) {
		case string:
			body = RespJson{Code: wcode, Message: msg.(string)}
			c.JSON(wcode, body)
		case error:
			body = RespJson{Code: wcode, Message: msg.(error).Error()}
			c.JSON(wcode, body)
		default:
			body = RespJson{Code: wcode, Message: "system type error. please ask admin for help"}
			c.JSON(wcode, body)
		}
	}
	return
}
