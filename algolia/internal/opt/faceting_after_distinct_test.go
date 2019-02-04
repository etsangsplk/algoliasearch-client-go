package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestFacetingAfterDistinct(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.FacetingAfterDistinctOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.FacetingAfterDistinct(false),
		},
		{
			opts:     []interface{}{opt.FacetingAfterDistinct(true)},
			expected: opt.FacetingAfterDistinct(true),
		},
		{
			opts:     []interface{}{opt.FacetingAfterDistinct(false)},
			expected: opt.FacetingAfterDistinct(false),
		},
	} {
		var (
			in  = ExtractFacetingAfterDistinct(c.opts...)
			out opt.FacetingAfterDistinctOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
