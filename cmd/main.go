package main

import (
	"context"
	"time"

	"github.com/JOKR-Services/ifood_nfs_rerun/env"
	"github.com/JOKR-Services/ifood_nfs_rerun/integration/ifood"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/db"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/orders"
	"github.com/JOKR-Services/logr-go"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	envs         *env.Environment
	mongoClient  *mongo.Client
	ifoodAdapter ifood.Adapter
	orderService orders.OrderService
)

func init() {
	envs = env.Get()

	mongoClient, err := db.NewMongoClient(envs.Storage.MongoUri, time.Duration(envs.Storage.MongoTimeout)).Connect()
	if err != nil {
		logr.LogPanic("error connecting to mongo", err, logr.KindInfra, nil)
	}

	ifoodAdapter = ifood.NewAdapter(envs.Ifood.URL, envs.Ifood.ClientID, envs.Ifood.ClientSecret)
	orderService = orders.NewOrderService(mongoClient, envs.Storage.DbName, "orders")
}

func main() {
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			logr.LogPanic("MongoClient", err, logr.KindInfra, nil)
		}
	}()
}
