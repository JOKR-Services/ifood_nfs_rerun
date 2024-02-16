package workers

import (
	"github.com/JOKR-Services/ifood_nfs_rerun/integration/ifood"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/graph"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/hub"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/orders"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/reader"
)

type Worker interface {
	IfoodOrdersToMongo() error
	MongoToBravalara() error
}

type worker struct {
	ifoodAdapter  ifood.Adapter
	orderService  orders.OrderService
	hubService    hub.HubService
	graphQlClient graph.GraphQlClient
	reader        reader.Reader
}

func NewWorker(ifoodAdapter ifood.Adapter, orderService orders.OrderService, hubService hub.HubService, graphQlClient graph.GraphQlClient, reader reader.Reader) Worker {
	return &worker{
		ifoodAdapter:  ifoodAdapter,
		orderService:  orderService,
		hubService:    hubService,
		graphQlClient: graphQlClient,
		reader:        reader,
	}
}
