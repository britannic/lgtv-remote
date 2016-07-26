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
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context"

	logging "github.com/op/go-logging"
)

// API struct for the LG TV API interface
type API struct {
	*logging.Logger
	AppID   string
	AppName string
	ctx     context.Context
	Found   bool
	ID      string
	IP      net.IP
	Name    string
	Pin     string
	Timeout time.Duration
}

// CmdMode sets which API command is used
type CmdMode struct {
	Pair string
	Send string
}

// TVCmds is a map[string]int map of commands
type TVCmds map[string]int

const (
	agent   = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/601.7.7 (KHTML, like Gecko) Version/9.1.2 Safari/601.7.7`
	cr      = "\r\n"
	httpStr = `http://`
	udp4    = "udp4"
)

var (
	// Cmd maps LG TV int commands to meaningful names
	Cmd = TVCmds{
		"Power":       1,
		"Num0":        2,
		"Num1":        3,
		"Num2":        4,
		"Num3":        5,
		"Num4":        6,
		"Num5":        7,
		"Num6":        8,
		"Num7":        9,
		"Num8":        10,
		"Num9":        11,
		"Up":          12,
		"Down":        13,
		"Left":        14,
		"Right":       15,
		"OK":          20,
		"Home":        21,
		"Menu":        22,
		"Back":        23,
		"Vol_Up":      24,
		"Vol_Dn":      25,
		"Mute":        26,
		"Ch_Up":       27,
		"Ch_Dn":       28,
		"Blue":        29,
		"Green":       30,
		"Red":         31,
		"Yellow":      32,
		"Play":        33,
		"Pause":       34,
		"Stop":        35,
		"FF":          36,
		"REW":         37,
		"Skip_FF":     38,
		"Skip_REW":    39,
		"REC":         40,
		"REC_List":    41,
		"Repeat":      42,
		"Live":        43,
		"EPG":         44,
		"Info":        45,
		"Aspect":      46,
		"Ext":         47,
		"PIP":         48,
		"Subtitle":    49,
		"Prog_List":   50,
		"Text":        51,
		"Mark":        52,
		"3D":          400,
		"3D_LR":       401,
		"Dash":        402,
		"Prev_Ch":     403,
		"Fave":        404,
		"Quick_Menu":  405,
		"Text_Opt":    406,
		"Audio_Desc":  407,
		"Netcast":     408,
		"Energy_Save": 409,
		"AV":          410,
		"SimpLink":    411,
		"Exit":        412,
		"Reserve":     413,
		"PIP_CH_Up":   414,
		"PIP_CH_Down": 415,
		"PIP_Switch":  416,
		"Apps":        417,
	}

	conn     *net.UDPConn // UDP Connection
	maxTries = 10

	mode = CmdMode{
		Pair: "/udap/api/pairing",
		Send: "/udap/api/command",
	}

	sock = false
)

// Send sends an http request to the LG smart tv specified by *TV
func (a *API) Send(cmd string, msg []byte) (int, io.Reader, error) {
	var (
		body    []byte
		err     error
		lgtvCMD = fmt.Sprintf("%v%v:8080%v%v", httpStr, a.IP.String(), cmd, bytes.NewReader(msg))
		resp    *http.Response
		req     *http.Request
	)

	a.Info("About to contact LG TV on address: ", lgtvCMD, " with command: ", string(msg))

	if req, err = http.NewRequest("POST", lgtvCMD, nil); err != nil {
		return http.StatusNotAcceptable, strings.NewReader(fmt.Sprintf("Unable to form HTTP request %v", lgtvCMD)), err
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Content-Length", strconv.Itoa(len(msg)))
	req.Header.Add("Connection", "Close")
	req.Header.Add("User-Agent", agent)

	if resp, err = (&http.Client{}).Do(req); err != nil {
		return resp.StatusCode, strings.NewReader(fmt.Sprintf("Unable to get response from %v", a.IP.String())), err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if len(body) < 1 {
		return resp.StatusCode, strings.NewReader(fmt.Sprintf("%s did not confirm command received", a.IP.String())), err

	}

	return resp.StatusCode, bytes.NewBuffer(body), err
}

// ShowPIN displays PIN on screen for pairing
func (a *API) ShowPIN() {
	if !sock {
		a.setUpSox()
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

	a.scan("1990", xmitStr)

	i := 0
	for !a.Found && i != maxTries {
		a.chkMsgs()
		i++
		switch {
		case !a.Found && i != maxTries:
			a.Warning("No LG TV detected yet...")
			time.Sleep(a.Timeout * time.Second)
		case !a.Found && i == maxTries:
			a.Critical("No LG TV detected, giving up!")
		case a.Found:
			a.Infof("LG TV %v with IDL %v found at %v", a.Name, a.ID, a.IP)
		}
	}
}

// Pair using a specified pin
func (a *API) Pair() {
	msg := []byte(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?><envelope><api type="pairing"><name>hello</name><value>%v</value><port>8080</port></api></envelope>`, a.Pin))

	a.Infof("Pairing with TV: %v using Pin: %v", a.Name, a.Pin)

	a.Send(mode.Pair, msg)
}

// Zap sends a command to the tv
func (a *API) Zap(cmd int) bool {
	i := strconv.Itoa(cmd)
	xmitStr := []byte(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?><envelope><api type="command"><name>HandleKeyInput</name><value>%v</value></api></envelope>`, i))
	zap := []byte(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?><envelope><api type="command"><name>HandleKeyInput</name><value>%v</value></api></envelope>`, i))

	a.Infof("Sending command %v to %v", i, a.Name)

	resp, _, _ := a.Send(mode.Send, zap)

	// Pairing required whenever the LG has been turned off and then on
	if resp != 200 {
		a.Pair()
		resp, _, _ = a.Send(mode.Send, xmitStr)
	}

	return resp == 200
}

func (a *API) scan(portAddr string, msg []byte) error {
	var timeout time.Time
	timeout.Add(7 * time.Second)
	conn.SetWriteDeadline(timeout)

	udpAddr, err := net.ResolveUDPAddr(udp4, fmt.Sprintf("%v:%v", net.IPv4bcast.String(), portAddr))
	if err != nil {
		a.Fatal(err)
	}

	a.Infof("Broadcasting %q on: %v:%v", msg, net.IPv4bcast.String(), portAddr)

	_, err = conn.WriteToUDP(msg, udpAddr)
	defer conn.Close()

	return err
}

func (a *API) getLocalIP() (string, error) {
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

func (a *API) setUpSox() {
	ip, err := a.getLocalIP()
	if err != nil {
		a.Fatal(err)
	}

	a.Infof("Found IP: %v", ip)

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

func (a *API) chkMsgs() (bool, error) {
	var (
		buf [1024]byte
		err error
		ok  bool
	)

	n, addr, _ := conn.ReadFromUDP(buf[0:])
	ip, err := a.getLocalIP()
	if n > 0 && addr.IP.String() != ip {
		msg := string(buf[0:n])
		ok, err = a.parseMsg(msg, addr)
	}

	return ok, err
}

// parseMsg parses a message found by CheckForMessages
func (a *API) parseMsg(msg string, addr *net.UDPAddr) (bool, error) {
	if msg == "" {
		panic("empty message")
	}

	if addr.Port == 1990 {
		rx := regexp.MustCompile(`SERVER: [\w//.]* [\w//.]* ([\w-]*)`)

		a.Found = true
		a.IP = addr.IP
		a.Name = rx.FindStringSubmatch(msg)[1]
		a.Infof("LG TV %v with IP %v responded", a.Name, a.IP.String())
		a.pairingRequest()
	}

	if addr.IP.String() == a.IP.String() && a.Found == true {
		a.Infof("LG TV %v with IP %v says: %q", a.Name, addr.IP, msg)
	}

	return true, nil
}

func (a *API) pairingRequest() error {
	xmitStr := []byte(`<?xml version="1.0" encoding="utf-8"?><envelope><api type="pairing"><name>showKey</name></api></envelope>`)

	if code, _, err := a.Send(mode.Pair, xmitStr); err != nil || code != 200 {
		return fmt.Errorf("Pairing error: %v", err)
	}

	a.Info("Pairing successful")

	return nil
}

// Pair holds key/value pairs
type Pair struct {
	Key   string
	Value int
}

// PairList is a slice of pairs that implements sort.Interface to sort by values
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func (tv TVCmds) String() (s string) {
	maxLen := func(keys PairList) int {
		smallest := len(keys[0].Key)
		largest := len(keys[0].Key)
		for i := range keys {
			if len(keys[i].Key) > largest {
				largest = len(keys[i].Key)
			} else if len(keys[i].Key) < smallest {
				smallest = len(keys[i].Key)
			}
		}
		return largest
	}

	i := 0
	keys := make(PairList, len(tv))
	for k, v := range tv {
		keys[i] = Pair{k, v}
		i++
	}

	pad := func(s string) string {
		return strings.Repeat(" ", maxLen(keys)-len(s)+1)
	}

	sort.Sort(keys)
	for _, k := range keys {
		s += fmt.Sprintf("\t%q%v%v %d\n", k.Key, pad(k.Key), "=>", k.Value)
	}

	return s
}
