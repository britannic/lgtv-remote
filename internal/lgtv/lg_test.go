package lgtv

import (
	"crypto/md5"
	"io"
	"net"
	"reflect"
	"testing"
	"time"

	logging "github.com/op/go-logging"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
)

func TestRespMapIDs(t *testing.T) {
	Convey("Testing respMapIDs()", t, func() {
		tests := []struct {
			rLen  int
			rkLen int
			name  string
			r     RespMap
			rTV   TVCmds
		}{
			{rLen: 5, rkLen: 2, name: "Single Record", rTV: TVCmds{"Single-Step": {Cmd1: "k", Cmd2: "z"}}},
			{rLen: 5, rkLen: 130, name: "Step Generator", rTV: TVCmds{"Multi-Step": {Cmd1: "k", Cmd2: "q", Max: 64}}},
			{rLen: 5, rkLen: 0, name: "WebOS Only", rTV: TVCmds{"WebOs": {}}},
		}

		for _, tt := range tests {
			tt.r = tt.rTV.GetRespMap()
			Convey("running test: "+tt.name, func() {
				So(len(tt.r), ShouldEqual, tt.rLen)
				for k := range tt.r {
					So(len(tt.r[k]), ShouldEqual, tt.rkLen)
				}
			})
		}
	})
}

func TestAPISend(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rLogger  *logging.Logger
		rAppID   string
		rAppName string
		rctx     context.Context
		rFound   bool
		rID      string
		rIP      net.IP
		rName    string
		rPin     string
		rTimeout time.Duration
		// Parameters.
		cmd string
		msg []byte
		// Expected results.
		want    int
		want1   io.Reader
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &API{
			Logger:  tt.rLogger,
			AppID:   tt.rAppID,
			AppName: tt.rAppName,
			Found:   tt.rFound,
			ID:      tt.rID,
			IP:      tt.rIP,
			Name:    tt.rName,
			Pin:     tt.rPin,
			Timeout: tt.rTimeout,
		}
		got, got1, err := a.Send(tt.cmd, tt.msg)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. API.Send() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. API.Send() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. API.Send() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestAPIShowPIN(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rLogger  *logging.Logger
		rAppID   string
		rAppName string
		rctx     context.Context
		rFound   bool
		rID      string
		rIP      net.IP
		rName    string
		rPin     string
		rTimeout time.Duration
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &API{
			Logger:  tt.rLogger,
			AppID:   tt.rAppID,
			AppName: tt.rAppName,
			Found:   tt.rFound,
			ID:      tt.rID,
			IP:      tt.rIP,
			Name:    tt.rName,
			Pin:     tt.rPin,
			Timeout: tt.rTimeout,
		}
		a.ShowPIN()
	}
}

func TestAPIPair(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rLogger  *logging.Logger
		rAppID   string
		rAppName string
		rctx     context.Context
		rFound   bool
		rID      string
		rIP      net.IP
		rName    string
		rPin     string
		rTimeout time.Duration
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &API{
			Logger:  tt.rLogger,
			AppID:   tt.rAppID,
			AppName: tt.rAppName,
			Found:   tt.rFound,
			ID:      tt.rID,
			IP:      tt.rIP,
			Name:    tt.rName,
			Pin:     tt.rPin,
			Timeout: tt.rTimeout,
		}
		a.Pair()
	}
}

func TestAPIZap(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rLogger  *logging.Logger
		rAppID   string
		rAppName string
		rctx     context.Context
		rFound   bool
		rID      string
		rIP      net.IP
		rName    string
		rPin     string
		rTimeout time.Duration
		// Parameters.
		cmd int
		// Expected results.
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &API{
			Logger:  tt.rLogger,
			AppID:   tt.rAppID,
			AppName: tt.rAppName,
			Found:   tt.rFound,
			ID:      tt.rID,
			IP:      tt.rIP,
			Name:    tt.rName,
			Pin:     tt.rPin,
			Timeout: tt.rTimeout,
		}
		if got := a.Zap(tt.cmd); got != tt.want {
			t.Errorf("%q. API.Zap() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestAPIScan(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rLogger  *logging.Logger
		rAppID   string
		rAppName string
		rctx     context.Context
		rFound   bool
		rID      string
		rIP      net.IP
		rName    string
		rPin     string
		rTimeout time.Duration
		// Parameters.
		portAddr string
		msg      []byte
		// Expected results.
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &API{
			Logger:  tt.rLogger,
			AppID:   tt.rAppID,
			AppName: tt.rAppName,
			Found:   tt.rFound,
			ID:      tt.rID,
			IP:      tt.rIP,
			Name:    tt.rName,
			Pin:     tt.rPin,
			Timeout: tt.rTimeout,
		}
		if err := a.scan(tt.portAddr, tt.msg); (err != nil) != tt.wantErr {
			t.Errorf("%q. API.scan() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestAPIGetLocalIP(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rLogger  *logging.Logger
		rAppID   string
		rAppName string
		rctx     context.Context
		rFound   bool
		rID      string
		rIP      net.IP
		rName    string
		rPin     string
		rTimeout time.Duration
		// Expected results.
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &API{
			Logger:  tt.rLogger,
			AppID:   tt.rAppID,
			AppName: tt.rAppName,
			Found:   tt.rFound,
			ID:      tt.rID,
			IP:      tt.rIP,
			Name:    tt.rName,
			Pin:     tt.rPin,
			Timeout: tt.rTimeout,
		}
		got, err := a.getLocalIP()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. API.getLocalIP() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. API.getLocalIP() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestAPISetUpSox(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rLogger  *logging.Logger
		rAppID   string
		rAppName string
		rctx     context.Context
		rFound   bool
		rID      string
		rIP      net.IP
		rName    string
		rPin     string
		rTimeout time.Duration
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &API{
			Logger:  tt.rLogger,
			AppID:   tt.rAppID,
			AppName: tt.rAppName,
			Found:   tt.rFound,
			ID:      tt.rID,
			IP:      tt.rIP,
			Name:    tt.rName,
			Pin:     tt.rPin,
			Timeout: tt.rTimeout,
		}
		a.setUpSox()
	}
}

func TestAPIChkMsgs(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rLogger  *logging.Logger
		rAppID   string
		rAppName string
		rctx     context.Context
		rFound   bool
		rID      string
		rIP      net.IP
		rName    string
		rPin     string
		rTimeout time.Duration
		// Expected results.
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &API{
			Logger:  tt.rLogger,
			AppID:   tt.rAppID,
			AppName: tt.rAppName,
			Found:   tt.rFound,
			ID:      tt.rID,
			IP:      tt.rIP,
			Name:    tt.rName,
			Pin:     tt.rPin,
			Timeout: tt.rTimeout,
		}
		got, err := a.chkMsgs()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. API.chkMsgs() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. API.chkMsgs() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestAPIParseMsg(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rLogger  *logging.Logger
		rAppID   string
		rAppName string
		rctx     context.Context
		rFound   bool
		rID      string
		rIP      net.IP
		rName    string
		rPin     string
		rTimeout time.Duration
		// Parameters.
		msg  string
		addr *net.UDPAddr
		// Expected results.
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &API{
			Logger:  tt.rLogger,
			AppID:   tt.rAppID,
			AppName: tt.rAppName,
			Found:   tt.rFound,
			ID:      tt.rID,
			IP:      tt.rIP,
			Name:    tt.rName,
			Pin:     tt.rPin,
			Timeout: tt.rTimeout,
		}
		got, err := a.parseMsg(tt.msg, tt.addr)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. API.parseMsg() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. API.parseMsg() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestAPIPairingRequest(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rLogger  *logging.Logger
		rAppID   string
		rAppName string
		rctx     context.Context
		rFound   bool
		rID      string
		rIP      net.IP
		rName    string
		rPin     string
		rTimeout time.Duration
		// Expected results.
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &API{
			Logger:  tt.rLogger,
			AppID:   tt.rAppID,
			AppName: tt.rAppName,
			Found:   tt.rFound,
			ID:      tt.rID,
			IP:      tt.rIP,
			Name:    tt.rName,
			Pin:     tt.rPin,
			Timeout: tt.rTimeout,
		}
		if err := a.pairingRequest(); (err != nil) != tt.wantErr {
			t.Errorf("%q. API.pairingRequest() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestRespMapString(t *testing.T) {
	Convey("Testing RespMapString", t, func() {
		tests := []struct {
			name string
			tv   TVCmds
			want [16]byte
		}{
			{
				name: "First Record",
				want: [16]uint8{215, 244, 60, 243, 171, 65, 7, 193, 207, 206, 234, 187, 146, 48, 125, 209},
				tv: TVCmds{
					"First": {
						Cmd1: "k",
						Cmd2: "z",
					},
				},
			},
			{
				name: "Second Record",
				want: [16]uint8{75, 64, 21, 55, 223, 99, 113, 37, 28, 68, 156, 180, 4, 151, 230, 111},
				tv: TVCmds{
					"Second": {
						Cmd1: "k",
						Cmd2: "q",
					},
				},
			},
			{
				name: "Third Record",
				want: [16]uint8{186, 207, 19, 245, 23, 50, 99, 222, 58, 24, 56, 195, 229, 58, 113, 80},
				tv:   nil,
			},
			{
				name: "Fourth Record",
				want: [16]uint8{83, 0, 124, 146, 32, 252, 218, 80, 35, 6, 225, 28, 219, 198, 101, 210},
				tv: TVCmds{
					"Second": {
						Cmd1: "m",
						Cmd2: "d",
						Max:  10,
					},
				},
			},
		}
		for _, tt := range tests {
			Convey("running test: "+tt.name, func() {
				So(md5.Sum([]byte(tt.tv.GetRespMap().String())), ShouldResemble, tt.want)
			})
		}
	})
}

func TestTVCmdsString(t *testing.T) {
	Convey("Testing RespMapString", t, func() {
		tests := []struct {
			name string
			tv   TVCmds
			want [16]byte
		}{
			{
				name: "First Record",
				want: [16]uint8{12, 26, 190, 252, 199, 125, 235, 5, 14, 113, 205, 173, 70, 96, 215, 196},
				tv: TVCmds{
					"First": {
						Cmd1: "k",
						Cmd2: "z",
					},
				},
			},
			{
				name: "Second Record",
				want: [16]uint8{177, 191, 37, 241, 128, 94, 189, 33, 119, 127, 36, 93, 29, 219, 240, 146},
				tv: TVCmds{
					"Second": {
						Cmd1: "k",
						Cmd2: "q",
					},
				},
			},
			{
				name: "Third Record",
				want: [16]uint8{55, 166, 37, 156, 192, 193, 218, 226, 153, 167, 134, 100, 137, 223, 240, 189},
				tv:   nil,
			},
			{
				name: "Fourth Record",
				want: [16]uint8{31, 168, 171, 57, 176, 252, 60, 162, 176, 56, 122, 33, 245, 228, 202, 2},
				tv: TVCmds{
					"Second": {
						Cmd1: "m",
						Cmd2: "d",
						Max:  10,
					},
				},
			},
		}
		for _, tt := range tests {
			Convey("running test: "+tt.name, func() {
				So(md5.Sum([]byte(tt.tv.String())), ShouldResemble, tt.want)
			})
		}
	})
}
