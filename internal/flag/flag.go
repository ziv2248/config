package flag

import (
	"flag"

	"github.com/structproto"
)

const (
	TagName = "arg"
)

var (
	help = flag.Bool("help", false, "Show this help")

	binder = &FlagBinder{}
)

func Process(target interface{}) error {
	prototype, err := structproto.Prototypify(target, &structproto.StructProtoResolveOption{
		TagName: TagName,
	})
	if err != nil {
		return err
	}
	return prototype.Bind(binder)
}
