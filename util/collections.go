package util

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

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

// More efficient at lower slice sizes requires type to be ordered
func Unique[T constraints.Ordered](slice []T) (unique []T) {
	newSlice := make([]T, 0)
	for _, val := range slice {
		if _, found := slices.BinarySearch(newSlice, val); !found {
			newSlice = append(newSlice, val)
		}
	}
	return newSlice
}

// Time efficient Memory inefficient
func UniqueV2[T comparable](slice []T) (unique []T) {
	keyMap := make(map[T]int)
	for _, val := range slice {
		keyMap[val] = 1
	}
	return maps.Keys(keyMap)
}

func Intersect[T constraints.Ordered](slice1 []T, slice2 []T) (union []T) {
	newSlice := make([]T, 0)

	unique1 := Unique(slice1)
	unique2 := Unique(slice2)

	for _, val := range unique1 {
		if idx := slices.Index(unique2, val); idx != -1 {
			newSlice = append(newSlice, val)
		}
	}

	return newSlice
}
