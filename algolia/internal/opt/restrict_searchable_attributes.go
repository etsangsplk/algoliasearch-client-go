package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractRestrictSearchableAttributes(opts ...interface{}) *opt.RestrictSearchableAttributesOption {
	for _, o := range opts {
		if v, ok := o.(opt.RestrictSearchableAttributesOption); ok {
			return &v
		}
	}
	return nil
}
