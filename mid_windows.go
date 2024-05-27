//go:build windows

package mid

import (
	"os/exec"
	"regexp"
	"syscall"
)

func UUID() string {
	c := exec.Command("wmic", "CsProduct", "Get", "UUID")
	c.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000,
	}
	output, err := c.Output()
	if err != nil {
		return newUUID()
	}
	pattern := regexp.MustCompile(`([\w\d]+)-([\w\d]+)-([\w\d]+)-([\w\d]+)-([\w\d]+)`)
	s := pattern.FindString(string(output))
	return s
}
