// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type ClickAnalyticsOption struct {
	value bool
}

func ClickAnalytics(v bool) ClickAnalyticsOption {
	return ClickAnalyticsOption{v}
}

func (o ClickAnalyticsOption) Get() bool {
	return o.value
}

func (o ClickAnalyticsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ClickAnalyticsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}
