package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractTypoTolerance(opts ...interface{}) *opt.TypoToleranceOption {
	for _, o := range opts {
		if v, ok := o.(opt.TypoToleranceOption); ok {
			return &v
		}
	}
	return nil
}
