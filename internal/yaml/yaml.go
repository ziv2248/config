package yaml

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadFile(filepath string, target interface{}) error {
	path := os.ExpandEnv(filepath)
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return LoadBytes(buffer, target)
}

func LoadBytes(buffer []byte, target interface{}) error {
	err := yaml.Unmarshal(buffer, target)
	if err != nil {
		return err
	}
	return nil
}
