package modules

type Options struct {
	PingAddress string
	WsFocused   string
	WsUnfocused string
	DateFormat  string
}

var defaultOptions = &Options{
	PingAddress: "google.it",
	WsFocused:   "\uf111",
	WsUnfocused: "\uf10c",
	DateFormat:  "15:04",
}
