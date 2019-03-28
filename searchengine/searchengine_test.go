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
