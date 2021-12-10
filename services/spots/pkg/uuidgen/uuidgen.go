package uuidgen

import "github.com/google/uuid"

type UUIDGen interface {
	New() string
}

type uuidgen struct{}

func New() UUIDGen {
	return &uuidgen{}
}

func (u uuidgen) New() string {
	return uuid.NewString()
}
