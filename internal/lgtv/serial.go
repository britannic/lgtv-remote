package lgtv

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/tarm/serial"
)

// MaxTVs sets how many TV sets are in use
var MaxTVs = 5

// Serializer implements Open and Xmit for LGTV serial control
type Serializer interface {
	Open() (*serial.Port, error)
	Xmit(ctx context.Context, id int, cmd string) (bool, error)
}

// Serial implements the Serializer interface
type Serial struct {
	Cmd            TVCmpMap
	Baud           int
	Port           string
	Parity         serial.Parity
	ReadTimeout    time.Duration
	RTSFlowControl bool
	StopBits       serial.StopBits
	XONFlowControl bool
}

// CmdMode sets which API command is used
type CmdMode struct {
	Pair string
	Send string
}

// LGCmd is a struct of serial and WebOS commands
type LGCmd struct {
	Cmd1 string `json:"1st cmd,omitempty"`
	Cmd2 string `json:"2nd cmd,omitempty"`
	Data string `json:"data,omitempty"`
	Max  int    `json:"max,omitempty"`
	Note string `json:"note,omitempty"`
	Web  int    `json:"WebOS,omitempty"`
}

// RespMap is a map of response keys mapped to LG TV functions.
type RespMap map[int]map[string]string

// TVCmpMap is a map of xmit comands and responses
type TVCmpMap map[int]map[string]XmitRes

// XmitRes is a transmit and response command struct
type XmitRes struct {
	Xmit []byte            `json:"Xmit,omitempty"`
	Resp map[string][]byte `json:"Resp,omitempty"`
}

// TVCmds is a map[string]LGCmd of RS-232C serial and WebOS commands.
type TVCmds map[string]LGCmd

func (r RespMap) respMapIDs() {
	for i := 0; i < MaxTVs; i++ {
		r[i] = make(map[string]string)
	}
}

// SetSerialCmds builds a set of serial commands
func (tv TVCmds) SetSerialCmds() TVCmpMap {
	ok := func(l LGCmd) bool {
		if l.Data == "FF" || (l.Cmd1 == "" && l.Cmd2 == "") {
			return false
		}
		return true
	}

	xmitres := func(cmd1, cmd2, id, data string) XmitRes {
		x := XmitRes{
			Resp: make(map[string][]byte),
			Xmit: []byte(fmt.Sprintf("%s %s %v %s\n", cmd1, cmd2, id, data)),
		}
		for _, code := range []string{"NG", "OK"} {
			x.Resp[code] = []byte(fmt.Sprintf("%s %s %s %s %sx", cmd1, cmd2, code, id, data))
		}
		return x
	}

	tvc := make(TVCmpMap)
	for i := 0; i < MaxTVs; i++ {
		tvc[i] = make(map[string]XmitRes)
	}

	for id := range tvc {
		idStr := strconv.Itoa(id)
		if id < 10 {
			idStr = "0" + strconv.Itoa(id)
		}
		for tvKey := range tv {
			v := tv[tvKey]
			if ok(v) {
				switch v.Max {
				case 0:
					tvc[id][tvKey+v.Data] = xmitres(v.Cmd1, v.Cmd2, idStr, v.Data)
				default:
					for i := 0; i <= v.Max; i++ {
						data := strconv.Itoa(i)
						if i < 10 {
							data += "0"
						}
						tvc[id][tvKey+data] = xmitres(v.Cmd1, v.Cmd2, idStr, data)
					}
				}
			}
		}
	}

	return tvc
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

func (r RespMap) String() string {
	b, _ := json.MarshalIndent(r, "", "\t")
	return string(b)
}

func (tv TVCmds) String() string {
	b, _ := json.MarshalIndent(tv, "", "\t")
	return string(b)
}

func (tv TVCmpMap) String() string {
	b, _ := json.MarshalIndent(tv, "", "\t")
	return string(b)
}

// Open opens an asynchronous communications port.
func (s *Serial) Open() (*serial.Port, error) {
	return serial.OpenPort(
		&serial.Config{
			Baud:        s.Baud,
			Name:        s.Port,
			Parity:      s.Parity,
			ReadTimeout: s.ReadTimeout,
		})
}

// Xmit sends a request using the a selected serial driver
func (s Serial) Xmit(ctx context.Context, id int, cmd string) (bool, error) {
	return true, nil
}
