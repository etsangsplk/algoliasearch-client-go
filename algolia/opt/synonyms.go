// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type SynonymsOption struct {
	value bool
}

func Synonyms(v bool) *SynonymsOption {
	return &SynonymsOption{v}
}

func (o SynonymsOption) Get() bool {
	return o.value
}

func (o SynonymsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *SynonymsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = true
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *SynonymsOption) Equal(o2 *SynonymsOption) bool {
	if o2 == nil {
		return o.value == true
	}
	return o.value == o2.value
}

func SynonymsEqual(o1, o2 *SynonymsOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
