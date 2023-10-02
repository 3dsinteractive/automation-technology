package main

import "github.com/3dsinteractive/ksuid"

// IUUIDGen is UUID generator
type IUUIDGen interface {
	NewUUID() string
	IsValid(id string) error
}

// UUIDGen is struct implement IUUIDGen
type UUIDGen struct {
}

// NewUUIDGen return new UUIDGen
func NewUUIDGen() *UUIDGen {
	return &UUIDGen{}
}

// NewUUID return new UUID as string
func (gen *UUIDGen) NewUUID() string {
	id := ksuid.New()
	return id.String()
}

// IsValid validate given id
func (gen *UUIDGen) IsValid(id string) error {
	_, err := ksuid.Parse(id)
	if err != nil {
		return err
	}

	return nil
}
