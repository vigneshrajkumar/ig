package conf

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// ParseConfFile is used to parse the configuration file in the given path
func ParseConfFile(filePath string) (m map[string]string, err error) {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(data, &m); err != nil {
		return
	}

	return
}
