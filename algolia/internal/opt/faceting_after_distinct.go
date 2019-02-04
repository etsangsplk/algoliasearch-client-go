package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractFacetingAfterDistinct(opts ...interface{}) *opt.FacetingAfterDistinctOption {
	for _, o := range opts {
		if v, ok := o.(opt.FacetingAfterDistinctOption); ok {
			return &v
		}
	}
	return nil
}
