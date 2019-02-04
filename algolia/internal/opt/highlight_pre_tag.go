package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractHighlightPreTag(opts ...interface{}) *opt.HighlightPreTagOption {
	for _, o := range opts {
		if v, ok := o.(opt.HighlightPreTagOption); ok {
			return &v
		}
	}
	return nil
}
