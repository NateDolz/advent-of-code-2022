package util

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](slice []T) (idx int, value T) {
	min := slice[0]
	minIdx := 0
	for idx, curr := range slice {
		if idx == 0 || curr < min {
			min = curr
			minIdx = idx
		}
	}
	return minIdx, min
}
