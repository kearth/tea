package structs

import "github.com/gogf/gf/v2/encoding/gjson"

type Demo struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (d *Demo) String() string {
	s, _ := gjson.EncodeString(d)
	return s
}
