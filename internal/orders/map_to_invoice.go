package orders

import (
	"fmt"

	"github.com/JOKR-Services/ifood_nfs_rerun/internal/graph"
)

// Esta constante é utilizada para adicionar um sufixo ao OrderId,
// dessa forma daremos um bypass no controle de duplicidade de invoices.
const appendAfterOrderId string = "X"

// Este produto será usado para adicionar um valor de 0.02 centavos ao pedido
const sampleProductSku string = "15000005677_sample"

func MapToInvoice(deliveryFeeOnly bool, order Order, mappedHubs map[string]string) graph.OrderData {
	orderData := graph.OrderData{
		OrderId:              order.OrderId + appendAfterOrderId,
		CustomerFederalTaxId: *order.Customer.CpfDocument,
		DeliveryFee:          float64(order.DeliveryFee),
		HubId:                mappedHubs[fmt.Sprintf("%d", order.StoreId)],
		IsInvoiceRerun:       false,
		LineItems:            mapInvoiceLineItems(order.LineItems),
		ShippingAddress:      mapShippingAddress(order.Customer.Name, order.DeliveryAddress),
	}

	// Remove 1 centavo do deliveryFee e adiciona um produto com valor de 0.02 centavos.
	if deliveryFeeOnly {
		orderData.DeliveryFee = float64(order.DeliveryFee - 0.01)
		orderData.LineItems = sampleLineItem()
	}

	return orderData
}

// Se adicionarmos um produto com o valor de 0.01 centavos, o valor do frete não é adicionado,
// então devemos ter um produto com um total de 0.02 centavos,
// e em seguida retirar 0.01 centavos do frete.
func sampleLineItem() []graph.LineItem {
	return []graph.LineItem{
		{
			TotalDiscount:     0,
			Sku:               sampleProductSku,
			Quantity:          1,
			Origin:            "CART",
			Name:              "Frete",
			LineNumber:        0,
			UnitPrice:         0.02,
			TotalUnitDiscount: 0,
			TotalPrice:        0.02,
		},
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
