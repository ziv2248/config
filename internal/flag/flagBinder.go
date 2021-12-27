package flag

import (
	"flag"
	"os"
	"reflect"

	"github.com/structproto"
)

var _ structproto.StructBinder = new(FlagBinder)

type FlagBinder struct{}

func (p *FlagBinder) Init(context *structproto.StructProtoContext) error {
	return nil
}

func (p *FlagBinder) Bind(field structproto.FieldInfo, rv reflect.Value) error {
	value := p.makeFlagValue(rv)
	flag.Var(value, field.Name(), field.Desc())
	return nil
}

func (p *FlagBinder) Deinit(context *structproto.StructProtoContext) error {
	// NOTE: ignore validate
	if !flag.Parsed() {
		flag.Parse()
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}
	return nil
}

func (p *FlagBinder) makeFlagValue(rv reflect.Value) flag.Value {
	if rv.CanInterface() {
		if rv.CanAddr() {
			rv = rv.Addr()
		}
		v := rv.Interface()
		switch v.(type) {
		case flag.Value:
			value, ok := v.(flag.Value)
			if ok {
				return value
			}
		}
	}
	return &FlagValue{rv}
}
