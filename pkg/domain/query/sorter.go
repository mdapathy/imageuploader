package query

type SortingOrder int

const (
	Ascending  SortingOrder = 1
	Descending SortingOrder = -1
)

type Sorter interface {
	Append(attr string, order SortingOrder) Sorter
	Build() interface{}
}

type SorterFactory interface {
	New() Sorter
}
