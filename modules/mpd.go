package modules

import (
	"errors"
	"log"

	"strings"

	"encoding/json"

	"io"

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
		":6060",
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
	w, err := mpd.NewWatcher("tcp", ":6600", "", "player")
	if err != nil {
		m.errOutput(err)
		return
	}
	c, err := mpd.Dial("tcp", ":6600")
	if err != nil {
		m.errOutput(err)
		return
	}
	for {
		var result, song, toggle string
		st, err := c.Status()
		if err != nil {
			if err == io.EOF {
				<-w.Event
				continue
			}
			m.errOutput(err)
			return
		}
		switch st["state"] {
		case "play":
			toggle = m.TogglePlaying
		case "pause":
			toggle = m.TogglePaused
		case "stop":
			output <- m.update("Stopped.")
			<-w.Event //Waits for mpd to start playing again.
			continue  //Then start from the top of the loop
		}
		result = strings.Replace(m.Format, "%prev", buttonify("mpc prev", m.PrevButton), -1)
		result = strings.Replace(result, "%next", buttonify("mpc next", m.NextButton), -1)

		result = strings.Replace(result, "%toggle", buttonify("mpc toggle", toggle), -1)
		s, err := c.CurrentSong()
		if err != nil {
			if err == io.EOF {
				<-w.Event
				continue
			}
			m.errOutput(err)
			return
		}
		song = strings.Replace(m.SongFormat, "%title", s["Title"], -1)
		song = strings.Replace(song, "%artist", s["Artist"], -1)
		song = strings.Replace(song, "%album", s["Album"], -1)
		result = strings.Replace(result, "%song", song, -1)
		log.Println(result)
		output <- m.update(result)
		<-w.Event //Blocks until something happens
	}
}
