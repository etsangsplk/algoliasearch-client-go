// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

type RestrictIndicesOption struct {
	value []string
}

func RestrictIndices(v ...string) *RestrictIndicesOption {
	return &RestrictIndicesOption{v}
}

func (o RestrictIndicesOption) Get() []string {
	return o.value
}

func (o RestrictIndicesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *RestrictIndicesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *RestrictIndicesOption) Equal(o2 *RestrictIndicesOption) bool {
	if o2 == nil {
		return reflect.DeepEqual(o.value, []string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

func RestrictIndicesEqual(o1, o2 *RestrictIndicesOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
