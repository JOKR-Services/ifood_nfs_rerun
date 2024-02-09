package orders

import "github.com/JOKR-Services/ifood_nfs_rerun/integration/ifood"

func MapFromIfood(ifoodOrder *ifood.OrderDetails) Order {
	return Order{
		OrderId:            ifoodOrder.OrderId,
		CreatedAt:          ifoodOrder.CreatedAt,
		StoreId:            ifoodOrder.StoreId,
		CustomerId:         ifoodOrder.CustomerId,
		Customer:           mapCustomer(ifoodOrder.Customer),
		OrderCode:          ifoodOrder.OrderCode,
		DeliveryAddress:    mapAddress(ifoodOrder.DeliveryAddress),
		DeliveryFee:        ifoodOrder.DeliveryFee,
		TotalPrice:         ifoodOrder.TotalPrice,
		DiscountValue:      ifoodOrder.DiscountValue,
		DocumentForTax:     ifoodOrder.DocumentForTax,
		LineItems:          mapLineItems(ifoodOrder.LineItems),
		ScheduledBeginDate: ifoodOrder.ScheduledBeginDate,
		ScheduledEndDate:   ifoodOrder.ScheduledEndDate,
		ScheduledBeginTime: ifoodOrder.ScheduledBeginTime,
		ScheduledEndTime:   ifoodOrder.ScheduledEndTime,
		Benefits:           mapBenefits(ifoodOrder.Benefits),
	}
}

func mapCustomer(ifoodCustomer ifood.Customer) Customer {
	return Customer{
		Id:           ifoodCustomer.Id,
		Name:         ifoodCustomer.Name,
		Email:        ifoodCustomer.Email,
		CpfDocument:  ifoodCustomer.CpfDocument,
		CnpjDocument: ifoodCustomer.CnpjDocument,
		Type:         ifoodCustomer.Type,
		Phone:        ifoodCustomer.Phone,
	}
}

func mapAddress(ifoodAddress ifood.Address) Address {
	return Address{
		Street:        ifoodAddress.Street,
		StreetNumber:  ifoodAddress.StreetNumber,
		Apartment:     ifoodAddress.Apartment,
		Neighbourhood: ifoodAddress.Neighbourhood,
		City:          ifoodAddress.City,
		Province:      ifoodAddress.Province,
		ProvinceCode:  ifoodAddress.ProvinceCode,
		Zip:           ifoodAddress.Zip,
		Latitude:      ifoodAddress.Latitude,
		Longitude:     ifoodAddress.Longitude,
	}
}

func mapLineItems(ifoodLineItems []ifood.LineItems) []LineItem {
	var lineItems []LineItem
	for _, item := range ifoodLineItems {
		lineItems = append(lineItems, LineItem{
			Id:           item.Id,
			Name:         item.Name,
			BarCode:      item.BarCode,
			InternalCode: item.InternalCode,
			UnitPrice:    item.UnitPrice,
			TotalPrice:   item.TotalPrice,
			Quantity:     item.Quantity,
		})
	}
	return lineItems
}

func mapBenefits(ifoodBenefits []ifood.Benefits) []Benefits {
	var benefits []Benefits
	for _, benefit := range ifoodBenefits {
		benefits = append(benefits, Benefits{
			Sponsor: benefit.Sponsor,
			Value:   benefit.Value,
			Type:    benefit.Type,
		})
	}
	return benefits
}
