package workers

import (
	"context"
	"fmt"

	"github.com/JOKR-Services/ifood_nfs_rerun/internal/orders"
)

func (w *worker) MongoToBravalara() error {
	dbOrders, err := w.orderService.GetOrders(context.Background())
	if err != nil {
		return err
	}

	//
	// Gerar produto fake 1 cent + deliveryFee
	//

	hubs, err := w.hubService.GetHubs(context.Background())
	if err != nil {
		return err
	}

	for _, order := range dbOrders {
		fmt.Println(orders.MapToInvoice(order, hubs))

		// if _, err := w.graphQlClient.GetInvoice(orders.MapToInvoice(order)); err != nil {
		// 	return err
		// }

		fmt.Println("Invoice request for order", order.OrderId, "sent")
	}

	return nil
}
