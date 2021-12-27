package json

import "github.com/config/internal/json"

func LoadFile(filepath string, target interface{}) error {
	return json.LoadFile(filepath, target)
}

func LoadBytes(buffer []byte, target interface{}) error {
	return json.LoadBytes(buffer, target)
}
