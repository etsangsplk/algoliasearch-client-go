// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractCamelCaseAttributes(opts ...interface{}) *opt.CamelCaseAttributesOption {
	for _, o := range opts {
		if v, ok := o.(*opt.CamelCaseAttributesOption); ok {
			return v
		}
	}
	return nil
}
