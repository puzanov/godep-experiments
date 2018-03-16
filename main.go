package main

import (
	"os"

	"github.com/google/logger"
	"github.com/op/go-logging"
)

func main() {
	// google logger
	logger.Info("I'm about to do something!")

	// go loggong pkg
	log := logging.MustGetLogger("example")
	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.ERROR, "")
	logging.SetBackend(backendLeveled, backendFormatter)
	log.Debugf("debug %s", "secret")
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("err")
	log.Critical("crit")
}
