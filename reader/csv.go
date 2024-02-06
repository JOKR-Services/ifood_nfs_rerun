package reader

import (
	"encoding/csv"
	"os"
)

func NewCSVReader(name string) ([]Order, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csv := csv.NewReader(file)
	data, err := csv.ReadAll()
	if err != nil {
		return nil, err
	}

	orders := mapToOrder(data)

	return orders, nil
}
