//go:build darwin || (openbsd && !mips64)

package mid

import (
	"os/exec"
	"regexp"
)

func UUID() string {
	c := exec.Command("ioreg", "-rd1", "-c", "IOPlatformExpertDevice")
	output, err := c.Output()
	if err != nil {
		return newUUID()
	}
	pattern := regexp.MustCompile(`IOPlatformUUID" = "(.*?)"`)
	ss := pattern.FindStringSubmatch(string(output))
	if len(ss) == 0 {
		return newUUID()
	}
	return ss[1]
}
