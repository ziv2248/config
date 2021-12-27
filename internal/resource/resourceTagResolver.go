package resource

import (
	"strings"

	"github.com/structproto"
)

var _ structproto.TagResolver = ResourceTagResolver

func ResourceTagResolver(fieldname, token string) (*structproto.Tag, error) {
	if len(token) > 0 {
		parts := strings.SplitN(token, ";", 2)
		var desc string
		if len(parts) == 2 {
			parts, desc = strings.Split(parts[0], ","), parts[1]
		} else {
			parts = strings.Split(token, ",")
		}
		name, flags := parts[0], parts[1:]

		var tag *structproto.Tag
		if len(name) > 0 && name != "-" {
			tag = &structproto.Tag{
				Name:  name,
				Flags: flags,
				Desc:  desc,
			}
		}
		return tag, nil
	}
	return nil, nil
}
