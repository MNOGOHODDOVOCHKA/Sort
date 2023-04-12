package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	arr_size = 1000
	max_num  = 200
)

func main() {
	rand.Seed(time.Now().Unix())

	mas := NewSortArray(arr_size)
	mas.Fill(max_num)
	fmt.Println("Array before sorting:")
	mas.Print(false)

	mas.quick_sort()
	fmt.Println("Array after sorting:")
	mas.Print(true)
}
