package workers

import (
	"context"

	"github.com/JOKR-Services/ifood_nfs_rerun/internal/orders"
	"golang.org/x/time/rate"
)

type Order struct {
	OrderId         string
	ExternalOrderId string
}

func (w *worker) IfoodOrdersToMongo() error {
	csvOrders, err := w.reader.ReadFromCSV()

	if err != nil {
		return err
	}

	rate := rate.NewLimiter(10, 1)
	for _, order := range csvOrders {
		if err := rate.Wait(context.Background()); err != nil {
			return err
		}

		orderDetails, err := w.ifoodAdapter.GetOrderDetails(order.ExternalOrderId)
		if err != nil {
			return err
		}

		if err := w.orderService.Save(context.Background(), orders.MapFromIfood(orderDetails)); err != nil {
			return err
		}
	}

	return nil
}
