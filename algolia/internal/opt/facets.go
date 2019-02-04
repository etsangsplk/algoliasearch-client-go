package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractFacets(opts ...interface{}) *opt.FacetsOption {
	for _, o := range opts {
		if v, ok := o.(opt.FacetsOption); ok {
			return &v
		}
	}
	return nil
}
