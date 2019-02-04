package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractSumOrFiltersScores(opts ...interface{}) *opt.SumOrFiltersScoresOption {
	for _, o := range opts {
		if v, ok := o.(opt.SumOrFiltersScoresOption); ok {
			return &v
		}
	}
	return nil
}
