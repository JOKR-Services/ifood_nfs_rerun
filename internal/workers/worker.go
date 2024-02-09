package workers

import (
	"github.com/JOKR-Services/ifood_nfs_rerun/integration/ifood"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/calculator"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/graph"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/orders"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/reader"
)

type Worker interface {
	IfoodOrdersToMongo() error
}

type worker struct {
	ifoodAdapter  ifood.Adapter
	orderService  orders.OrderService
	graphQlClient graph.GraphQlClient
	reader        reader.Reader
	calculator    calculator.Calculator
}

func NewWorker(ifoodAdapter ifood.Adapter, orderService orders.OrderService, graphQlClient graph.GraphQlClient, reader reader.Reader) Worker {
	return &worker{
		ifoodAdapter:  ifoodAdapter,
		orderService:  orderService,
		graphQlClient: graphQlClient,
		reader:        reader,
	}
}
