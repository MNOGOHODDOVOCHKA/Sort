package main

import "math"

func (s *SortArray) bubble_sort() {
	s.comparisons = 0
	s.swaps = 0

	for i := 1; i < s.Len(); i++ {
		for j := 0; j < s.Len()-i; j++ {
			s.comparisons++
			if s.array[j] > s.array[j+1] {
				s.swaps++
				s.array[j], s.array[j+1] = s.array[j+1], s.array[j]
			}
		}
	}
}

func (s *SortArray) shaker_sort() {
	s.comparisons = 0
	s.swaps = 0

	l_corner := -1
	r_corner := s.Len()
	forward := true
	for l_corner != r_corner {
		if forward {
			for i := l_corner + 1; i < r_corner-1; i++ {
				s.comparisons++
				if s.array[i] > s.array[i+1] {
					s.swaps++
					s.array[i], s.array[i+1] = s.array[i+1], s.array[i]
				}
			}
			forward = false
			r_corner--
		} else {
			for i := r_corner - 1; i > l_corner+1; i-- {
				s.comparisons++
				if s.array[i] < s.array[i-1] {
					s.swaps++
					s.array[i], s.array[i-1] = s.array[i-1], s.array[i]
				}
			}
			forward = true
			l_corner++
		}
	}
}

func (s *SortArray) insertion_sort() {
	s.comparisons = 0
	s.swaps = 0

	res := make([]int, s.Len())
	var newpos int

	for i := 0; i < s.Len(); i++ {
		newpos = 0
		for j := len(res) - 1; j >= 0; j-- {
			newpos = j
			s.comparisons++
			if res[j] < s.array[i] {
				break
			}
		}

		for k := 0; k < newpos; k++ {
			s.swaps++
			res[k] = res[k+1]
		}
		s.swaps++
		res[newpos] = s.array[i]
	}
	copy(s.array, res)
}

func (s *SortArray) selection_sort() {
	s.comparisons = 0
	s.swaps = 0

	for i := 0; i < s.Len(); i++ {
		min := s.array[i]
		index := i
		for j := i + 1; j < s.Len(); j++ {
			s.comparisons++
			if min > s.array[j] {
				min = s.array[j]
				index = j
			}
		}
		s.swaps++
		s.array[i], s.array[index] = s.array[index], s.array[i]
	}
}

var comparisons int
var swaps int

type node struct {
	parentInd int
	isLeaf    bool
	value     int
}

func makeTree(array []int) []node {
	heap := make([]node, len(array))
	curLvl := 0
	curLvlMax := 1
	curParentId := 0
	curChildCount := 0

	for i := 0; i < len(array); i++ {
		heap[i].value = array[i]
		heap[i].isLeaf = true
		if i == 0 {
			heap[i].parentInd = -1
			continue
		}

		if i >= curLvlMax {
			curLvl++
			curLvlMax += int(math.Pow(2, float64(curLvl)))
		}
		if curChildCount >= 2 {
			curParentId++
			curChildCount = 0
		}
		heap[i].parentInd = curParentId
		heap[curParentId].isLeaf = false
		curChildCount++
	}

	return heap
}

func getChilds(heap []node, parentId int) (int, int) {
	leftChildId := -1
	rightChildId := -1

	for i := 0; i <= len(heap)-1; i++ {
		if heap[i].parentInd == parentId {
			leftChildId = i
			break
		}
	}
	for j := len(heap) - 1; j >= 0; j-- {
		if heap[j].parentInd == parentId {
			rightChildId = j
			break
		}
	}

	return leftChildId, rightChildId
}

func sortNode(heap []node, nodeId int) {
	lch, rch := getChilds(heap, nodeId)
	//2 childs
	if lch >= 0 && rch >= 0 {
		comparisons++
		if heap[lch].value < heap[rch].value {
			swaps++
			heap[lch].value, heap[rch].value = heap[rch].value, heap[lch].value
		}
		comparisons++
		if heap[lch].value > heap[nodeId].value {
			swaps++
			heap[lch].value, heap[nodeId].value = heap[nodeId].value, heap[lch].value

			comparisons++
			if heap[lch].value < heap[rch].value {
				swaps++
				heap[lch].value, heap[rch].value = heap[rch].value, heap[lch].value
			}
		}
	} else {
		//1 child
		if lch >= 0 && lch == rch {
			comparisons++
			if heap[lch].value > heap[nodeId].value {
				swaps++
				heap[lch].value, heap[nodeId].value = heap[nodeId].value, heap[lch].value
			}
		}
	}
}

func (s *SortArray) heap_sort() {
	comparisons = 0
	swaps = 0

	//making heap
	heap := makeTree(s.array)
	//making sort
	id := 0
	for ; len(heap) > 0; heap = heap[:len(heap)-1] {
		for j := len(heap) - 1; j >= 0; j-- {
			if !heap[j].isLeaf {
				sortNode(heap, j)
			}
		}
		s.array[id] = heap[0].value
		heap[0].value = heap[len(heap)-1].value
		id++
	}

	s.comparisons = comparisons
	s.swaps = swaps
}

func sort_arr_part(arr_part []int) {
	if len(arr_part) == 0 {
		return
	}

	base := len(arr_part) - 1
	for i := 0; i < base; i++ {
		comparisons++
		if arr_part[i] > arr_part[base] {
			swaps++
			el := arr_part[i]
			p1 := arr_part[:i]
			p2 := arr_part[i+1:]
			arr_part = append(p1, p2...)
			arr_part = append(arr_part, el)
			base--
			i--
		}
	}

	sort_arr_part(arr_part[:base])
	sort_arr_part(arr_part[base+1:])
}

func (s *SortArray) quick_sort() {
	comparisons = 0
	swaps = 0

	sort_arr_part(s.array)

	s.comparisons = comparisons
	s.swaps = swaps
}
