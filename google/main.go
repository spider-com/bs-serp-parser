package google

import (
	"encoding/json"
	i "github.com/spider-com/bs-serp-parser"
	"io"
)

func ParseJSON(r io.Reader) (res []byte, err error) {
	v, err := parseDesktop(r)
	if err != nil {
		return
	}

	return json.Marshal(v)
}

func ParsePage(r io.Reader, f i.PageFormat) ([]byte, error) {
	switch f {
	default:
		{
			return ParseJSON(r)
		}
	case i.TabletPage:
		v, err := parseTablet(r)
		if err != nil {
			return nil, err
		}
		return json.Marshal(v)
	case i.MobilePage:
		{
			v, err := parseMobile(r)
			if err != nil {
				return nil, err
			}
			return json.Marshal(v)
		}
	}
}

