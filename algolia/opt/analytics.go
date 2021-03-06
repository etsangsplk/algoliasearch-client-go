// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type AnalyticsOption struct {
	value bool
}

func Analytics(v bool) *AnalyticsOption {
	return &AnalyticsOption{v}
}

func (o AnalyticsOption) Get() bool {
	return o.value
}

func (o AnalyticsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AnalyticsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = true
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *AnalyticsOption) Equal(o2 *AnalyticsOption) bool {
	if o2 == nil {
		return o.value == true
	}
	return o.value == o2.value
}

func AnalyticsEqual(o1, o2 *AnalyticsOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
