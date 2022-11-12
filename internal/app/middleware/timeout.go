package middleware

import (
	"net/http"
	"time"

	// external packages
	"github.com/gin-gonic/gin"
	timeout "github.com/vearne/gin-timeout"
	// project packages
)

var (
	defaultTimeOutMsg = `{"code": -1, "msg":"http: Handler timeout"}`
	timeOutSecond     = time.Millisecond * 1500 // 1500ms
)

func TimeOutHandler() gin.HandlerFunc {
	return timeout.Timeout(
		timeout.WithTimeout(timeOutSecond),
		timeout.WithErrorHttpCode(http.StatusRequestTimeout),
		timeout.WithDefaultMsg(defaultTimeOutMsg),
	)
}
