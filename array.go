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

func (arr *staticArray) get_first() int {
	return arr.items[0]
}
func (arr *staticArray) set_first(item int) {
	arr.items[0] = item
}
func (arr *staticArray) get_last() int {
	return arr.items[len(arr.items)-1]
}
func (arr *staticArray) set_last(item int) {
	arr.items[len(arr.items)-1] = item
}

type dynamicArray struct {
	items []int
}

func (arr *dynamicArray) insert_at(item int, idx int) {
	//shifting item from index i to end to the right
	temp1 := arr.items[:idx]
	temp2 := arr.items[idx+1:]
	temp1 = append(temp1, item)
	temp1 = append(temp1, temp2...)
	arr.items = temp1
}

func (arr *dynamicArray) delete_at(idx int) {
	//shifting item from index i to end to the left
	temp1 := arr.items[:idx]
	temp2 := arr.items[idx+1:]
	temp1 = append(temp1, temp2...)
	arr.items = temp1
}

func (arr *dynamicArray) insert_first(item int) {
	var tempSlice = []int{}
	temp1 := arr.items
	tempSlice = append(tempSlice, item)
	temp1 = append(tempSlice, temp1...)
	arr.items = temp1
}

func (arr *dynamicArray) delete_last() {
	arr.items = arr.items[:len(arr.items)]
}
func (arr *dynamicArray) build(items []int) {
	// O(n)
	arr.items = items
}

func (arr *dynamicArray) get_at(idx int) int {
	//constant time O(1)
	return arr.items[idx]
}

func (arr *dynamicArray) set_at(item int, idx int) {
	//constant time O(1)
	arr.items[idx] = item
}

func (arr *dynamicArray) get_length() int {
	// O(n)
	return len(arr.items)
}
func (arr *dynamicArray) iter_seq() []int {
	// O(n)
	return arr.items
}

func (arr *dynamicArray) get_first() int {
	return arr.items[0]
}
func (arr *dynamicArray) set_first(item int) {
	arr.items[0] = item
}
func (arr *dynamicArray) get_last() int {
	return arr.items[len(arr.items)-1]
}
func (arr *dynamicArray) set_last(item int) {
	arr.items[len(arr.items)-1] = item
}
