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

// Memory efficient Time inefficient
func Unique[T comparable](slice []T) (unique []T) {
	newSlice := make([]T, 0)
	for _, val := range slice {
		if idx := slices.Index(newSlice, val); idx == -1 {
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

func Intersect[T comparable](slice1 []T, slice2 []T) (union []T) {
	newSlice := make([]T, 0)

	unique1 := UniqueV2(slice1)
	unique2 := UniqueV2(slice2)

	for _, val := range unique1 {
		if idx := slices.Index(unique2, val); idx != -1 {
			newSlice = append(newSlice, val)
		}
	}

	return newSlice
}
