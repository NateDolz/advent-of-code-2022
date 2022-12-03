package util

import (
	"fmt"
	"math/rand"
	"testing"

	"golang.org/x/exp/slices"
)

type uniqueTest[T comparable] struct {
	slice    []T
	expected []T
}

type benchmarkTest struct {
	size  int
	slice []int
}

var intTests = []uniqueTest[int]{
	{slice: []int{1, 3, 3, 2}, expected: []int{1, 3, 2}},
	{slice: []int{10, 3, 1, 2}, expected: []int{10, 3, 1, 2}},
	{slice: []int{1, 1, 1, 1}, expected: []int{1}},
}

var stringTests = []uniqueTest[string]{
	{slice: []string{"a", "b", "b", "c"}, expected: []string{"a", "b", "c"}},
	{slice: []string{"a", "b", "c", "d"}, expected: []string{"a", "b", "c", "d"}},
	{slice: []string{"a", "a", "a", "a"}, expected: []string{"a"}},
}

var benchmarkTests = []benchmarkTest{
	{size: 1000, slice: rand.Perm(1000)},
	{size: 10000, slice: rand.Perm(10000)},
	{size: 100000, slice: rand.Perm(100000)},
}

func TestUniqueInt(t *testing.T) {
	for _, test := range intTests {
		unique := Unique(test.slice)
		slices.Sort(unique)
		slices.Sort(test.expected)
		if compareRes := slices.Compare(unique, test.expected); compareRes != 0 {
			t.Error("Non Unique Result returned")
		}
	}
}

func TestUniqueString(t *testing.T) {
	for _, test := range stringTests {
		unique := Unique(test.slice)
		slices.Sort(unique)
		slices.Sort(test.expected)
		if compareRes := slices.Compare(unique, test.expected); compareRes != 0 {
			t.Error("Non Unique Result returned")
		}
	}
}

func TestUniqueV2Int(t *testing.T) {
	for _, test := range intTests {
		unique := UniqueV2(test.slice)
		slices.Sort(unique)
		slices.Sort(test.expected)
		if compareRes := slices.Compare(unique, test.expected); compareRes != 0 {
			t.Error("Non Unique Result returned")
		}
	}
}

func TestUniqueV2String(t *testing.T) {
	for _, test := range stringTests {
		unique := UniqueV2(test.slice)
		slices.Sort(unique)
		slices.Sort(test.expected)
		if compareRes := slices.Compare(unique, test.expected); compareRes != 0 {
			t.Error("Non Unique Result returned")
		}
	}
}

func BenchmarkUnique(b *testing.B) {
	for _, test := range benchmarkTests {
		b.Run(fmt.Sprintf("input_size_%d", test.size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Unique(test.slice)
			}
		})
	}
}

func BenchmarkUniqueV2(b *testing.B) {
	for _, test := range benchmarkTests {
		b.Run(fmt.Sprintf("input_size_%d", test.size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UniqueV2(test.slice)
			}
		})
	}
}
