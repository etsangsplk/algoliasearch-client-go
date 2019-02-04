package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAroundPrecision(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.AroundPrecisionOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AroundPrecisionOption{},
		},
		{
			opts:     []interface{}{opt.AroundPrecision(1)},
			expected: opt.AroundPrecision(1),
		},
		{
			opts:     []interface{}{opt.AroundPrecision(42)},
			expected: opt.AroundPrecision(42),
		},
	} {
		var (
			in  = ExtractAroundPrecision(c.opts...)
			out opt.AroundPrecisionOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
