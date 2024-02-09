package orders

import "github.com/JOKR-Services/ifood_nfs_rerun/internal/graph"

func MapToInvoice(order Order) graph.OrderData {
	return graph.OrderData{
		OrderId:              order.OrderId,
		CustomerFederalTaxId: *order.Customer.CpfDocument,
		DeliveryFee:          float64(order.DeliveryFee),
		HubId:                "",
		IsInvoiceRerun:       false,
		LineItems:            mapInvoiceLineItems(order.LineItems),
		ShippingAddress:      mapShippingAddress(order.Customer.Name, order.DeliveryAddress),
	}
}

func mapInvoiceLineItems(lineItems []LineItem) []graph.LineItem {
	var items []graph.LineItem
	for lineNumber, item := range lineItems {
		items = append(items, graph.LineItem{
			TotalDiscount:     0, // Definir
			Sku:               *item.InternalCode,
			Quantity:          int(*item.Quantity),
			Origin:            "CART",
			Name:              *item.Name,
			LineNumber:        lineNumber,
			UnitPrice:         float64(*item.UnitPrice),
			TotalUnitDiscount: 0, // Definir
			TotalPrice:        float64(*item.TotalPrice),
		})
	}
	return items
}

func mapShippingAddress(name string, address Address) graph.ShippingAddress {
	return graph.ShippingAddress{
		Name:            name,
		BusinessName:    name,
		Street:          address.Street,
		Number:          address.StreetNumber,
		ApartmentNumber: address.Apartment,
		Neighborhood:    address.Neighbourhood,
		CityName:        address.City,
		State:           address.Province,
		Zip:             address.Zip,
	}
}
