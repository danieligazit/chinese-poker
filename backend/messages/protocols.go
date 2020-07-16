package communication

import (
	"encoding/json"
)

type protocol string

type marshalHandler struct {
	unmarshal func(data []byte, v interface{}) error
	marshal   func(v interface{}) ([]byte, error)
}

const JsonProtocol protocol = "json"

var protocol2marshalHandler map[string]marshalHandler

func init() {
	protocol2marshalHandler = map[string]marshalHandler{
		JsonProtocol: marshalHandler{
			unmarshal: json.Unmarshal,
			marshal:   json.Marshal,
		},
	}
}
