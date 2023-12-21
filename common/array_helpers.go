package common

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Sum[T Number](numbers []T) T {
	var sum T
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Reduce[T, U any](items []T, initialValue U, reducer func(U, T) U) U {
	result := initialValue
	for _, item := range items {
		result = reducer(result, item)
	}
	return result
}
