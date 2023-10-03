package main

import "github.com/3dsinteractive/testify/mock"

type RandomMock struct {
	mock.Mock
}

func NewRandomMock() *RandomMock {
	return &RandomMock{}
}

func (rnd *RandomMock) Random() string {
	args := rnd.Called()
	return args.String(0)
}
