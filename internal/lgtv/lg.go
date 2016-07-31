package lgtv

import (
	"bytes"
	"encoding/json"
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

// API struct for the LG TV API interface
type API struct {
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

// CmdMode sets which API command is used
type CmdMode struct {
	Pair string
	Send string
}

const (
	agent   = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/601.7.7 (KHTML, like Gecko) Version/9.1.2 Safari/601.7.7`
	cr      = "\r\n"
	httpStr = `http://`
	maxTVs  = 5
	udp4    = "udp4"
)

// LGCmd is a struct of serial and WebOS commands
type LGCmd struct {
	Cmd1 string `json:"1st cmd,omitempty"`
	Cmd2 string `json:"2nd cmd,omitempty"`
	Data string `json:"data,omitempty"`
	Max  int    `json:"max,omitempty"`
	Note string `json:"note,omitempty"`
	Web  int    `json:"WebOS,omitempty"`
}

// IDCmdMap is a map TVCmds keyed to IDs
type IDCmdMap map[int]TVCmds

// RespMap is a a map of response keys mapped to LG TV functions.
type RespMap map[int]map[string]string

// TVCmds is a map[string]struct{} map of commands
type TVCmds map[string]LGCmd

var (
	// Cmd maps LG TV int commands to meaningful names
	Cmd = TVCmds{
		// %v is a placeholder for ID (if multi 2nd %v setting), 00 = units
		"3D_LR":           {Web: 401},
		"3D":              {Web: 400},
		"AbnormalRead":    {Cmd1: "k", Cmd2: "z", Data: "FF"},
		"Abnormal0":       {Cmd2: "z", Data: "00", Note: "Normal (Power on and signal exist)"},
		"Abnormal1":       {Cmd2: "z", Data: "01", Note: "No signal (Power on)"},
		"Abnormal2":       {Cmd2: "z", Data: "02", Note: "Turn the monitor off by remote control"},
		"Abnormal3":       {Cmd2: "z", Data: "03", Note: "Turn the monitor off by sleep time function"},
		"Abnormal4":       {Cmd2: "z", Data: "04", Note: "Turn the monitor off by RS-232C function"},
		"Abnormal6":       {Cmd2: "z", Data: "06", Note: "AC down"},
		"Abnormal8":       {Cmd2: "z", Data: "08", Note: "Turn the monitor off by off time function"},
		"Abnormal9":       {Cmd2: "z", Data: "09", Note: "Turn the monitor off by auto off function"},
		"AfterImgInv":     {Cmd1: "j", Cmd2: "p", Data: "01"},
		"AfterImgNorm":    {Cmd1: "j", Cmd2: "p", Data: "08"},
		"AfterImgOrbit":   {Cmd1: "j", Cmd2: "p", Data: "02"},
		"AfterImgWtWash":  {Cmd1: "j", Cmd2: "p", Data: "04"},
		"Apps":            {Web: 417},
		"Aspect1:1(PC)":   {Cmd1: "k", Cmd2: "c", Data: "09", Web: 46},
		"Aspect14:9":      {Cmd1: "k", Cmd2: "c", Data: "07", Web: 46},
		"Aspect16:9":      {Cmd1: "k", Cmd2: "c", Data: "02", Web: 46},
		"Aspect4:3":       {Cmd1: "k", Cmd2: "c", Data: "01", Web: 46},
		"AspectFull":      {Cmd1: "k", Cmd2: "c", Data: "08", Web: 46},
		"AspectHoriz":     {Cmd1: "k", Cmd2: "c", Data: "03", Web: 46},
		"AspectStatus":    {Cmd1: "k", Cmd2: "c", Data: "FF", Web: 46},
		"AspectZoom1":     {Cmd1: "k", Cmd2: "c", Data: "04", Web: 46},
		"AspectZoom2":     {Cmd1: "k", Cmd2: "c", Data: "05", Web: 46},
		"AudioDesc":       {Web: 407},
		"AutoConf(RGB)PC": {Cmd1: "j", Cmd2: "u", Data: "01"},
		"AV":              {Web: 410},
		"Back":            {Web: 23},
		"BalanceLevel":    {Cmd1: "k", Cmd2: "t", Data: "FF"},
		"BalanceSet":      {Cmd1: "k", Cmd2: "t", Max: 64},
		"Blue":            {Web: 29},
		"BrightLevel":     {Cmd1: "k", Cmd2: "h", Data: "FF"},
		"BrightSet":       {Cmd1: "k", Cmd2: "h", Max: 64},
		"Ch_Dn":           {Web: 28},
		"Ch_Up":           {Web: 27},
		"ColorCool":       {Cmd1: "k", Cmd2: "u", Data: "01"},
		"ColorLevel":      {Cmd1: "k", Cmd2: "i", Data: "FF"},
		"ColorNormal":     {Cmd1: "k", Cmd2: "u", Data: "00"},
		"ColorSet":        {Cmd1: "k", Cmd2: "i", Max: 64},
		"ColorTempLvl":    {Cmd1: "k", Cmd2: "u", Data: "FF"},
		"ColorUser":       {Cmd1: "k", Cmd2: "u", Data: "03"},
		"ColorWarm":       {Cmd1: "k", Cmd2: "u", Data: "02"},
		"ContrastLvl":     {Cmd1: "k", Cmd2: "g", Data: "FF"},
		"ContrastSet":     {Cmd1: "k", Cmd2: "g", Max: 64},
		"Dash":            {Web: 402},
		"Down":            {Web: 13},
		"EnergySave":      {Web: 409},
		"EPG":             {Web: 44},
		"Exit":            {Web: 412},
		"Ext":             {Web: 47},
		"Fave":            {Web: 404},
		"FF":              {Web: 36},
		"Green":           {Web: 30},
		"Home":            {Web: 21},
		"Info":            {Web: 45},
		"InputAV":         {Cmd1: "k", Cmd2: "b", Data: "02"},
		"InputCmpnt1":     {Cmd1: "k", Cmd2: "b", Data: "04"},
		"InputCmpnt2":     {Cmd1: "k", Cmd2: "b", Data: "05"},
		"InputHDMI(DTV)":  {Cmd1: "k", Cmd2: "b", Data: "08"},
		"InputHDMI(PC)":   {Cmd1: "k", Cmd2: "b", Data: "09"},
		"InputRGB(DTV)":   {Cmd1: "k", Cmd2: "b", Data: "06"},
		"InputRGB(PC)":    {Cmd1: "k", Cmd2: "b", Data: "07"},
		"InternalTemp":    {Cmd1: "d", Cmd2: "n", Data: "FF", Note: "The data is 1 byte long in Hexadecimal."},
		"LampCheck":       {Cmd1: "d", Cmd2: "p", Data: "FF"},
		"LampFault":       {Cmd1: "d", Cmd2: "p", Data: "00"},
		"LampOk":          {Cmd1: "d", Cmd2: "p", Data: "01"},
		"Left":            {Web: 14},
		"Live":            {Web: 43},
		"Mark":            {Web: 52},
		"Menu":            {Web: 22},
		"MuteStatus":      {Cmd1: "k", Cmd2: "e", Data: "FF"},
		"MuteOff":         {Cmd1: "k", Cmd2: "e", Data: "01", Web: 26},
		"MuteOn":          {Cmd1: "k", Cmd2: "e", Data: "00", Web: 26},
		"Netcast":         {Web: 408},
		"Num0":            {Cmd1: "m", Cmd2: "c", Data: "02", Web: 2},
		"Num1":            {Cmd1: "m", Cmd2: "c", Data: "03", Web: 3},
		"Num2":            {Cmd1: "m", Cmd2: "c", Data: "04", Web: 4},
		"Num3":            {Cmd1: "m", Cmd2: "c", Data: "05", Web: 5},
		"Num4":            {Cmd1: "m", Cmd2: "c", Data: "06", Web: 6},
		"Num5":            {Cmd1: "m", Cmd2: "c", Data: "07", Web: 7},
		"Num6":            {Cmd1: "m", Cmd2: "c", Data: "08", Web: 7},
		"Num7":            {Cmd1: "m", Cmd2: "c", Data: "09", Web: 9},
		"Num8":            {Cmd1: "m", Cmd2: "c", Data: "10", Web: 10},
		"Num9":            {Cmd1: "m", Cmd2: "c", Data: "11", Web: 11},
		"OK":              {Web: 20},
		"OSDOff":          {Cmd1: "k", Cmd2: "l", Data: "00"},
		"OSDOn":           {Cmd1: "k", Cmd2: "l", Data: "01"},
		"Pause":           {Web: 34},
		"PIP_CH_Down":     {Web: 415},
		"PIP_CH_Up":       {Web: 414},
		"PIP_Switch":      {Web: 416},
		"PIP":             {Web: 48},
		"Play":            {Web: 33},
		"PowerOff":        {Cmd1: "k", Cmd2: "a", Data: "00", Web: 0},
		"PowerOn":         {Cmd1: "k", Cmd2: "a", Data: "01", Web: 1},
		"PowerStatus":     {Cmd1: "k", Cmd2: "a", Data: "FF"},
		"PrevCh":          {Web: 403},
		"ProgList":        {Web: 50},
		"QuickMenu":       {Web: 405},
		"REC_List":        {Web: 41},
		"REC":             {Web: 40},
		"Red":             {Web: 31},
		"RemoteDisable":   {Cmd1: "k", Cmd2: "m", Data: "00"},
		"RemoteEnable":    {Cmd1: "k", Cmd2: "m", Data: "01"},
		"Repeat":          {Web: 42},
		"Reserve":         {Web: 413},
		"REW":             {Web: 37},
		"Right":           {Web: 15},
		"ScreenOff":       {Cmd1: "k", Cmd2: "d", Data: "00"},
		"ScreenOn":        {Cmd1: "k", Cmd2: "d", Data: "01"},
		"SharpLevel":      {Cmd1: "k", Cmd2: "k", Data: "FF"},
		"SharpSet":        {Cmd1: "k", Cmd2: "k", Max: 64},
		"SimpLink":        {Web: 411},
		"SkipFF":          {Web: 38},
		"SkipREW":         {Web: 39},
		"Stop":            {Web: 35},
		"Subtitle":        {Web: 49},
		"Text_Opt":        {Web: 406},
		"Text":            {Web: 51},
		"Tile1x2":         {Cmd1: "d", Cmd2: "d", Data: "12", Note: "(column x row)"},
		"Tile1x3":         {Cmd1: "d", Cmd2: "d", Data: "13", Note: "(column x row)"},
		"Tile1x4":         {Cmd1: "d", Cmd2: "d", Data: "14", Note: "(column x row)"},
		"Tile2x2":         {Cmd1: "d", Cmd2: "d", Data: "22", Note: "(column x row)"},
		"Tile2x3":         {Cmd1: "d", Cmd2: "d", Data: "23", Note: "(column x row)"},
		"Tile2x4":         {Cmd1: "d", Cmd2: "d", Data: "24", Note: "(column x row)"},
		"Tile3x2":         {Cmd1: "d", Cmd2: "d", Data: "32", Note: "(column x row)"},
		"Tile3x3":         {Cmd1: "d", Cmd2: "d", Data: "33", Note: "(column x row)"},
		"Tile3x4":         {Cmd1: "d", Cmd2: "d", Data: "34", Note: "(column x row)"},
		"Tile4x2":         {Cmd1: "d", Cmd2: "d", Data: "42", Note: "(column x row)"},
		"Tile4x3":         {Cmd1: "d", Cmd2: "d", Data: "43", Note: "(column x row)"},
		"Tile4x4":         {Cmd1: "d", Cmd2: "d", Data: "44", Note: "(column x row)"},
		"TileID":          {Cmd1: "d", Cmd2: "i", Max: 10},
		"TileOff":         {Cmd1: "d", Cmd2: "d", Data: "00"},
		"TileSizeH":       {Cmd1: "d", Cmd2: "g", Max: 64},
		"TileSizeV":       {Cmd1: "d", Cmd2: "h", Max: 64},
		"TimeElapsed":     {Cmd1: "d", Cmd2: "l", Data: "FF", Note: "The data means used hours. (Hexadecimal code)"},
		"TintLevel":       {Cmd1: "k", Cmd2: "j", Data: "FF"},
		"TintSet":         {Cmd1: "k", Cmd2: "j", Max: 64},
		"Up":              {Web: 12},
		"VolDn":           {Web: 25},
		"VolLvl":          {Cmd1: "k", Cmd2: "f", Data: "FF"},
		"VolSet":          {Cmd1: "k", Cmd2: "f", Max: 64},
		"VolUp":           {Web: 24},
		"Yellow":          {Web: 32},
	}

	conn     *net.UDPConn // UDP Connection
	maxTries = 10

	mode = CmdMode{
		Pair: "/udap/api/pairing",
		Send: "/udap/api/command",
	}

	sock = false
)

func (r RespMap) respMapIDs() {
	for i := 0; i < 5; i++ {
		r[i] = make(map[string]string)
	}
}

// GetRespMap creates a map of response keys mapped to LG TV functions.
func (tv TVCmds) GetRespMap() RespMap {
	ok := func(v LGCmd) bool {
		if v.Data == "FF" || (v.Cmd1 == "" && v.Cmd2 == "") {
			return false
		}
		return true
	}

	r := make(RespMap)
	r.respMapIDs()
	for id := range r {
		idStr := strconv.Itoa(id)
		if id < 10 {
			idStr = "0" + strconv.Itoa(id)
		}
		for tvKey := range tv {
			v := tv[tvKey]
			if ok(v) {
				for _, code := range []string{"NG", "OK"} {
					switch v.Max {
					case 0:
						r[id][fmt.Sprintf("%s %s %s %s %s x", v.Cmd1, v.Cmd2, code, idStr, v.Data)] = tvKey
					default:
						for i := 0; i <= v.Max; i++ {
							data := strconv.Itoa(i)
							if i < 10 {
								data = "0" + data
							}
							r[id][fmt.Sprintf("%s %s %s %s %s x", v.Cmd1, v.Cmd2, code, idStr, data)] = tvKey
						}
					}
				}
			}
		}
	}
	return r
}

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

func (r RespMap) String() string {
	b, _ := json.MarshalIndent(r, "", "\t")
	return string(b)
}

func (tv TVCmds) String() string {
	b, _ := json.MarshalIndent(tv, "", "\t")
	return string(b)
}
