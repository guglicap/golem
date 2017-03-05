package modules

import (
	"encoding/json"
	"os/user"
	"strings"
)

type WhoamiModule struct {
	ModuleBase
	Format string
}

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
	output <- Update{m.slot, m.colors, result}
}
