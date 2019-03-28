package searchengine_test

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
	return List{}
}

// Add inserts a new element into the list
func (l List) Add(s string) {
}

// GetUniqueElementsCount returns the number of uniq elements
func (l List) GetUniqueElementsCount() int {
	ct := 0

	return ct
}

// GetTopElements returns the top n elements (by weight)
func (l List) GetTopElements(n int) []Item {
	return nil
}
