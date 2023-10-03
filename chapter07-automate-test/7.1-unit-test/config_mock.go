// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import (
	"github.com/3dsinteractive/testify/mock"
)

// NewConfigMock implement IConfig
type ConfigMock struct {
	mock.Mock
}

// NewConfigMock return new config instance
func NewConfigMock() *ConfigMock {
	return &ConfigMock{}
}

// ServiceID return ID of service
func (cfg *ConfigMock) ServiceID() string {
	args := cfg.Called()
	return args.String(0)
}

// CacheServer return redis server
func (cfg *ConfigMock) CacheServer() string {
	args := cfg.Called()
	return args.String(0)
}

// MQServers return Kafka servers
func (cfg *ConfigMock) MQServers() string {
	args := cfg.Called()
	return args.String(0)
}

// CitizenRegisteredTopic return topic name for registered event
func (cfg *ConfigMock) CitizenRegisteredTopic() string {
	args := cfg.Called()
	return args.String(0)
}

// CitizenConfirmedTopic return topic name for confirmed event
func (cfg *ConfigMock) CitizenConfirmedTopic() string {
	args := cfg.Called()
	return args.String(0)
}

// CitizenValidationAPI return API to validate citizen information
func (cfg *ConfigMock) CitizenValidationAPI() string {
	args := cfg.Called()
	return args.String(0)
}

// CitizenDeliveryAPI return API to request delivery citizen ID card
func (cfg *ConfigMock) CitizenDeliveryAPI() string {
	args := cfg.Called()
	return args.String(0)
}

// BatchDeliverAPI return API to batch delivery citizen ID card
func (cfg *ConfigMock) BatchDeliverAPI() string {
	args := cfg.Called()
	return args.String(0)
}
