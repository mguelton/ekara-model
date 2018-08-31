package model

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Parameters map[string]interface{}

func createParameters(src map[string]interface{}) Parameters {
	dst := make(map[string]interface{})
	mergeParams(dst, src)
	return dst
}

func (p Parameters) inherit(parents ...map[string]interface{}) Parameters {
	dst := make(map[string]interface{})
	for i := len(parents) - 1; i >= 0; i-- {
		mergeParams(dst, parents[i])
	}
	mergeParams(dst, p)
	return dst
}

func mergeParams(dst map[string]interface{}, src map[string]interface{}) {
	// TODO take into account map merging
	for k, v := range src {
		dst[k] = v
	}
}

func (t Parameters) MarshalJSON() ([]byte, error) {
	dest := make(map[interface{}]interface{})
	for k, v := range t {
		dest[k] = v
	}
	b, err := MarshalJSONMap(dest)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBufferString("")
	buffer.WriteString(fmt.Sprintf("\"Parameters\":[%s]", string(buffer.Bytes())))
	return b, err
}

func MarshalJSONMap(m map[interface{}]interface{}) ([]byte, error) {
	buffer := bytes.NewBufferString("[")

	length := len(m)
	count := 0
	for key, value := range m {
		mII, ok := value.(map[interface{}]interface{})
		if !ok {
			jsonValue, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			buffer.WriteString(fmt.Sprintf("{\"%s\":%s}", key, string(jsonValue)))
		} else {
			b, err := MarshalJSONMap(mII)
			if err != nil {
				return nil, err
			}
			buffer.WriteString(fmt.Sprintf("{\"%s\":[%s]}", key, string(b)))
		}
		count++
		if count < length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")
	return buffer.Bytes(), nil
}
