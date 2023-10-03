package main

type IRandom interface {
	Random() string
}

type Random struct{}

func NewRandom() *Random {
	return &Random{}
}

func (r *Random) Random() string {
	return randString()
}
