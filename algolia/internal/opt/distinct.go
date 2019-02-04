package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractDistinct(opts ...interface{}) *opt.DistinctOption {
	for _, o := range opts {
		if v, ok := o.(opt.DistinctOption); ok {
			return &v
		}
	}
	return nil
}
