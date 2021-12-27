package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
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
	err := json.Unmarshal(buffer, target)
	if err != nil {
		return err
	}
	return nil
}
