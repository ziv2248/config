package env

import (
	"os"
	"strings"

	"github.com/structproto"
	"github.com/structproto/valuebinder"
)

const (
	TagName = "env"
)

func Process(prefix string, target interface{}) error {
	if len(prefix) > 0 {
		prefix += "_"
	}

	prototype, err := structproto.Prototypify(target, &structproto.StructProtoResolveOption{
		TagName: TagName,
	})
	if err != nil {
		return err
	}

	var table structproto.FieldValueMap = make(structproto.FieldValueMap)
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		name, value := parts[0], parts[1]
		if strings.HasPrefix(name, prefix) {
			table[name[len(prefix):]] = value
		}
	}
	err = prototype.BindValues(table, valuebinder.BuildStringArgsBinder)
	if err != nil {
		return err
	}
	return nil
}
