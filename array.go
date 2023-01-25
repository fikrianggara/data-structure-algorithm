package main

type staticArray struct {
	items []int
}

func (arr *staticArray) build(items []int) {
	// O(n)
	arr.items = items
}

func (arr *staticArray) get_at(idx int) int {
	//constant time O(1)
	return arr.items[idx]
}

func (arr *staticArray) set_at(item int, idx int) {
	//constant time O(1)
	arr.items[idx] = item
}

func (arr *staticArray) get_length() int {
	// O(n)
	return len(arr.items)
}
func (arr *staticArray) iter_seq() []int {
	// O(n)
	return arr.items
}
