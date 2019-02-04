package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestSumOrFiltersScores(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.SumOrFiltersScoresOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.SumOrFiltersScores(false),
		},
		{
			opts:     []interface{}{opt.SumOrFiltersScores(true)},
			expected: opt.SumOrFiltersScores(true),
		},
		{
			opts:     []interface{}{opt.SumOrFiltersScores(false)},
			expected: opt.SumOrFiltersScores(false),
		},
	} {
		var (
			in  = ExtractSumOrFiltersScores(c.opts...)
			out opt.SumOrFiltersScoresOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
