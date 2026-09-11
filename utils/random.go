package utils

import "math/rand/v2"

func Pick[T any](s []T) T {
	return s[rand.IntN(len(s))]
}

type Weighted[T any] struct {
	Value  T
	Weight float64
}

func PickWeighted[T any](items []Weighted[T]) T {
	total := 0.0
	for _, it := range items {
		total += it.Weight
	}

	r := rand.Float64() * total
	for _, it := range items {
		r -= it.Weight
		if r < 0 {
			return it.Value
		}
	}

	return items[len(items)-1].Value
}
