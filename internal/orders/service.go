package orders

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderService interface {
	Save(orders []Order) error
}

type orderService struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewOrderService(client *mongo.Client, db, coll string) *orderService {
	database := client.Database(db)
	collection := database.Collection(coll)

	return &orderService{
		db:   database,
		coll: collection,
	}
}

func (o *orderService) Save(orders []Order) error {
	_, err := o.coll.InsertMany(context.Background(), []interface{}{orders})
	if err != nil {
		return err
	}

	return nil
}
