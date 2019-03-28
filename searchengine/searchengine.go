package searchengine

import "sort"

// Item is a single element
type Item struct {
	Name   string
	Weight int
}

// List structures the data
type List struct {
	Indexes []string
	Set     map[string]int
}

// NewList creates an new empty List
func NewList() List {
	l := List{}
	l.Set = make(map[string]int)
	return l
}

// Add inserts a new element into the list
func (l List) Add(s string) {
	l.Indexes = append(l.Indexes, s)
	v, ok := l.Set[s]
	if !ok {
		l.Set[s] = 1
	} else {
		l.Set[s] = v + 1
	}
}

// GetUniqueElementsCount returns the number of uniq elements
func (l List) GetUniqueElementsCount() int {
	ct := 0
	for _, v := range l.Set {
		if v == 1 {
			ct++
		}
	}
	return ct
}

// GetTopElements returns the top n elements (by weight)
func (l List) GetTopElements(n int) []Item {
	var temp []Item
	for k, v := range l.Set {
		temp = append(temp, Item{k, v})
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i].Weight > temp[j].Weight
	})
	return temp[0:n]
}
