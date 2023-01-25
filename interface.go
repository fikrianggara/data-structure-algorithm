package main

// interface to implement
type staticSequenceInterface interface {
	get_at(idx int) int
	set_at(item int, idx int)
	get_length() int
	iter_seq() []int
	build(items []int)
	get_first() int
	set_first(item int)
	get_last() int
	set_last(item int)
}

type dynamicSequenceInterface interface {
	staticSequenceInterface
	delete_at(idx int)
	insert_at(item int, idx int)
}
