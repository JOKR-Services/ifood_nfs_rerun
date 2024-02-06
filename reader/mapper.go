package reader

import (
	"strconv"
	"time"
)

type Order struct {
	Date        time.Time
	OrderId     string
	DeliveryFee float64
}

func mapToOrder(data [][]string) []Order {
	const (
		dateIndex        = 0
		orderIdIndex     = 1
		deliveryFeeIndex = 2
	)

	var orders []Order
	for line, record := range data {
		if line == 0 {
			continue
		}

		var order Order
		for index, field := range record {
			switch index {
			case dateIndex:
				order.Date, _ = time.Parse(time.RFC3339, field)
			case orderIdIndex:
				order.OrderId = field
			case deliveryFeeIndex:
				order.DeliveryFee, _ = strconv.ParseFloat(field, 64)
			}
			orders = append(orders, order)
		}
	}

	return orders
}
