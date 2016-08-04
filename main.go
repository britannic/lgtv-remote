package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/britannic/lgtv-remote/internal/lgtv"
	"github.com/tarm/serial"
)

// trap SIGINT and exit if received
func sigExit(i int) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		os.Exit(i)
	}()
}

func main() {
	sigExit(1)
	port := flag.String("port", "/dev/ttys000", "set serial device")
	flag.Parse()

	s := lgtv.Serial{
		Baud:        9600,
		Cmd:         lgtv.Cmd.SetSerialCmds(),
		Parity:      serial.ParityNone,
		Port:        *port,
		ReadTimeout: 1 * time.Second,
	}

	tty, err := s.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()
}
