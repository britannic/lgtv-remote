
# lgtv-remote

# lgtv
    import "github.com/britannic/lgtv-remote/internal/lgtv"





## Variables
``` go
var (
    // Cmd maps LG TV int commands to meaningful names.
    Cmd = TVCmds{
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
)
```


## type API
``` go
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
```
API struct for the LG TV API interface











### func (\*API) Pair
``` go
func (a *API) Pair()
```
Pair using the LG TV's PIN



### func (\*API) Send
``` go
func (a *API) Send(cmd string, msg []byte) (int, io.Reader, error)
```
Send xmits a WebOS request to the LG TV.



### func (\*API) ShowPIN
``` go
func (a *API) ShowPIN()
```
ShowPIN displays the LG TV's PIN (Pairing ID Number) on its screen.



### func (\*API) Zap
``` go
func (a *API) Zap(cmd int) bool
```
Zap xmits a WebOS command.



## type CmdMode
``` go
type CmdMode struct {
    Pair string
    Send string
}
```
CmdMode sets which API command is used











## type IDCmdMap
``` go
type IDCmdMap map[int]TVCmds
```
IDCmdMap is a map TVCmds keyed to IDs











## type LGCmd
``` go
type LGCmd struct {
    Cmd1 string `json:"1st cmd,omitempty"`
    Cmd2 string `json:"2nd cmd,omitempty"`
    Data string `json:"data,omitempty"`
    Max  int    `json:"max,omitempty"`
    Note string `json:"note,omitempty"`
    Web  int    `json:"WebOS,omitempty"`
}
```
LGCmd is a struct of serial and WebOS commands











## type RespMap
``` go
type RespMap map[int]map[string]string
```
RespMap is a map of response keys mapped to LG TV functions.











### func (RespMap) String
``` go
func (r RespMap) String() string
```


## type TVCmds
``` go
type TVCmds map[string]LGCmd
```
TVCmds is a map[string]LGCmd map of RS-232C serial and WebOS commands.











### func (TVCmds) GetRespMap
``` go
func (tv TVCmds) GetRespMap() RespMap
```
GetRespMap creates a map of response keys mapped to LG TV functions.



### func (TVCmds) String
``` go
func (tv TVCmds) String() string
```






- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)