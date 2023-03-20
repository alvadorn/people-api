package domain

import (
	"context"
)

type Person struct {
	Name             string
	ShortDescription string
}

type PersonGetterByName interface {
	GetByName(ctx context.Context, name string) (*Person, error)
}

type PeopleRepository interface {
	PersonGetterByName
}
