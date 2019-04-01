package search

import (
	"fmt"
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	iopt "github.com/algolia/algoliasearch-client-go/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

type Index struct {
	appID        string
	name         string
	maxBatchSize int
	client       *Client
	transport    *transport.Transport
}

func newIndex(client *Client, name string) *Index {
	return &Index{
		appID:        client.appID,
		client:       client,
		name:         name,
		maxBatchSize: client.maxBatchSize,
		transport:    client.transport,
	}
}

func (i *Index) path(format string, a ...interface{}) string {
	prefix := fmt.Sprintf("/1/indexes/%s", i.name)
	suffix := fmt.Sprintf(format, a...)
	return prefix + suffix
}

func (i *Index) WaitTask(taskID int) error {
	return waitWithRetry(func() (bool, error) {
		res, err := i.GetStatus(taskID)
		if err != nil {
			return true, err
		}
		if res.Status == "published" {
			return true, nil
		}
		return false, nil
	})
}

func (i *Index) operation(destination, op string, opts ...interface{}) (res UpdateTaskRes, err error) {
	var scopes []string
	if opt := iopt.ExtractScopes(opts...); opt != nil {
		scopes = opt.Get()
	}
	req := IndexOperation{
		Destination: destination,
		Operation:   op,
		Scopes:      scopes,
	}
	path := i.path("/operation")
	err = i.transport.Request(&res, http.MethodPost, path, req, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

func (i *Index) GetAppID() string {
	return i.appID
}

func (i *Index) Clear(opts ...interface{}) (res UpdateTaskRes, err error) {
	path := i.path("/clear")
	err = i.transport.Request(&res, http.MethodPost, path, nil, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

func (i *Index) Delete(opts ...interface{}) (res DeleteTaskRes, err error) {
	path := i.path("")
	err = i.transport.Request(&res, http.MethodDelete, path, nil, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

func (i *Index) GetStatus(taskID int) (res TaskStatusRes, err error) {
	path := i.path("/task/%d", taskID)
	err = i.transport.Request(&res, http.MethodGet, path, nil, call.Read)
	return
}
