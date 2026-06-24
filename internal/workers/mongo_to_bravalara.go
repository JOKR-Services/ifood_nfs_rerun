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

	hubs, err := w.hubService.GetHubs(context.Background())
	if err != nil {
		return err
	}

	order := dbOrders[0]

	//for _, order := range slices. {
	if _, err := w.graphQlClient.GetInvoice(orders.MapToInvoice(true, order, hubs)); err != nil {
		return err
	}

	fmt.Println("Invoice generation request for order", order.OrderId, "sent")
	//}

	return nil
}
