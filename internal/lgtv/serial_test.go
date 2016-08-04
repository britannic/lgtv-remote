package lgtv

import (
	"crypto/md5"
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// func TestSetSerialCmds(t *testing.T) {
// 	Convey("Testing SetSerialCmds()", t, func() {
// 		tests := []struct {
// 			rLen  int
// 			rkLen int
// 			name  string
// 			r     TVCmpMap
// 			rTV   TVCmds
// 		}{
// 			{rLen: 5, rkLen: 1, name: "Single Record", rTV: TVCmds{"Single-Step": {Cmd1: "k", Cmd2: "z"}}},
// 			{rLen: 5, rkLen: 65, name: "Step Generator", rTV: TVCmds{"Multi-Step": {Cmd1: "k", Cmd2: "q", Max: 64}}},
// 			{rLen: 5, rkLen: 0, name: "WebOS Only", rTV: TVCmds{"WebOs": {}}},
// 		}
//
// 		for _, tt := range tests {
// 			tt.r = tt.rTV.SetSerialCmds()
// 			Convey("running test: "+tt.name, func() {
// 				So(len(tt.r), ShouldEqual, tt.rLen)
// 				for k := range tt.r {
// 					So(len(tt.r[k]), ShouldEqual, tt.rkLen)
// 				}
// 			})
// 		}
// 	})
// }

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

func TestRespMapString(t *testing.T) {
	Convey("Testing RespMapString()", t, func() {
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
	Convey("Testing TVCmdsString()", t, func() {
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

func TestTVCmpMapString(t *testing.T) {
	Convey("Testing TVCmpMapString()", t, func() {
		tests := []struct {
			name string
			r    TVCmpMap
			tv   TVCmds
			want [16]byte
		}{
			{
				name: "First Record",
				want: [16]uint8{59, 194, 93, 163, 170, 13, 235, 154, 152, 174, 32, 35, 72, 133, 115, 213},
				tv: TVCmds{
					"First": {
						Cmd1: "k",
						Cmd2: "z",
					},
				},
			},
			{
				name: "Second Record",
				want: [16]uint8{159, 138, 231, 87, 31, 145, 29, 225, 69, 148, 218, 27, 119, 93, 241, 89},
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
				want: [16]uint8{28, 129, 53, 217, 86, 33, 246, 108, 221, 252, 213, 23, 212, 204, 230, 240},
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
				tt.r = tt.tv.SetSerialCmds()
				So(md5.Sum([]byte(tt.r.String())), ShouldResemble, tt.want)
			})
		}
	})
}

func TestSetSerialCmds(t *testing.T) {
	Convey("Testing SetSerialCmds()", t, func() {
		tests := []struct {
			name string
			r    TVCmpMap
			tv   TVCmds
			want []string
		}{
			{
				name: "First Record",
				want: []string{"k z 00 01\n", "k z 01 01\n", "k z 02 01\n", "k z 03 01\n", "k z 04 01\n", "k z NG 00 01x", "k z NG 01 01x", "k z NG 02 01x", "k z NG 03 01x", "k z NG 04 01x", "k z OK 00 01x", "k z OK 01 01x", "k z OK 02 01x", "k z OK 03 01x", "k z OK 04 01x"},
				tv: TVCmds{
					"First": {
						Cmd1: "k",
						Cmd2: "z",
						Data: "01",
					},
				},
			},
			{
				name: "Second Record",
				want: []string{"k q 00 03\n", "k q 01 03\n", "k q 02 03\n", "k q 03 03\n", "k q 04 03\n", "k q NG 00 03x", "k q NG 01 03x", "k q NG 02 03x", "k q NG 03 03x", "k q NG 04 03x", "k q OK 00 03x", "k q OK 01 03x", "k q OK 02 03x", "k q OK 03 03x", "k q OK 04 03x"},
				tv: TVCmds{
					"Second": {
						Cmd1: "k",
						Cmd2: "q",
						Data: "03",
					},
				},
			},
			{
				name: "Third Record",
				want: nil,
				tv:   nil,
			},
			{
				name: "Fourth Record",
				want: []string{"m d 00 00\n", "m d 00 10\n", "m d 00 20\n", "m d 00 30\n", "m d 00 40\n", "m d 00 50\n", "m d 00 60\n", "m d 00 70\n", "m d 00 80\n", "m d 00 90\n", "m d 01 00\n", "m d 01 10\n", "m d 01 20\n", "m d 01 30\n", "m d 01 40\n", "m d 01 50\n", "m d 01 60\n", "m d 01 70\n", "m d 01 80\n", "m d 01 90\n", "m d 02 00\n", "m d 02 10\n", "m d 02 20\n", "m d 02 30\n", "m d 02 40\n", "m d 02 50\n", "m d 02 60\n", "m d 02 70\n", "m d 02 80\n", "m d 02 90\n", "m d 03 00\n", "m d 03 10\n", "m d 03 20\n", "m d 03 30\n", "m d 03 40\n", "m d 03 50\n", "m d 03 60\n", "m d 03 70\n", "m d 03 80\n", "m d 03 90\n", "m d 04 00\n", "m d 04 10\n", "m d 04 20\n", "m d 04 30\n", "m d 04 40\n", "m d 04 50\n", "m d 04 60\n", "m d 04 70\n", "m d 04 80\n", "m d 04 90\n", "m d NG 00 00x", "m d NG 00 10x", "m d NG 00 20x", "m d NG 00 30x", "m d NG 00 40x", "m d NG 00 50x", "m d NG 00 60x", "m d NG 00 70x", "m d NG 00 80x", "m d NG 00 90x", "m d NG 01 00x", "m d NG 01 10x", "m d NG 01 20x", "m d NG 01 30x", "m d NG 01 40x", "m d NG 01 50x", "m d NG 01 60x", "m d NG 01 70x", "m d NG 01 80x", "m d NG 01 90x", "m d NG 02 00x", "m d NG 02 10x", "m d NG 02 20x", "m d NG 02 30x", "m d NG 02 40x", "m d NG 02 50x", "m d NG 02 60x", "m d NG 02 70x", "m d NG 02 80x", "m d NG 02 90x", "m d NG 03 00x", "m d NG 03 10x", "m d NG 03 20x", "m d NG 03 30x", "m d NG 03 40x", "m d NG 03 50x", "m d NG 03 60x", "m d NG 03 70x", "m d NG 03 80x", "m d NG 03 90x", "m d NG 04 00x", "m d NG 04 10x", "m d NG 04 20x", "m d NG 04 30x", "m d NG 04 40x", "m d NG 04 50x", "m d NG 04 60x", "m d NG 04 70x", "m d NG 04 80x", "m d NG 04 90x", "m d OK 00 00x", "m d OK 00 10x", "m d OK 00 20x", "m d OK 00 30x", "m d OK 00 40x", "m d OK 00 50x", "m d OK 00 60x", "m d OK 00 70x", "m d OK 00 80x", "m d OK 00 90x", "m d OK 01 00x", "m d OK 01 10x", "m d OK 01 20x", "m d OK 01 30x", "m d OK 01 40x", "m d OK 01 50x", "m d OK 01 60x", "m d OK 01 70x", "m d OK 01 80x", "m d OK 01 90x", "m d OK 02 00x", "m d OK 02 10x", "m d OK 02 20x", "m d OK 02 30x", "m d OK 02 40x", "m d OK 02 50x", "m d OK 02 60x", "m d OK 02 70x", "m d OK 02 80x", "m d OK 02 90x", "m d OK 03 00x", "m d OK 03 10x", "m d OK 03 20x", "m d OK 03 30x", "m d OK 03 40x", "m d OK 03 50x", "m d OK 03 60x", "m d OK 03 70x", "m d OK 03 80x", "m d OK 03 90x", "m d OK 04 00x", "m d OK 04 10x", "m d OK 04 20x", "m d OK 04 30x", "m d OK 04 40x", "m d OK 04 50x", "m d OK 04 60x", "m d OK 04 70x", "m d OK 04 80x", "m d OK 04 90x"},
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
				tt.r = tt.tv.SetSerialCmds()
				var x []string
				for id := range tt.r {
					for k := range tt.r[id] {
						x = append(x, string(tt.r[id][k].Resp["OK"]))
						x = append(x, string(tt.r[id][k].Resp["NG"]))
						x = append(x, string(tt.r[id][k].Xmit))
					}
				}
				sort.Strings(x)
				So(x, ShouldResemble, tt.want)
			})
		}
	})
}
