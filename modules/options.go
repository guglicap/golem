package modules

//Options holds all of the possible options for a module.
type Options struct {
	//Ws
	WsFocused   string
	WsUnfocused string
	//Date
	DateFormat string
	//Text
	Text string
	//NetAddress
	NetInterface string
	//PowerTray
	PowerOffText   string
	RebootText     string
	SuspendText    string
	PowerHubFormat string
	//Button
	Command    string
	ButtonText string
	//Whoami
	WhoamiFormat string
	//IconTray
	IconTrayCommands string
	IconTrayText     string
	//DiskInfo
	DiskMount      string
	DiskInfoFormat string
	//MemInfo
	MemInfoFormat string
}

/*func defaultOptions() *Options {
	opts := new(Options)
	opts.WsFocused = "\uf111"
	opts.WsUnfocused = "\uf10c"
	opts.PowerHubFormat = "P"
	opts.PowerOffText = "\uf011"
	opts.RebootText = "\uf021"
	opts.SuspendText = "\uf186"
	opts.DateFormat = "15:04 02/01/2006"
	opts.Text = ""
	opts.Command = ""
	opts.ButtonText = ""
	opts.NetInterface = ""
	return opts
}*/

var defaultOptions = map[string]*Options{
	"workspaces": &Options{
		WsFocused:   "\uf111",
		WsUnfocused: "\uf10c",
	},

	"powertray": &Options{
		PowerHubFormat: "%P",
		PowerOffText:   "\uf011",
		RebootText:     "\uf021",
		SuspendText:    "\uf186",
	},

	"date": &Options{
		DateFormat: "15:04 02/01/2006",
	},
	"text": &Options{
		Text: "",
	},
	"button": &Options{
		Command:    "",
		ButtonText: "",
	},
	"netAddress": &Options{
		NetInterface: "enp3s0",
	},
	"whoami": &Options{
		WhoamiFormat: "\uf2be %uname",
	},
	"icontray": &Options{
		IconTrayCommands: "firefox,xterm",
		IconTrayText:     "\uf269,\uf120",
	},
	"diskinfo": &Options{
		DiskInfoFormat: "%mount: %usePerc",
		DiskMount:      "/",
	},
	"meminfo": &Options{
		MemInfoFormat: "Memory: %usedMB / %totalMB",
	},
}
