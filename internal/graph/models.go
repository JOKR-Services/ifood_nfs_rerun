package graph

type OrderData struct {
	OrderId              string          `json:"orderId"`
	CustomerFederalTaxId string          `json:"customerFederalTaxId"`
	DeliveryFee          float64         `json:"deliveryFee"`
	HubId                string          `json:"hubId"`
	IsInvoiceRerun       bool            `json:"isInvoiceRerun"`
	LineItems            []LineItem      `json:"lineItems"`
	ShippingAddress      ShippingAddress `json:"shippingAddress"`
}

type LineItem struct {
	TotalDiscount     float64 `json:"totalDiscount"`
	Sku               string  `json:"sku"`
	Quantity          int     `json:"quantity"`
	Origin            string  `json:"origin"`
	Name              string  `json:"name"`
	LineNumber        int     `json:"lineNumber"`
	UnitPrice         float64 `json:"unitPrice"`
	TotalUnitDiscount float64 `json:"totalUnitDiscount"`
	TotalPrice        float64 `json:"totalPrice"`
}

type ShippingAddress struct {
	ApartmentNumber string `json:"apartmentNumber"`
	BusinessName    string `json:"businessName"`
	CityName        string `json:"cityName"`
	Name            string `json:"name"`
	Neighborhood    string `json:"neighborhood"`
	Number          string `json:"number"`
	State           string `json:"state"`
	Street          string `json:"street"`
	Zip             string `json:"zip"`
}
