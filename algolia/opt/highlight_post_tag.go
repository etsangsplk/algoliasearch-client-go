// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type HighlightPostTagOption struct {
	value string
}

func HighlightPostTag(v string) *HighlightPostTagOption {
	return &HighlightPostTagOption{v}
}

func (o HighlightPostTagOption) Get() string {
	return o.value
}

func (o HighlightPostTagOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *HighlightPostTagOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = "</em>"
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *HighlightPostTagOption) Equal(o2 *HighlightPostTagOption) bool {
	if o2 == nil {
		return o.value == "</em>"
	}
	return o.value == o2.value
}

func HighlightPostTagEqual(o1, o2 *HighlightPostTagOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
