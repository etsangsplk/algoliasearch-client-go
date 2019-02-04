package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAroundLatLngViaIP(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.AroundLatLngViaIPOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AroundLatLngViaIP(false),
		},
		{
			opts:     []interface{}{opt.AroundLatLngViaIP(true)},
			expected: opt.AroundLatLngViaIP(true),
		},
		{
			opts:     []interface{}{opt.AroundLatLngViaIP(false)},
			expected: opt.AroundLatLngViaIP(false),
		},
	} {
		var (
			in  = ExtractAroundLatLngViaIP(c.opts...)
			out opt.AroundLatLngViaIPOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
