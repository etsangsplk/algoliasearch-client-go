package search_client

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/cts"
	"github.com/stretchr/testify/require"
)

func TestGetLogs(t *testing.T) {
	t.Parallel()
	c := cts.InitSearchClient1(t)

	for i := 0; i < 2; i++ {
		_, err := c.ListIndexes()
		require.NoError(t, err)
	}

	res, err := c.GetLogs(
		opt.Offset(0),
		opt.Length(2),
		opt.Type("all"),
	)
	require.NoError(t, err)
	require.Len(t, res.Logs, 2)
}
