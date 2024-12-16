package warehouse

func (w *Warehouse) AddItem(name string, stock int) error {
	if _, exists := w.stocks[name]; exists {
		return ErrItemExists
	}
	w.stocks[name] = stock
	return nil
}

func (w *Warehouse) UpdateStock(name string, stock int) error {
	if _, exists := w.stocks[name]; !exists {
		return ErrItemNotFound
	}
	w.stocks[name] = stock
	return nil
}

func (w *Warehouse) CheckStock(name string) (int, error) {
	stock, exists := w.stocks[name]
	if !exists {
		return 0, ErrItemNotFound
	}
	return stock, nil
}

func (w *Warehouse) DeleteItem(name string) error {
	if _, exists := w.stocks[name]; !exists {
		return ErrItemNotFound
	}
	delete(w.stocks, name)
	return nil
}
