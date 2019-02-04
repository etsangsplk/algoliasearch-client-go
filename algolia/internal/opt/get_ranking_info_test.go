package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestGetRankingInfo(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.GetRankingInfoOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.GetRankingInfo(false),
		},
		{
			opts:     []interface{}{opt.GetRankingInfo(true)},
			expected: opt.GetRankingInfo(true),
		},
		{
			opts:     []interface{}{opt.GetRankingInfo(false)},
			expected: opt.GetRankingInfo(false),
		},
	} {
		var (
			in  = ExtractGetRankingInfo(c.opts...)
			out opt.GetRankingInfoOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
