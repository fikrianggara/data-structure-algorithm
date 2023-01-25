package main

import "fmt"

func staticSequenceArray() {
	var static1 staticSequenceInterface
	//implementation of static sequence using array
	static1 = &staticArray{}
	static1.build([]int{1, 2, 3, 4, 5})
	fmt.Println("items :", static1.iter_seq())
	fmt.Println("element at 0-index: ", static1.get_at(0))
	fmt.Println("length of sequence :", static1.get_length())
	static1.set_at(3, 0)
	fmt.Print("update element at 0 into ", static1.get_at(0))
	fmt.Println("items :", static1.iter_seq())
}

func dynamicSequenceLinkedList() {
	ll1 := singlyLinkedList{}
	ll1.insert(3)
	ll1.insert(4)
	ll1.insert(54)
	ll1.insert(9)

	ll1.traverse()

	ll1.insert_at(2, 10)

	ll1.traverse()

}
func main() {
	// staticSequenceArray()
	dynamicSequenceLinkedList()

}
