package main

import "fmt"

//type Node struct {
//	next  *Node
//	Value rune
//

//type Odnosvyaz struct {
//	begin  *Node
//	length int
//}

//func Create_sp() *Odnosvyaz {
//	return &Odnosvyaz{begin: nil, length: 0}
//}

//func (sp *Odnosvyaz) Print_sp() {
//	if sp.length == 0 {
//		fmt.Println("Список пуст.")
//	}
//	now_node := sp.begin
//	for now_node != nil {
//		fmt.Println(string(now_node.Value))
//		now_node = now_node.next
//	}
//}

func (sp *Odnosvyaz) Add(value rune) {
	new_node := &Node{next: nil, Value: value}
	if sp.length == 0 {
		sp.begin = new_node
		sp.length++
	} else {
		node := sp.begin
		for node.next != nil {
			node = node.next
		}
		node.next = new_node
		sp.length++
	}
}

func (sp *Odnosvyaz) get_el(ind int) rune {
	node := sp.begin
	for i := 0; i < ind; i++ {
		node = node.next
	}
	return node.Value
}

func (sp *Odnosvyaz) Swap(ind1 int, ind2 int) {
	memory_link_1 := sp.begin
	memory_link_2 := sp.begin

	if ind1 == ind2 {
		return
	}
	if ind1 > ind2 {
		for i := 0; i < ind2; i++ {
			memory_link_1 = memory_link_1.next
			memory_link_2 = memory_link_2.next
		}
		for i := ind2; i < ind1; i++ {
			memory_link_1 = memory_link_1.next
		}
	} else {
		for i := 0; i < ind1; i++ {
			memory_link_1 = memory_link_1.next
			memory_link_2 = memory_link_2.next
		}
		for i := ind1; i < ind2; i++ {
			memory_link_2 = memory_link_2.next
		}
	}
	bufer := memory_link_1.Value
	memory_link_1.Value = memory_link_2.Value
	memory_link_2.Value = bufer

}
func (list *Odnosvyaz) HeapSort() {
	n := list.length
	for i := n/2 - 1; i >= 0; i-- {
		heapify(list, n, i)
	}
}

func heapify(list *Odnosvyaz, n int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && list.get_el(left) > list.get_el(largest) {
		largest = left
	}
	if right < n && list.get_el(right) > list.get_el(largest) {
		largest = right
	}
	if largest != i {
		list.Swap(i, largest)
		heapify(list, largest, i+1)
	}
}

func (sp *Odnosvyaz) Count_el_3() bool {
	sp.HeapSort()
	counter := 1

	for i := 0; i < sp.length-1; i++ {
		if sp.get_el(i) == sp.get_el(i+1) {
			counter++
			if counter == 3 {
				return true
			}

		} else {
			counter = 1
		}
	}
	return false
}

func main() {
	P := Create_sp()
	P.Add('A')
	P.Add('B')
	P.Add('C')
	P.Add('D')
	P.Add('C')
	P.Add('C')
	P.Add('D')

	fmt.Println(P.Count_el_3())
}
