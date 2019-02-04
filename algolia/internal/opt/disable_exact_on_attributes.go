package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractDisableExactOnAttributes(opts ...interface{}) *opt.DisableExactOnAttributesOption {
	for _, o := range opts {
		if v, ok := o.(opt.DisableExactOnAttributesOption); ok {
			return &v
		}
	}
	return nil
}
