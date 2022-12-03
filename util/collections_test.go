package util

import (
	"fmt"
	"math/rand"
	"testing"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// #region unique test sets

type uniqueTest[T comparable] struct {
	slice    []T
	expected []T
}

type benchmarkUniqueTest struct {
	size  int
	slice []int
}

var intUniqueTests = []uniqueTest[int]{
	{slice: []int{1, 3, 3, 2}, expected: []int{1, 3, 2}},
	{slice: []int{10, 3, 1, 2}, expected: []int{10, 3, 1, 2}},
	{slice: []int{1, 1, 1, 1}, expected: []int{1}},
}

var stringUniqueTests = []uniqueTest[string]{
	{slice: []string{"a", "b", "b", "c"}, expected: []string{"a", "b", "c"}},
	{slice: []string{"a", "b", "c", "d"}, expected: []string{"a", "b", "c", "d"}},
	{slice: []string{"a", "a", "a", "a"}, expected: []string{"a"}},
}

var bnechmarkUniqueTests = []benchmarkUniqueTest{
	{size: 1000, slice: rand.Perm(1000)},
	{size: 10000, slice: rand.Perm(10000)},
	{size: 100000, slice: rand.Perm(100000)},
}

// #endregion

// #region min test sets

type minTest[T constraints.Ordered] struct {
	slice    []T
	expected T
}

var minIntTests = []minTest[int]{
	{slice: []int{1, 3, 3, 2}, expected: 1},
	{slice: []int{10, 3, 1, 2}, expected: 1},
	{slice: []int{1, 1, 1, 1}, expected: 1},
	{slice: []int{1000, 4000, 8}, expected: 8},
	{slice: rand.Perm(100), expected: 0},
}

var minStringTests = []minTest[string]{
	{slice: []string{"a", "c", "d"}, expected: "a"},
	{slice: []string{"z", "y", "x"}, expected: "x"},
	{slice: []string{"A", "c", "x"}, expected: "A"},
	{slice: []string{"Z", "a", "b"}, expected: "Z"},
}

// #endregion

// #region intersect test sets

type intersectTest[T comparable] struct {
	slice1   []T
	slice2   []T
	expected []T
}

var intersectIntTests = []intersectTest[int]{
	{slice1: []int{1, 5, 7}, slice2: []int{2, 4, 6}, expected: []int{}},
	{slice1: []int{1, 5, 7}, slice2: []int{1, 5, 7}, expected: []int{1, 5, 7}},
	{slice1: []int{1, 1, 5, 7}, slice2: []int{1, 1, 5, 7}, expected: []int{1, 5, 7}},
	{slice1: []int{1, 5, 7}, slice2: []int{2, 3, 7}, expected: []int{7}},
	{slice1: []int{1, 1, 2, 5, 7}, slice2: []int{1, 1, 2}, expected: []int{1, 2}},
}

var intersectStringTests = []intersectTest[string]{
	{slice1: []string{"a", "c", "e"}, slice2: []string{"b", "d", "f"}, expected: []string{}},
	{slice1: []string{"a", "b", "c"}, slice2: []string{"a", "b", "c"}, expected: []string{"a", "b", "c"}},
	{slice1: []string{"string"}, slice2: []string{"String"}, expected: []string{}},
	{slice1: []string{"a", "string", "hello", "world"}, slice2: []string{"hello", "World"}, expected: []string{"hello"}},
	{slice1: []string{"a", "a", "b", "c", "d"}, slice2: []string{"a", "a", "b"}, expected: []string{"a", "b"}},
}

// #endregion

// #region unique tests

func TestUnique(t *testing.T) {
	for _, test := range intUniqueTests {
		unique := Unique(test.slice)
		slices.Sort(unique)
		slices.Sort(test.expected)
		if compareRes := slices.Compare(unique, test.expected); compareRes != 0 {
			t.Error("Non Unique Result returned")
		}
	}

	for _, test := range stringUniqueTests {
		unique := Unique(test.slice)
		slices.Sort(unique)
		slices.Sort(test.expected)
		if compareRes := slices.Compare(unique, test.expected); compareRes != 0 {
			t.Error("Non Unique Result returned")
		}
	}
}

func TestUniqueV2(t *testing.T) {
	for _, test := range intUniqueTests {
		unique := UniqueV2(test.slice)
		slices.Sort(unique)
		slices.Sort(test.expected)
		if compareRes := slices.Compare(unique, test.expected); compareRes != 0 {
			t.Error("Non Unique Result returned")
		}
	}

	for _, test := range stringUniqueTests {
		unique := UniqueV2(test.slice)
		slices.Sort(unique)
		slices.Sort(test.expected)
		if compareRes := slices.Compare(unique, test.expected); compareRes != 0 {
			t.Error("Non Unique Result returned")
		}
	}
}

func BenchmarkUnique(b *testing.B) {
	for _, test := range bnechmarkUniqueTests {
		b.Run(fmt.Sprintf("input_size_%d", test.size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Unique(test.slice)
			}
		})
	}
}

func BenchmarkUniqueV2(b *testing.B) {
	for _, test := range bnechmarkUniqueTests {
		b.Run(fmt.Sprintf("input_size_%d", test.size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UniqueV2(test.slice)
			}
		})
	}
}

// #endregion

// #region min tests

func TestMinInt(t *testing.T) {
	for _, test := range minIntTests {
		if _, min := Min(test.slice); min != test.expected {
			fmt.Println(test.slice)
			t.Errorf("Expected min value %d, got %d", test.expected, min)
		}
	}
}

func TestMinString(t *testing.T) {
	for _, test := range minStringTests {
		if _, min := Min(test.slice); min != test.expected {
			fmt.Println(test.slice)
			t.Errorf("Expected min value %s, got %s", test.expected, min)
		}
	}
}

// #endregion

// #region intersect tests

func TestIntersectInts(t *testing.T) {
	for _, test := range intersectIntTests {
		result := Intersect(test.slice1, test.slice2)
		slices.Sort(result)
		slices.Sort(test.expected)
		if res := slices.Compare(result, test.expected); res != .0 {
			fmt.Println(test)
			t.Errorf("Expected value %v, got %v", test.expected, result)
		}
	}
}

func TestIntersectStrings(t *testing.T) {
	for _, test := range intersectStringTests {
		result := Intersect(test.slice1, test.slice2)
		slices.Sort(result)
		slices.Sort(test.expected)
		if res := slices.Compare(result, test.expected); res != .0 {
			fmt.Println(test)
			t.Errorf("Expected value %v, got %v", test.expected, result)
		}
	}
}

// #endregion intersect tests
