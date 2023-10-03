// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import (
	"time"

	"github.com/3dsinteractive/testify/mock"
)

type ContextMock struct {
	mock.Mock
}

func NewContextMock() *ContextMock {
	return &ContextMock{}
}

func (ctx *ContextMock) Log(message string) {
	ctx.Called(message)
}

func (ctx *ContextMock) Param(name string) string {
	args := ctx.Called(name)
	return args.String(0)
}

func (ctx *ContextMock) QueryParam(name string) string {
	args := ctx.Called(name)
	return args.String(0)
}

func (ctx *ContextMock) Response(responseCode int, responseData interface{}) {
	ctx.Called(responseCode, responseData)

}

func (ctx *ContextMock) ReadInput() string {
	args := ctx.Called()
	return args.String(0)
}

func (ctx *ContextMock) ReadInputs() []string {
	args := ctx.Called()
	return args.Get(0).([]string)
}

func (ctx *ContextMock) Now() time.Time {
	args := ctx.Called()
	return args.Get(0).(time.Time)
}

func (ctx *ContextMock) Cacher(server string) ICacher {
	args := ctx.Called(server)
	return args.Get(0).(ICacher)
}

func (ctx *ContextMock) Producer(servers string) IProducer {
	args := ctx.Called(servers)
	return args.Get(0).(IProducer)
}

func (ctx *ContextMock) MQ(servers string) IMQ {
	args := ctx.Called(servers)
	return args.Get(0).(IMQ)
}

func (ctx *ContextMock) Requester(baseURL string, timeout time.Duration) IRequester {
	args := ctx.Called(baseURL, timeout)
	return args.Get(0).(IRequester)
}
