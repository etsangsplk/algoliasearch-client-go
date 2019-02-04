package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractAttributesToHighlight(opts ...interface{}) *opt.AttributesToHighlightOption {
	for _, o := range opts {
		if v, ok := o.(opt.AttributesToHighlightOption); ok {
			return &v
		}
	}
	return nil
}
