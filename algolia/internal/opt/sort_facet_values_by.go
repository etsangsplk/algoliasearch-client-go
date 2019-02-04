package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractSortFacetValuesBy(opts ...interface{}) *opt.SortFacetValuesByOption {
	for _, o := range opts {
		if v, ok := o.(opt.SortFacetValuesByOption); ok {
			return &v
		}
	}
	return nil
}
