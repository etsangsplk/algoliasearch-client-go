package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractAttributesToSnippet(opts ...interface{}) *opt.AttributesToSnippetOption {
	for _, o := range opts {
		if v, ok := o.(opt.AttributesToSnippetOption); ok {
			return &v
		}
	}
	return nil
}
