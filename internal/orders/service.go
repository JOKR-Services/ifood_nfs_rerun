package orders

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderService interface {
	Save(ctx context.Context, order Order) error
	GetOrders(ctx context.Context) ([]Order, error)
}

type orderService struct {
	coll *mongo.Collection
}

func NewOrderService(client *mongo.Client, db, coll string) *orderService {
	database := client.Database(db)
	collection := database.Collection(coll)

	return &orderService{
		coll: collection,
	}
}

func (o *orderService) Save(ctx context.Context, order Order) error {
	_, err := o.coll.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (o *orderService) GetOrders(ctx context.Context) ([]Order, error) {
	var orders []Order
	cursor, err := o.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
