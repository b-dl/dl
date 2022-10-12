package router

import (
	"net/http"

	"github.com/b-dl/dl/tool"
)

func Ping(rw http.ResponseWriter, r *http.Request) {
	rw.Write(tool.Str2Bytes(`{"code":0,"msg":"服务正常"}`))
}

func Token(rw http.ResponseWriter, r *http.Request) {
	rw.Write(nil)
}
