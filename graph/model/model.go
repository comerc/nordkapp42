package model

type Model[T comparable] interface {
	GetID() T
}
