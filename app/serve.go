package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/b-dl/dl/logger"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func serveAction(c *cli.Context) error {
	start := time.Now()
	cli.VersionPrinter(c)
	logger.Init(path.Join("log"), c.String("level"))

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", c.Uint("port")),
		Handler: http.DefaultServeMux,
	}

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
	logrus.Infof("Service listening port number: %d", c.Uint("port"))
	logrus.Infof("Service startup time: %d μs", time.Since(start).Microseconds())

	err = srv.Serve(ln)
	if err != nil && err != http.ErrServerClosed {
		logrus.Errorf("Service startup error: %# v", err)
	} else {
		logrus.Debugf("Successfully shut down the service: %# v", err)
	}

	return nil
}
