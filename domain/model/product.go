package model

type ProductType int

const (
	Movie ProductType = iota
	Series
)

func (pt ProductType) String() string {
	return [...]string{"movie", "series"}[pt]
}

type Product struct {
	Name string
	Type ProductType
}
