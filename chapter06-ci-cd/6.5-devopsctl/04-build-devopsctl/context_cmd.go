// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// CmdContext implement IContext it is context for Consumer
type CmdContext struct {
	ms   *Microservice
	args map[string]string
}

// NewCmdContext is the constructor function for CmdContext
func NewCmdContext(ms *Microservice, args map[string]string) *CmdContext {
	return &CmdContext{
		ms:   ms,
		args: args,
	}
}

// Log will log a message
func (ctx *CmdContext) Log(message string) {
	_, fn, line, _ := runtime.Caller(1)
	fns := strings.Split(fn, "/")
	fmt.Println("CMD:", fns[len(fns)-1], line, message)
}

// Param return parameter by name (empty in scheduler)
func (ctx *CmdContext) Param(name string) string {
	if len(ctx.args) == 0 {
		return ""
	}

	val, ok := ctx.args[name]
	if ok {
		// When define non required flag with empty default value, the library need one value instead of accept empty default value
		// So we replace empty value with this UUID, and will replace it with empty value when get with ctx.Param("name")
		// We expect you don't use this value as the value of flag, If you use this value, please change!
		// EMPTY_DEFAULT_VALUE = "22c211c6-8774-11eb-8dcd-0242ac130003"
		if val == EMPTY_DEFAULT_VALUE {
			return ""
		}
		return val
	}

	return ""
}

// QueryParam return empty in scheduler
func (ctx *CmdContext) QueryParam(name string) string {
	return ""
}

// ReadInput return message (return empty in scheduler)
func (ctx *CmdContext) ReadInput() string {
	return ""
}

// ReadInputs return messages in batch (return nil in scheduler)
func (ctx *CmdContext) ReadInputs() []string {
	return nil
}

// Response return response to client
func (ctx *CmdContext) Response(responseCode int, responseData interface{}) {
	return
}

// Now return now
func (ctx *CmdContext) Now() time.Time {
	return time.Now()
}

// Cacher return cacher
func (ctx *CmdContext) Cacher(server string) ICacher {
	return ctx.ms.getCacher(server)
}

// Producer return producer
func (ctx *CmdContext) Producer(servers string) IProducer {
	return ctx.ms.getProducer(servers)
}

// MQ return MQ
func (ctx *CmdContext) MQ(servers string) IMQ {
	return NewMQ(servers, ctx.ms)
}

// Requester return Requester
func (ctx *CmdContext) Requester(baseURL string, timeout time.Duration) IRequester {
	return NewRequester(baseURL, timeout, ctx.ms)
}
