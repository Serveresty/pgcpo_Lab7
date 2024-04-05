package main

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
)

func TestRandomData(t *testing.T) {
	N := rand.Intn(10-1) + 1
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(100)
	}

	arr = sorting(arr)

	if !isSorted(arr) {
		t.Errorf("Error: not sorted")
	}
}

func TestEmptyArr(t *testing.T) {
	var arr []int

	arr1 := sorting(arr)

	if !reflect.DeepEqual(arr, arr1) {
		t.Errorf("Error: not sorted")
	}
}

func TestAlreadySorted(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := sorting(arr)

	if !reflect.DeepEqual(arr1, arr2) {
		t.Errorf("Error")
	}
}

func TestLessThanZero(t *testing.T) {
	arr := []int{-5, -4, -3, -2, -1}
	arr1 := sorting(arr)

	if !isSorted(arr1) {
		t.Errorf("Error: not sorted")
	}
}

func TestMixedNumbers(t *testing.T) {
	arr := []int{-5, 4, -3, 2, -1}
	arr1 := sorting(arr)

	if !isSorted(arr1) {
		t.Errorf("Error: not sorted")
	}
}

func TestSingleElement(t *testing.T) {
	arr := []int{2}
	arr1 := sorting(arr)

	if !isSorted(arr1) {
		t.Errorf("Error: not sorted")
	}
}

func TestRepeatElements(t *testing.T) {
	arr := []int{2, 3, 2, 5, 3}
	arr1 := sorting(arr)

	if !isSorted(arr1) {
		t.Errorf("Error: not sorted")
	}
}

func TestLargeNumbers(t *testing.T) {
	maxInt := math.MaxInt
	arr := []int{maxInt, maxInt / 10, maxInt / 100, maxInt / 2, maxInt / 5}

	arr1 := sorting(arr)

	if !isSorted(arr1) {
		t.Errorf("Error: not sorted")
	}
}

func TestLargeNegativeNumbers(t *testing.T) {
	maxInt := math.MaxInt
	arr := []int{-maxInt, -maxInt / 10, -maxInt / 100, -maxInt / 2, -maxInt / 5}

	arr1 := sorting(arr)

	if !isSorted(arr1) {
		t.Errorf("Error: not sorted")
	}
}
