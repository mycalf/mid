package mid

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func getMachineUidPath() string {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		homeDir = "."
	}
	return filepath.Join(homeDir, ".muid")
}

func newUUID() string {
	filePath := getMachineUidPath()
	id := ""
	data, err := ioutil.ReadFile(filePath)
	if err != nil || strings.TrimSpace(string(data)) == "" {
		id = fmt.Sprintf("%s", uuid.NewV4())
		ioutil.WriteFile(filePath, []byte(id), 0644)
	} else {
		return strings.TrimSpace(string(data))
	}
	return id
}

func RemoveTempUidFile() error {
	return os.Remove(getMachineUidPath())
}
