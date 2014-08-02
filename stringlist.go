package main

import "math/rand"

type StringSlice struct {
	data []string
}

func NewStringSlice(size int) *StringSlice {
	return &StringSlice{data: make([]string, size)}
}

func (slice *StringSlice) Append(item string) {
	slice.data = append(slice.data, item)
}

func (slice *StringSlice) Pop() string {
	size := slice.Len() - 1

	if (size == -1) {
		panic("Trying to pop from an empty StringSlice")
	}

	str := slice.data[size]
	slice.data = slice.data[:size]

	return str
}

func (slice *StringSlice) Len() int {
	return len(slice.data)
}

func (slice *StringSlice) Shuffle() {
	for i := range slice.data {
		j := rand.Intn(i + 1)
		slice.data[i], slice.data[j] = slice.data[j], slice.data[i]
	}
}

func (slice *StringSlice) Clone() *StringSlice {
	newSlice := NewStringSlice(slice.Len())
	copy(newSlice.data, slice.data)

	return newSlice
}

type StringList struct {
	masterList *StringSlice
	options *StringSlice
}

func NewStringList() *StringList {
	return &StringList{
		masterList: NewStringSlice(0),
		options: NewStringSlice(0),
	}
}

func (slist *StringList) AddString(str string) {
	slist.masterList.Append(str)
}

func (slist *StringList) RandomString() string {
	// Clone and shuffle the master list if we have no strings
	if (slist.options.Len() < 1) {
		slist.options = slist.masterList.Clone()
		slist.options.Shuffle()
	}

	return slist.options.Pop()
}
