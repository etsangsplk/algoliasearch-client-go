package search

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/algolia/algoliasearch-client-go/algolia/rand"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

func defaultHosts(appID string) (hosts []*transport.StatefulHost) {
	hosts = append(hosts, transport.NewStatefulHost(appID+"-dsn.algolia.net", call.IsRead))
	hosts = append(hosts, transport.NewStatefulHost(appID+".algolia.net", call.IsWrite))
	hosts = append(hosts, transport.Shuffle(
		[]*transport.StatefulHost{
			transport.NewStatefulHost(appID+"-1.algolianet.com", call.IsReadWrite),
			transport.NewStatefulHost(appID+"-2.algolianet.com", call.IsReadWrite),
			transport.NewStatefulHost(appID+"-3.algolianet.com", call.IsReadWrite),
		},
	)...)
	return
}

func noWait(_ int) error {
	return nil
}

func waitWithRetry(f func() (bool, error)) error {
	var maxDuration = time.Second

	for {
		done, err := f()
		if done {
			return err
		}

		sleepDuration := rand.Duration(maxDuration)
		time.Sleep(sleepDuration)

		// Increase the upper boundary used to generate the sleep duration
		if maxDuration < 10*time.Minute {
			maxDuration *= 2
			if maxDuration > 10*time.Minute {
				maxDuration = 10 * time.Minute
			}
		}
	}
}

func getObjectID(object interface{}) (string, bool) {
	data, err := json.Marshal(object)
	if err != nil {
		return "", false
	}
	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return "", false
	}
	objectID, ok := m["objectID"]
	if !ok {
		return "", false
	}

	switch v := objectID.(type) {
	case string:
		return v, true
	case float64:
		return fmt.Sprintf("%d", int(v)), true
	default:
		return "", false
	}
}

func hasObjectID(object interface{}) bool {
	_, ok := getObjectID(object)
	return ok
}
