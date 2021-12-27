package resource

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"

	"github.com/structproto"
	"github.com/structproto/valuebinder"
)

var (
	typeOfByteArray = reflect.TypeOf([]byte{})
)

var _ structproto.StructBinder = new(ResourceBinder)

type ResourceBinder struct {
	BaseDir string
}

func (p *ResourceBinder) Init(context *structproto.StructProtoContext) error {
	return nil
}

func (p *ResourceBinder) Bind(field structproto.FieldInfo, rv reflect.Value) error {
	filename := path.Join(p.BaseDir, field.Name())

	fileinfo, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if fileinfo.Mode().IsRegular() {
		buffer, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		switch rv.Type() {
		case typeOfByteArray:
			rv.Set(reflect.ValueOf(buffer))
			return nil
		}
		return valuebinder.BytesArgsBinder(rv).Bind(buffer)
	}
	return nil
}

func (p *ResourceBinder) Deinit(context *structproto.StructProtoContext) error {
	return nil
}
