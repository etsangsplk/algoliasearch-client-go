package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractAroundLatLngViaIP(opts ...interface{}) *opt.AroundLatLngViaIPOption {
	for _, o := range opts {
		if v, ok := o.(opt.AroundLatLngViaIPOption); ok {
			return &v
		}
	}
	return nil
}
