package reader

type Order struct {
	OrderId         string
	ExternalOrderId string
}

func mapToOrder(data [][]string) []Order {
	const (
		orderIdIndex         = 1
		externalOrderIdIndex = 2
	)

	var orders []Order
	for line, record := range data {
		if line == 0 {
			continue
		}

		var order Order
		for i, field := range record {
			switch i {
			case orderIdIndex:
				order.OrderId = field
			case externalOrderIdIndex:
				order.ExternalOrderId = field
			}
		}
		orders = append(orders, order)
	}

	return orders
}
