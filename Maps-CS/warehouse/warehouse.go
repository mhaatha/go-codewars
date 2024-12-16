package warehouse

import "errors"

type Warehouse struct {
	stocks map[string]int
}

var (
	ErrItemNotFound = errors.New("item not found")
	ErrItemExists   = errors.New("item already exists")
)

// NewWarehouse creates a new instance of Warehouse
func NewWarehouse() *Warehouse {
	return &Warehouse{
		stocks: make(map[string]int),
	}
}
