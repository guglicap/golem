package modules

import (
	"fmt"
	"os/user"
	"strings"
)

func powerHub(m Module) {
	result := m.options.PowerHubFormat
	result = strings.Replace(result, "%P", buttonify("poweroff", m.options.PowerOffText), -1)
	result = strings.Replace(result, "%R", buttonify("reboot", m.options.RebootText), -1)
	result = strings.Replace(result, "%S", buttonify("systemctl suspend", m.options.SuspendText), -1)
	fmt.Println(m.options)
	output <- Update{m.position, m.index, result}
}

func whoami(m Module) {
	u, err := user.Current()
	if err != nil {
		errOutput(m, err)
		return
	}
	result := m.options.WhoamiFormat
	result = strings.Replace(result, "%uname", u.Username, -1)
	result = strings.Replace(result, "%name", u.Name, -1)
	result = strings.Replace(result, "%gid", u.Gid, -1)
	result = strings.Replace(result, "%uid", u.Uid, -1)
	output <- Update{m.position, m.index, result}
}

func resources(m Module) {
	
}
