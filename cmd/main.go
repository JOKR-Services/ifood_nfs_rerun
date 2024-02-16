package main

import (
	"context"

	"github.com/JOKR-Services/ifood_nfs_rerun/env"
	"github.com/JOKR-Services/ifood_nfs_rerun/integration/ifood"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/db"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/graph"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/hub"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/orders"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/reader"
	"github.com/JOKR-Services/ifood_nfs_rerun/internal/workers"
	"github.com/JOKR-Services/logr-go"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	envs             *env.Environment
	mongoClient      *mongo.Client
	twoplMongoClient *mongo.Client
	worker           workers.Worker
)

func init() {
	envs = env.Get()
}

func main() {
	mongoClient, err := db.NewMongoClient(envs.Storage.MongoUri, 10).Connect()
	if err != nil {
		logr.LogPanic("error connecting to mongo", err, logr.KindInfra, nil)
	}

	twoplMongoClient, err := db.NewMongoClient(envs.TwoPlMongoUri, 10).Connect()
	if err != nil {
		logr.LogPanic("error connecting to 2pl mongo", err, logr.KindInfra, nil)
	}

	worker = workers.NewWorker(
		ifood.NewAdapter(envs.Ifood.URL, envs.Ifood.ClientID, envs.Ifood.ClientSecret),
		orders.NewOrderService(mongoClient, envs.Storage.DbName, "orders"),
		hub.NewHubService(twoplMongoClient, envs.TwoPlDbName, "ifood_stores"),
		graph.NewGraphQlClient(envs.GraphQl.URL, envs.GraphQl.APIKey),
		reader.NewReader("input/Ifood-orders-zero-deliveryfee.csv"),
	)

	conns := []*mongo.Client{mongoClient, twoplMongoClient}
	for _, conn := range conns {
		defer func(conn *mongo.Client) {
			if conn == nil {
				return
			}
			if err := conn.Disconnect(context.Background()); err != nil {
				logr.LogPanic("MongoClient", err, logr.KindInfra, nil)
			}
		}(conn)
	}

	err = worker.MongoToBravalara()
	if err != nil {
		logr.LogPanic("error processing ifood orders to mongo", err, logr.KindDomain, nil)
	}
}
