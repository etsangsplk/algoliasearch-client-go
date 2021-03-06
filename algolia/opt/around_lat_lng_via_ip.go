// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type AroundLatLngViaIPOption struct {
	value bool
}

func AroundLatLngViaIP(v bool) *AroundLatLngViaIPOption {
	return &AroundLatLngViaIPOption{v}
}

func (o AroundLatLngViaIPOption) Get() bool {
	return o.value
}

func (o AroundLatLngViaIPOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AroundLatLngViaIPOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *AroundLatLngViaIPOption) Equal(o2 *AroundLatLngViaIPOption) bool {
	if o2 == nil {
		return o.value == false
	}
	return o.value == o2.value
}

func AroundLatLngViaIPEqual(o1, o2 *AroundLatLngViaIPOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
