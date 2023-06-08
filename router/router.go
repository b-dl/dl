package router

import (
	"net/http"

	"github.com/b-dl/dl/logger"
	"github.com/sirupsen/logrus"
)

func Ping(rw http.ResponseWriter, r *http.Request) {
	logger.Request(r)
	rw.WriteHeader(http.StatusOK)
}

func Token(rw http.ResponseWriter, r *http.Request) {
	logger.Request(r)
	authorization := r.Header.Get("Authorization")
	logrus.Debug(authorization)
}
