// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import "time"

// IContext is the context for service
type IContext interface {
	Log(message string)
	Error(err error, servers string)
	Param(name string) string
	QueryParam(name string) string
	Response(responseCode int, responseData interface{})
	ReadInput() string
	ReadInputs() []string

	// Time
	Now() time.Time

	// Dependency
	Cacher(server string) ICacher
	Indexer(servers string) IIndexer
	Producer(servers string) IProducer
	MQ(servers string) IMQ
	Requester(baseURL string, timeout time.Duration) IRequester
}
