// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractDisableTypoToleranceOnWords(opts ...interface{}) *opt.DisableTypoToleranceOnWordsOption {
	for _, o := range opts {
		if v, ok := o.(*opt.DisableTypoToleranceOnWordsOption); ok {
			return v
		}
	}
	return nil
}
