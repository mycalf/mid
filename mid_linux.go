//go:build linux

package mid

import (
	"io/ioutil"
	"strings"
)

func UUID() string {
	data, err := ioutil.ReadFile("/var/lib/dbus/machine-id")
	if err != nil {
		return newUUID()
	}
	id := strings.TrimSpace(string(data))
	if len(id) > 20 {
		return id[0:8] + "-" + id[8:12] + "-" + id[12:16] + "-" + id[16:20] + "-" + id[20:]
	}
	return newUUID()
}
