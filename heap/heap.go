package main

import "fmt"


type Heap struct {
	space []int

}

func NewHeap(size int) Heap {
	h := Heap{space : make([]int, size)}
	return h
}

func (h *Heap) Set (index int, value int)  {
	h.space[index] = value
}

func (h *Heap) Build ()  {
	startIndex := (len(h.space)/2) - 1

	for i := startIndex; i >= 0; i-- {
		h.Heapify(len(h.space), i)
	}
}

func (h *Heap) Heapify (n int, i int)  {

	largest := i
	l := i*2 + 1
	r := i*2 + 2
	
	if l < n && h.space[l]  > h.space[largest] {
		largest = l
	}

	if r < n && h.space[r]  > h.space[largest] {
		largest = r
	}

	if largest != i {
		h.space[i], h.space[largest] = h.space[largest], h.space[i]
		h.Heapify(n, largest)
	}
	
}

func (h *Heap) Print ()  {
	fmt.Print("[")
	
	for i, el := range h.space  {
		fmt.Print(el)
		if i != len(h.space)-1{
			fmt.Print(" ")
		}
	}
	
	fmt.Print("]")
}

func (h *Heap) HeapSort ()  {
	

	for i := len(h.space) - 1; i >= 0; i-- {
		h.space[0],h.space[i] = h.space[i], h.space[0]

		h.Heapify(i, 0)
	}
	
}

func main() {
	fmt.Printf("%s\n", "Heaping\n")

	hep := NewHeap(10)
	
	hep.Set(4, 15)
	hep.Set(0, 10)
	hep.Set(2, 1)
	hep.Set(3, 23)
	hep.Set(1, 4)
	hep.Set(5, 7)
	hep.Set(6, 43)
	hep.Set(7, 3)
	hep.Set(8, 9)
	hep.Set(9, 12)

	hep.Build()
	fmt.Print("Heap:")
	hep.Print()

	fmt.Println("")

	hep.HeapSort()
	fmt.Print("Sort:")
	hep.Print()
	
	
	
}
