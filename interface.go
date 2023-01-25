package main

// interface to implement
type staticSequenceInterface interface {
	get_at(idx int) int
	set_at(item int, idx int)
	get_length() int
	iter_seq() []int
	build(items []int)
}

type dynamicSequenceInterface interface {
	staticSequenceInterface
	delete_at(idx int)
	add_at(item int, idx int)
}
