package main

import (
	"fmt"
)

// ============================================================================
// we make an interface that defines the strategy behaviour
type Sorter interface {
	Sort([]int) []int
}

// ============================================================================
// one strategy
type BubbleSort struct{}

// btw, slices are inherently pass by refference so no need for a point
func (bubble *BubbleSort) Sort(arr []int) []int {
	//implement
	return arr
}

// ============================================================================
// another strategy
type InsertionSort struct{}

func (insertion *InsertionSort) Sort(arr []int) []int {
	// implement
	return arr
}

// ============================================================================
// the thing that can accept different strategies
// has a field of some tupe that implements the sorter interface
type Context struct {
	sorter Sorter
}

func (context *Context) SetSorter(sorter Sorter) {
	context.sorter = sorter
}

func (context *Context) ExecuteSort(arr []int) []int {
	return context.sorter.Sort(arr)
}

// ============================================================================
// choose the sorting strategy
func main() {
	arr := []int{4, 3, 2, 1}
	fmt.Println(arr)
}
