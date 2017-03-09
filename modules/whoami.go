package modules

import (
	"encoding/json"
	"os/user"
	"strings"
)

//WhoamiModule displays info about the current user
type WhoamiModule struct {
	ModuleBase
	Format string
}

//BuildWhoami initializes a WhoamiModule
func BuildWhoami(ms *ModuleSpec) Module {
	opts := struct {
		Format string
	}{
		"\uf2be %uname",
	}
	json.Unmarshal([]byte(ms.Options), &opts)
	return &WhoamiModule{
		buildModuleBase(ms),
		opts.Format,
	}
}

//Run starts the module
func (m *WhoamiModule) Run() {
	u, err := user.Current()
	if err != nil {
		m.errOutput(err)
		return
	}
	result := m.Format
	result = strings.Replace(result, "%uname", u.Username, -1)
	result = strings.Replace(result, "%name", u.Name, -1)
	result = strings.Replace(result, "%gid", u.Gid, -1)
	result = strings.Replace(result, "%uid", u.Uid, -1)
	output <- m.update(result)
}
