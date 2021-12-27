package flag

import (
	"flag"
	"reflect"

	"github.com/structproto/valuebinder"
)

var _ flag.Value = new(FlagValue)

type FlagValue struct {
	value reflect.Value
}

func (fv *FlagValue) String() string {
	return fv.value.String()
}

func (fv *FlagValue) Set(v string) error {
	return valuebinder.StringArgsBinder(fv.value).Bind(v)
}
