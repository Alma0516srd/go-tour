package main

import (
	"encoding/json"
	"fmt"
)

func marshalType[T json.Marshaler](val T) (string, error) {
	bytes, err := json.Marshal(val)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type phone struct {
	model     string
	osVersion float64
}

func (p phone) MarshalJSON() ([]byte, error) {
	v := fmt.Sprintf("{\"model\": %q, \"os\": %v}", p.model, p.osVersion)
	return []byte(v), nil
}

func main() {
	iphone := phone{
		osVersion: 23.2,
		model:     "Iphone",
	}
	s, err := marshalType(&iphone)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}
