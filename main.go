package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/britannic/lgtv-remote/internal/lgtv"
	logging "github.com/op/go-logging"
	"github.com/pkg/term"
	"github.com/pkg/term/termios"
)

var (
	prog    = path.Base(os.Args[0])
	log     = logging.MustGetLogger(prog)
	logFile = prog + ".log"
)

func main() {
	fmt.Println(lgtv.Cmd.GetRespMap())

	os.Exit(0)

	err := newLog()
	if err != nil {
		log.Errorf("Unable to open log file: %v, error: %v\n", logFile, err)
	}

	ts, err := term.Open("/dev/ptyp0", term.Speed(9600))
	if err != nil {
		log.Fatalf("term.Open error: %v", err)
	}
	defer ts.Close()

	tm, err := term.Open("/dev/ttyp0", term.Speed(9600))
	if err != nil {
		log.Fatalf("term.Open error: %v", err)
	}
	defer tm.Close()

	if err = ts.SetOption(term.Speed(9600)); err != nil {
		log.Errorf("slave.SetOption error: %v", err)
	}

	if err = ts.SetReadTimeout(1); err != nil {
		log.Errorf("slave.SetReadTimeout error: %v", err)
	}

	if err = tm.SetReadTimeout(1); err != nil {
		log.Errorf("master.SetReadTimeout error: %v", err)
	}

	if err = ts.SetRaw(); err != nil {
		log.Errorf("slave.SetRaw error: %v", err)
	}

	if err = tm.SetRaw(); err != nil {
		log.Errorf("master.SetRaw error: %v", err)
	}

	b := []byte("This is a test and it is really massive!\nLorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.\n")

	if _, err = ts.Write(b); err != nil {
		log.Errorf("ts.Write error: %v", err)
	}

	i, err := ts.Buffered()
	if err != nil {
		log.Errorf("ts.Buffered error: %v", err)
	}

	log.Infof("There are %d bytes buffered for ts", i)

	i, err = tm.Buffered()
	if err != nil {
		log.Errorf("tm.Buffered error: %v", err)
	}

	log.Infof("There are %d bytes buffered for tm", i)

	n := 0
	read := 0

	for n != 50 {
		buf := make([]byte, 32)
		n++
		i, err = tm.Read(buf)
		read = read + i
		time.Sleep(time.Millisecond * 50)
		switch {
		case err != nil:
			log.Errorf("tm.Read error: %v", err)
		case i > 0:
			if _, err = ts.Write([]byte(fmt.Sprintf("%d times", n))); err != nil {
				log.Errorf("ts.Write error: %v", err)
			}
			break
		}

		log.Infof("tm.Read() = %s\n", buf[:i])
	}

	log.Infof("There were %d bytes read from tm", read)

	// tv := lgtv.API{Logger: log, Timeout: 0}
	// _ = lgtv.Cmd

	// tv.Critical("Test")
	// tv.ShowPIN()
	// if !tv.Zap(cmd["Power"]) {
	// 	log.Error("Unable to contact TV...")
	// }

	// fmt.Println(lgtv.Cmd)

}

func newLog() error {
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

	return err
}

func opendev() (*term.Term, string) {
	_, pts, err := termios.Pty()
	if err != nil {
		log.Fatalf("termios.Pty() error: %v", err)
	}
	defer pts.Close()

	slave, err := term.Open(pts.Name())
	if err != nil {
		log.Fatalf("slave.Open() error %v", err)
	}

	return slave, pts.Name()
}
