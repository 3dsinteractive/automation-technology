package main

import (
	"net/http"
	"testing"

	"github.com/3dsinteractive/testify/suite"
	"github.com/stretchr/testify/assert"
)

type MainTestSuite struct {
	suite.Suite
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}

func (ts *MainTestSuite) TestOnPostClient() {
	is := assert.New(ts.T())

	var ctx *ContextMock
	var cfg *ConfigMock
	var rnd *RandomMock
	var prod *ProducerMock
	var err error
	var status map[string]interface{}
	var citizen map[string]interface{}

	// Case: random start with 0
	ctx = NewContextMock()
	status = map[string]interface{}{
		"status": "failed",
	}
	ctx.On("Response", http.StatusOK, status)
	cfg = NewConfigMock()
	rnd = NewRandomMock()
	rnd.On("Random").Return("0000") // start with 0
	err = onPostClient(ctx, cfg, rnd)
	is.NoError(err)
	ctx.AssertExpectations(ts.T())

	// Case: random not start with 0
	ctx = NewContextMock()
	status = map[string]interface{}{
		"status":     "success",
		"citizen_id": "1111",
	}
	citizen = map[string]interface{}{
		"citizen_id": "1111",
	}
	ctx.On("Response", http.StatusOK, status)
	cfg = NewConfigMock()
	rnd = NewRandomMock()
	prod = NewProducerMock()
	cfg.On("CitizenRegisteredTopic").Return("mock-citizen-registered-topic")
	cfg.On("MQServers").Return("mock-mq-servers")
	ctx.On("Producer", "mock-mq-servers").Return(prod)
	prod.On("SendMessage", "mock-citizen-registered-topic", "", citizen).Return(nil)
	rnd.On("Random").Return("1111") // not start with 0
	err = onPostClient(ctx, cfg, rnd)
	is.NoError(err)

	ctx.AssertExpectations(ts.T())
	prod.AssertExpectations(ts.T())
	cfg.AssertExpectations(ts.T())

	// Case: random is empty
	ctx = NewContextMock()
	status = map[string]interface{}{
		"status": "failed",
	}
	ctx.On("Response", http.StatusOK, status)
	cfg = NewConfigMock()
	rnd = NewRandomMock()
	rnd.On("Random").Return("") // empty random
	err = onPostClient(ctx, cfg, rnd)
	is.NoError(err)

	ctx.AssertExpectations(ts.T())
}
