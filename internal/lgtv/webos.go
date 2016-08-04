package lgtv

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	logging "github.com/op/go-logging"
)

const (
	agent   = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/601.7.7 (KHTML, like Gecko) Version/9.1.2 Safari/601.7.7`
	cr      = "\r\n"
	httpStr = `http://`
	udp4    = "udp4"
)

// WebOS struct for the LG TV WebOS API interface
type WebOS struct {
	*logging.Logger
	AppID   string
	AppName string
	Found   bool
	ID      string
	IP      net.IP
	Name    string
	Pin     string
	Timeout time.Duration
}

var (
	conn     *net.UDPConn
	maxTries = 10
	mode     = CmdMode{Pair: "/udap/api/pairing", Send: "/udap/api/command"}
	sock     = false
)

func (w *WebOS) chkMsgs() (bool, error) {
	var (
		buf [1024]byte
		err error
		ok  bool
	)

	n, addr, _ := conn.ReadFromUDP(buf[0:])
	ip, err := w.getLocalIP()
	if n > 0 && addr.IP.String() != ip {
		msg := string(buf[0:n])
		ok, err = w.parseMsg(msg, addr)
	}

	return ok, err
}

func (w *WebOS) getLocalIP() (string, error) {
	var s string

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return s, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return s, errors.New("unable to detect a connected ethernet interface")
}

// Pair using the LG TV's PIN
func (w *WebOS) Pair() {
	msg := []byte(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?><envelope><api type="pairing"><name>hello</name><value>%v</value><port>8080</port></api></envelope>`, w.Pin))

	w.Infof("Pairing with TV: %v using Pin: %v", w.Name, w.Pin)

	w.Send(mode.Pair, msg)
}

func (w *WebOS) scan(portAddr string, msg []byte) error {
	var timeout time.Time
	timeout.Add(7 * time.Second)
	conn.SetWriteDeadline(timeout)

	udpAddr, err := net.ResolveUDPAddr(udp4, fmt.Sprintf("%v:%v", net.IPv4bcast.String(), portAddr))
	if err != nil {
		w.Fatal(err)
	}

	w.Infof("Broadcasting %q on: %v:%v", msg, net.IPv4bcast.String(), portAddr)

	_, err = conn.WriteToUDP(msg, udpAddr)
	defer conn.Close()

	return err
}

func (w *WebOS) pairingRequest() error {
	xmitStr := []byte(`<?xml version="1.0" encoding="utf-8"?><envelope><api type="pairing"><name>showKey</name></api></envelope>`)

	if code, _, err := w.Send(mode.Pair, xmitStr); err != nil || code != 200 {
		return fmt.Errorf("Pairing error: %v", err)
	}

	w.Info("Pairing successful")
	return nil
}

func (w *WebOS) parseMsg(msg string, addr *net.UDPAddr) (bool, error) {
	if msg == "" {
		return false, fmt.Errorf("message cannot be empty")
	}

	if addr.Port == 1990 {
		rx := regexp.MustCompile(`SERVER: [\w//.]* [\w//.]* ([\w-]*)`)
		w.Found = true
		w.IP = addr.IP
		w.Name = rx.FindStringSubmatch(msg)[1]
		w.Infof("LG TV %v with IP %v responded", w.Name, w.IP.String())
		w.pairingRequest()
	}

	if addr.IP.String() == w.IP.String() && w.Found == true {
		w.Infof("LG TV %v with IP %v says: %q", w.Name, addr.IP, msg)
	}

	return true, nil
}

// Send xmits a WebOS request to the LG TV.
func (w *WebOS) Send(cmd string, msg []byte) (int, io.Reader, error) {
	var (
		body    []byte
		err     error
		lgtvCMD = fmt.Sprintf("%v%v:8080%v%v", httpStr, w.IP.String(), cmd, bytes.NewReader(msg))
		resp    *http.Response
		req     *http.Request
	)

	w.Infof("About to contact LG TV on address: %s with command: %s", lgtvCMD, string(msg))

	if req, err = http.NewRequest("POST", lgtvCMD, nil); err != nil {
		return http.StatusNotAcceptable, strings.NewReader(fmt.Sprintf("Unable to form HTTP request %v", lgtvCMD)), err
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Content-Length", strconv.Itoa(len(msg)))
	req.Header.Add("Connection", "Close")
	req.Header.Add("User-Agent", agent)

	if resp, err = (&http.Client{}).Do(req); err != nil {
		return resp.StatusCode, strings.NewReader(fmt.Sprintf("Unable to get response from %v", w.IP.String())), err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if len(body) < 1 {
		return resp.StatusCode, strings.NewReader(fmt.Sprintf("%s did not confirm command received", w.IP.String())), err

	}

	return resp.StatusCode, bytes.NewBuffer(body), err
}

func (w *WebOS) setUpSox() {
	ip, err := w.getLocalIP()
	if err != nil {
		w.Fatal(err)
	}

	w.Infof("Found IP: %v", ip)

	udpAddr, err := net.ResolveUDPAddr(udp4, ":1990")
	if err != nil {
		fmt.Println("Resolve:", err)
		os.Exit(1)
	}
	conn, err = net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Listen:", err)
		os.Exit(1)
	}

	sock = true
}

// ShowPIN displays the LG TV's PIN (Pairing ID Number) on its screen.
func (w *WebOS) ShowPIN() {
	if !sock {
		w.setUpSox()
	}

	_ = []byte(`M-SEARCH * HTTP/1.1` + cr +
		`HOST: 239.255.255.250:1900` + cr +
		`MAN: "ssdp:discover"` + cr +
		`MX: 2` + cr +
		`ST: urn:schemas-upnp-org:device:MediaRenderer:1` + cr + cr)

	xmitStr := []byte(`B-SEARCH * HTTP/1.1` + cr +
		`HOST: 239.255.255.250:1990` + cr +
		`MAN: "ssdp:discover` + cr + `MX: 3` + cr +
		`ST: urn:schemas-udap:service:smartText:1` + cr +
		`USER-AGENT:` + agent + cr + cr)

	w.scan("1990", xmitStr)

	i := 0
	for !w.Found && i != maxTries {
		w.chkMsgs()
		i++
		switch {
		case !w.Found && i != maxTries:
			w.Warning("No LG TV detected yet...")
			time.Sleep(w.Timeout * time.Second)
		case !w.Found && i == maxTries:
			w.Critical("No LG TV detected, giving up!")
		case w.Found:
			w.Infof("LG TV %v with IDL %v found at %v", w.Name, w.ID, w.IP)
		}
	}
}

// Zap xmits a WebOS command.
func (w *WebOS) Zap(cmd int) bool {
	i := strconv.Itoa(cmd)
	xmitStr := []byte(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?><envelope><api type="command"><name>HandleKeyInput</name><value>%v</value></api></envelope>`, i))
	zap := []byte(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?><envelope><api type="command"><name>HandleKeyInput</name><value>%v</value></api></envelope>`, i))

	w.Infof("Sending command %v to %v", i, w.Name)

	resp, _, _ := w.Send(mode.Send, zap)

	// Pairing required after the LG TV has been turned off
	if resp != 200 {
		w.Pair()
		resp, _, _ = w.Send(mode.Send, xmitStr)
	}

	return resp == 200
}
