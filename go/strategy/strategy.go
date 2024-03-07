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
func main() {
	// **********
	// this doesnt make sense if the array does not belong to the context
	// otherwise we could just pass the data structure into the function
	// like a normal procedural paradigm
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	// choose the sorting strategy
	bubbleSort := &BubbleSort{}
	// insertionSort := &InsertionSort{}

	context := &Context{}
	context.SetSorter(bubbleSort)
	context.ExecuteSort(arr)

	fmt.Println(arr)
}
