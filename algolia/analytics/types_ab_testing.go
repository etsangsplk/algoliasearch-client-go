package analytics

import (
	"encoding/json"
	"fmt"
	"time"
)

type ABTest struct {
	Name     string
	Variants []Variant
	EndAt    time.Time
}

const ISO8601 = "2006-01-02T15:04:05Z"

func (abTest ABTest) MarshalJSON() ([]byte, error) {
	jsonVariants, err := json.Marshal(abTest.Variants)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal variants: %s", err)
	}
	jsonAbTest := fmt.Sprintf(
		`{"name":"%s","variants":%s,"endAt":"%s"}`,
		abTest.Name,
		jsonVariants,
		abTest.EndAt.In(time.UTC).Format(ISO8601),
	)
	return []byte(jsonAbTest), nil
}

type Variant struct {
	Index                  string                 `json:"index"`
	TrafficPercentage      int                    `json:"trafficPercentage"`
	Description            string                 `json:"description,omitempty"`
	CustomSearchParameters map[string]interface{} `json:"customSearchParameters,omitempty"` // TODO: Check type
}
