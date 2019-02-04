package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractSynonyms(opts ...interface{}) *opt.SynonymsOption {
	for _, o := range opts {
		if v, ok := o.(opt.SynonymsOption); ok {
			return &v
		}
	}
	return nil
}
