package lgtv

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/pkg/term"
)

type serializer interface {
	Open(name string, options ...func(*term.Term) error) (*term.Term, error)
	Xmit(ctx context.Context, s serializer, tv TVCmds) (bool, error)
}

// CmdMode sets which API command is used
type CmdMode struct {
	Pair string
	Send string
}

const maxTVs = 5

// LGCmd is a struct of serial and WebOS commands
type LGCmd struct {
	Cmd1 string `json:"1st cmd,omitempty"`
	Cmd2 string `json:"2nd cmd,omitempty"`
	Data string `json:"data,omitempty"`
	Max  int    `json:"max,omitempty"`
	Note string `json:"note,omitempty"`
	Web  int    `json:"WebOS,omitempty"`
}

// IDCmdMap is a map of TVCmds keyed to IDs
type IDCmdMap map[int]TVCmds

// RespMap is a map of response keys mapped to LG TV functions.
type RespMap map[int]map[string]string

// TVCmds is a map[string]LGCmd map of RS-232C serial and WebOS commands.
type TVCmds map[string]LGCmd

var (
	conn     *net.UDPConn
	maxTries = 10
	mode     = CmdMode{Pair: "/udap/api/pairing", Send: "/udap/api/command"}
	sock     = false
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

func (r RespMap) String() string {
	b, _ := json.MarshalIndent(r, "", "\t")
	return string(b)
}

func (tv TVCmds) String() string {
	b, _ := json.MarshalIndent(tv, "", "\t")
	return string(b)
}

// Open opens an asynchronous communications port.
func (tv TVCmds) Open(name string, options ...func(*term.Term) error) (*term.Term, error) {
	return term.Open(name, options...)
}

// Xmit sends a request using the a selected serial driver
func (tv TVCmds) Xmit(ctx context.Context, r RespMap, s serializer) (bool, error) {

	return true, nil
}
