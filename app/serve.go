package app

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/b-dl/dl/logger"
	"github.com/b-dl/dl/request"
	"github.com/b-dl/dl/router"
	"github.com/b-dl/dl/ws"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const port = "10101"

func serveAction(c *cli.Context) error {
	start := time.Now()
	cli.VersionPrinter(c)
	logger.Init(path.Join("log"), c.String("level"))

	request.SetOptions(request.RequestOptions{
		RetryTimes: int(c.Uint("retry-times")),
		Timeout:    int64(c.Uint("timeout")),
	})

	hub := ws.NewHub()
	go hub.Run()

	srv := http.Server{
		Addr:    ":" + port,
		Handler: http.DefaultServeMux,
	}

	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		ws.Ws(hub, rw, r)
	})

	http.HandleFunc("/ping", router.Ping)
	http.HandleFunc("/token", router.Token)

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		data := <-sigint
		logrus.Infof("%s Close signal", data.String())
		logrus.Debug("Start shutting down the service......")
		if err := srv.Shutdown(context.Background()); err != nil {
			logrus.Errorf("Service shutdown failed: %# v", err)
		}
		os.Exit(1)
	}()

	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}
	logrus.Infof("Service listening port number: %s", port)
	logrus.Infof("Service startup time: %.3f s", time.Since(start).Abs().Seconds())

	err = srv.Serve(ln)
	if err != nil && err != http.ErrServerClosed {
		logrus.Errorf("Service startup error: %# v", err)
	} else {
		logrus.Debugf("Successfully shut down the service: %# v", err)
	}

	return nil
}
