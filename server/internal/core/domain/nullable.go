package domain

type Nulleable[T any] struct {
	Value *T
	Set   bool
}
