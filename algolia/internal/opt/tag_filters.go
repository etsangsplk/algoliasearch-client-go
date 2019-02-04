package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractTagFilters(opts ...interface{}) *opt.TagFiltersOption {
	for _, o := range opts {
		if v, ok := o.(opt.TagFiltersOption); ok {
			return &v
		}
	}
	return nil
}
