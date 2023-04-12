package main

import (
	"fmt"
	"math/rand"
)

type SortArray struct {
	array       []int
	comparisons int
	swaps       int
}

func NewSortArray(size int) *SortArray {
	return &SortArray{array: make([]int, size)}
}

func (s *SortArray) Len() int {
	return len(s.array)
}

func (s *SortArray) Fill(max int) {
	for i := 0; i < len(s.array); i++ {
		s.array[i] = rand.Intn(max)
	}
}

func (s *SortArray) Print(print_stats bool) {
	for i := 0; i < len(s.array); i++ {
		fmt.Printf("%d ", s.array[i])
	}
	fmt.Println()
	if print_stats {
		fmt.Printf("Comparisons: %d\nSwaps: %d\n", s.comparisons, s.swaps)
	}
}
