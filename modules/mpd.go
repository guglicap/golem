package modules

import (
	"errors"

	"strings"

	"encoding/json"

	"time"

	"github.com/fhs/gompd/mpd"
)

//MpdModule displays what song is currently playing and allows control such as pause, next, prev
type MpdModule struct {
	ModuleBase
	Format, SongFormat                                  string
	PrevButton, NextButton, TogglePlaying, TogglePaused string
	Address, Password                                   string
}

//BuildMpd initializes a MpdModule
func BuildMpd(ms *ModuleSpec) Module {
	opts := struct {
		Format, SongFormat                                  string
		PrevButton, NextButton, TogglePlaying, TogglePaused string
		Address, Password                                   string
	}{
		"%song  %toggle",
		"%artist - %title",
		"\uf04a",
		"\uf04e",
		"\uf04c",
		"\uf04b",
		":6600",
		"",
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &MpdModule{
		buildModuleBase(ms),
		opts.Format,
		opts.SongFormat,
		opts.PrevButton,
		opts.NextButton,
		opts.TogglePlaying,
		opts.TogglePaused,
		opts.Address,
		opts.Password,
	}
}

//Run starts the module
func (m *MpdModule) Run() {
	if ok := inPATH("mpd"); !ok {
		m.errOutput(errors.New("Can't find mpd in $PATH"))
		return
	}
	if ok := inPATH("mpc"); !ok {
		m.errOutput(errors.New("Can't find mpc in $PATH"))
		return
	}
	w, err := mpd.NewWatcher("tcp", m.Address, m.Password, "player")
	if err != nil {
		m.errOutput(err)
		return
	}
	defer w.Close()
	c, err := mpd.Dial("tcp", m.Address)
	if err != nil {
		m.errOutput(err)
		return
	}
	q := make(chan bool)
	e := keepAlive(q, c) //Keep the connection to mpd alive
	defer func() {
		q <- true
	}() //Ensure the keepAlive goroutine stops when this one stops
	var result, song, toggle string
	for {
		st, _ := c.Status() //Ignore errors, keepAlive will tell us in case something's wrong
		switch st["state"] {
		case "play":
			toggle = m.TogglePlaying
		case "pause":
			toggle = m.TogglePaused
		}
		//Place our buttons
		result = strings.Replace(m.Format, "%prev", buttonify("mpc prev", m.PrevButton), -1)
		result = strings.Replace(result, "%next", buttonify("mpc next", m.NextButton), -1)
		result = strings.Replace(result, "%toggle", buttonify("mpc toggle", toggle), -1)
		//Get current song info
		s, _ := c.CurrentSong()
		song = strings.Replace(m.SongFormat, "%title", s["Title"], -1)
		song = strings.Replace(song, "%artist", s["Artist"], -1)
		song = strings.Replace(song, "%album", s["Album"], -1)
		//Put it all together
		result = strings.Replace(result, "%song", song, -1)
		if st["state"] == "stop" {
			result = "mpd stopped."
		}
		output <- m.update(result)
		select {
		case <-w.Event: //wait for something to happen
			break
		case err := <-e: //we couldn't connect to mpd, stop.
			m.errOutput(err)
			return
		}
	}
}

//Apparently, this is needed because mpd will close our connection after some time of inactivity
func keepAlive(quit chan bool, c *mpd.Client) chan error {
	e := make(chan error)
	go func() {
		t := time.NewTicker(45 * time.Second) //mpd default timeout is one minute
		for {
			select {
			case <-t.C:
				err := c.Ping()
				if err != nil {
					e <- err
				}
			case <-quit:
				return
			}
		}
	}()
	return e
}
