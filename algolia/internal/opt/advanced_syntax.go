package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractAdvancedSyntax(opts ...interface{}) *opt.AdvancedSyntaxOption {
	for _, o := range opts {
		if v, ok := o.(opt.AdvancedSyntaxOption); ok {
			return &v
		}
	}
	return nil
}
