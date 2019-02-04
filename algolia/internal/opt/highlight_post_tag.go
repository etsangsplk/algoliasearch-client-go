package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractHighlightPostTag(opts ...interface{}) *opt.HighlightPostTagOption {
	for _, o := range opts {
		if v, ok := o.(opt.HighlightPostTagOption); ok {
			return &v
		}
	}
	return nil
}
