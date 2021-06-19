package model

import "github.com/google/uuid"

type Id struct {
	value string
}

func NewId() *Id {
	value := uuid.NewString()

	return &Id{value}
}

func ValueOf(value string) (*Id, error) {
	if _, err := uuid.Parse(value); err != nil {
		return nil, err
	}

	return &Id{value}, nil
}

func (id *Id) Value() string {
	return id.value
}
