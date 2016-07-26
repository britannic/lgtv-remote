package main

import (
	"os"
	"path"

	"github.com/britannic/lgtv-remote/internal/lgtv"
	logging "github.com/op/go-logging"
)

var (
	prog     = path.Base(os.Args[0])
	log, err = newLog()
	logFile  = prog + ".log"
)

func main() {
	if log, err = newLog(); err != nil {
		log.Errorf("Unable to open log file: %v, error: %v\n", logFile, err)
	}

	tv := lgtv.API{Logger: log}
	_ = lgtv.Cmd

	tv.ShowPIN()
	// if !tv.Zap(cmd["Power"]) {
	// 	log.Error("Unabled to contact TV...")
	// }

}

func newLog() (*logging.Logger, error) {
	fdFmt := logging.MustStringFormatter(
		`%{level:.4s}[%{id:03x}]%{time:2006-01-02 15:04:05.000} ▶ %{message}`,
	)

	scrFmt := logging.MustStringFormatter(
		`%{color:bold}%{level:.4s}%{color:reset}[%{id:03x}]%{time:15:04:05.000} ▶ %{message}`,
	)

	fd, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	fdlog := logging.NewLogBackend(fd, "", 0)
	fdFmttr := logging.NewBackendFormatter(fdlog, fdFmt)

	scr := logging.NewLogBackend(os.Stderr, "", 0)
	scrFmttr := logging.NewBackendFormatter(scr, scrFmt)

	logging.SetBackend(fdFmttr, scrFmttr)

	return logging.MustGetLogger(prog), err
}
