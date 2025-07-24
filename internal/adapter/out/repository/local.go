package repository

import (
	"errors"
	"fmt"
)

type Local[T any] struct {
	storage map[string]*T
}

func NewLocal[T any]() *Local[T] {
	return &Local[T]{make(map[string]*T)}
}

func (l *Local[T]) Save(id string, elem *T) error {
	l.storage[id] = elem
	return nil
}

func (l *Local[T]) GetById(id string) (*T, error) {
	elem, ok := l.storage[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Not Found with ID [%s]", id))
	}
	return elem, nil
}