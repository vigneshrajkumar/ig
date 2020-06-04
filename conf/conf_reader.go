package conf

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var SysHome string

func init() {
	SysHome = os.Getenv("SYS_HOME")
	if SysHome == "" {
		SysHome = "."
	}
}

// ParseConfFile is used to parse the configuration file in the given path
func ParseConfFile(filePath string) (m map[string]string, err error) {
	data, err := ioutil.ReadFile(SysHome + filePath)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(data, &m); err != nil {
		return
	}

	return
}
