package env

import "github.com/config/internal/env"

func Process(prefix string, target interface{}) error {
	return env.Process(prefix, target)
}
