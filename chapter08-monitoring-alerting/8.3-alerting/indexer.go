// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import (
	"context"
	"strings"

	"github.com/3dsinteractive/elastic"
)

// IIndexer is the interface to search the index
type IIndexer interface {
	Index(index string, id string, indoc interface{}) ( /*id*/ string, error)
	Query(index string, query string) ( /*results*/ []string, int64, error)
}

type Indexer struct {
	client  *elastic.Client
	servers string
	ms      IMicroservice
}

func NewIndexer(servers string, ms IMicroservice) *Indexer {
	return &Indexer{
		client:  nil,
		servers: servers,
		ms:      ms,
	}
}

func (idx *Indexer) Query(index string, query string) ( /*results*/ []string, int64, error) {

	client, err := idx.getClient(idx.servers)
	if err != nil {
		return nil, 0, err
	}

	// 7. Build with all options and send the query to elasticsearch
	svc := client.Search().
		Index(index).
		Source(query).
		IgnoreUnavailable(true)

	var res *elastic.SearchResult

	res, err = svc.Do(context.Background())
	if err != nil {
		// IF NOT 404, return error
		if !idx.isError404NotFound(err) {
			return nil, 0, err
		}

		return make([]string, 0), 0, nil
	}

	// Just prevent if result is nil, must not go next line
	if res == nil {
		return make([]string, 0), 0, nil
	}

	// create the results as array of string
	docs := make([]string, 0)
	if len(res.Hits.Hits) > 0 {
		count := len(res.Hits.Hits)
		docs = make([]string, count)
		for i, hit := range res.Hits.Hits {
			docs[i] = string(*hit.Source)
		}
	}

	totalHits := res.Hits.TotalHits

	return docs, totalHits, nil
}

func (idx *Indexer) Index(
	index string,
	id string,
	indoc interface{}) (string /*id*/, error) {

	client, err := idx.getClient(idx.servers)
	if err != nil {
		return "", err
	}

	cmd := client.Index().
		Index(index).
		Type(index).
		Id(id).
		BodyJson(indoc)

	var result *elastic.IndexResponse

	result, err = cmd.Do(context.Background())
	if err != nil {
		// If not 404, return error
		if !idx.isError404NotFound(err) {
			return "", err
		}
	}

	// Just prevent if result is nil, must not go next line
	if result == nil {
		return "", nil
	}

	return result.Id, nil
}

func (idx *Indexer) isError404NotFound(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(strings.ToLower(err.Error()),
		strings.ToLower(ElasticError404NotFound))
}

func (idx *Indexer) getClient(servers string) (*elastic.Client, error) {
	client := idx.client
	if client == nil {

		var c *elastic.Client

		var err error

		splitServers := strings.Split(servers, ",")
		// No user & password mean no basic auth required
		c, err = elastic.NewClient(
			elastic.SetURL(splitServers...),
			elastic.SetSniff(false),
			elastic.SetGzip(false),
		)
		if err != nil {
			return nil, err
		}

		client = c
		idx.client = c
	}

	return client, nil
}

const (
	// DONT CHANGE THIS VALUE, it depends on elasticsearch
	// ElasticError404NotFound error message when document not found
	ElasticError404NotFound     string = "elastic: Error 404 (Not Found)"
	ElasticErrorNoIndex         string = "elastic: Error 404 (Not Found): no such index"
	ElasticErrorVersionConflict string = "elastic: Error 409 (Conflict)"
	ElasticErrorEOF             string = "EOF"
)
