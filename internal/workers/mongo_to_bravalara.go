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

	for _, order := range dbOrders {
		if _, err := w.graphQlClient.GetInvoice(orders.MapToInvoice(order)); err != nil {
			return err
		}
		fmt.Println("Invoice request for order", order.OrderId, "sent")
	}

	return nil
}
