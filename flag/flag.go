package flag

import "github.com/config/internal/flag"

func Process(target interface{}) error {
	return flag.Process(target)
}
