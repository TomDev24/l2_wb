package main

import (
	"fmt"
)

type Sorter interface {
	Sort([]int) []int
}

type BubbleSort struct{}

func (bs *BubbleSort) Sort(arr []int) []int {
	// Bubble Sort implementation
	return arr
}

type InsertionSort struct{}

func (is *InsertionSort) Sort(arr []int) []int {
	// Insertion Sort implementation
	return arr
}

type Context struct {
	sorter Sorter
}

func (c *Context) SetSorter(sorter Sorter) {
	c.sorter = sorter
}

func (c *Context) ExecuteSort(arr []int) []int {
	return c.sorter.Sort(arr)
}

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	bubbleSort := &BubbleSort{}
	insertionSort := &InsertionSort{}

	context := &Context{}
	context.SetSorter(bubbleSort)
	fmt.Println("Bubble Sort:", context.ExecuteSort(arr))

	context.SetSorter(insertionSort)
	fmt.Println("Insertion Sort:", context.ExecuteSort(arr))
}
