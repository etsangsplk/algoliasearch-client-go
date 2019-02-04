package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractEnableRules(opts ...interface{}) *opt.EnableRulesOption {
	for _, o := range opts {
		if v, ok := o.(opt.EnableRulesOption); ok {
			return &v
		}
	}
	return nil
}
