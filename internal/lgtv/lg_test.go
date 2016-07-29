package lgtv

import (
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
			l    int
			name string
			r    RespMap
		}{
			{l: 100, name: "Vanilla", r: Cmd.GetRespMap()},
		}
		for _, tt := range tests {
			tt.r.respMapIDs()
			So(len(tt.r), ShouldEqual, tt.l)
		}
	})
}

func TestTVCmdsGetRespMap(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver.
		tv TVCmds
		// Expected results.
		want RespMap
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.tv.GetRespMap(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. TVCmds.GetRespMap() = %v, want %v", tt.name, got, tt.want)
		}
	}
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
	tests := []struct {
		// Test description.
		name string
		// Receiver.
		r RespMap
		// Expected results.
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.r.String(); got != tt.want {
			t.Errorf("%q. RespMap.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestTVCmdsString(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver.
		tv TVCmds
		// Expected results.
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.tv.String(); got != tt.want {
			t.Errorf("%q. TVCmds.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
