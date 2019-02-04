package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestSynonyms(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.SynonymsOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.Synonyms(true),
		},
		{
			opts:     []interface{}{opt.Synonyms(true)},
			expected: opt.Synonyms(true),
		},
		{
			opts:     []interface{}{opt.Synonyms(false)},
			expected: opt.Synonyms(false),
		},
	} {
		var (
			in  = ExtractSynonyms(c.opts...)
			out opt.SynonymsOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
