package reader

import (
	"encoding/csv"
	"os"
)

func (r *reader) ReadFromCSV() ([]Order, error) {
	file, err := os.Open(r.csvName)
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
