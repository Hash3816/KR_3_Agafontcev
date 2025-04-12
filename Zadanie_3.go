package main

import (
	"fmt"
)

//type Node struct {
//	next  *Node
//	Value rune
//}

//type Odnosvyaz struct {
//	begin  *Node
//	length int
//}

//func Create_sp() *Odnosvyaz {
//	return &Odnosvyaz{begin: nil, length: 0}
//}
func (sp *Odnosvyaz) Add(value rune) {
	new_node := &Node{next: nil, Value: value}
	if sp.length == 0 {
		sp.begin = new_node
	}
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

func (sp *Odnosvyaz) Count_el_3() bool {
	el_nums_sp := Create_sp()
	now_node := sp.begin

	for i := 0; i < sp.length; i++ {
		value_1 := now_node.Value
		now_node = now_node.next
		el_node := el_nums_sp.begin

		counter := 0
		for j := 0; j < el_nums_sp.length; j++ {
			value_2 := el_node.Value
			el_node = el_node.next
			if value_1 == value_2 {
				counter++
			}
		}
		if counter == 0 {
			el_nums_sp.Add(value_1)
		}
	}

	el_node := el_nums_sp.begin

	for i := 0; i < el_nums_sp.length; i++ {
		counter := 0
		value_el_node := el_node.Value
		el_node = el_node.next
		sp1_node := sp.begin

		for j := 0; j < sp.length; j++ {
			value_old_sp := sp1_node.Value
			sp1_node = sp1_node.next
			if value_el_node == value_old_sp {
				counter++
			}
		}
		if counter == 3 {
			return true
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
