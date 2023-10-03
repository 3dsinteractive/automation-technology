// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import "github.com/3dsinteractive/testify/mock"

// Producer implement IProducer, is the service to send message to Kafka
type ProducerMock struct {
	mock.Mock
}

// NewProducerMock return new instance of ProducerMock
func NewProducerMock() *ProducerMock {
	return &ProducerMock{}
}

// SendMessage send message to topic synchronously
func (p *ProducerMock) SendMessage(topic string, key string, message interface{}) error {
	args := p.Called(topic, key, message)
	return args.Error(0)
}

// Close the producer
func (p *ProducerMock) Close() error {
	args := p.Called()
	return args.Error(0)
}
