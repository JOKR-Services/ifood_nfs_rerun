package hub

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type HubService interface {
	GetHubs(ctx context.Context) (map[string]string, error)
}

type hubService struct {
	coll *mongo.Collection
}

func NewHubService(client *mongo.Client, db, coll string) *hubService {
	collection := client.Database(db).Collection(coll)

	return &hubService{
		coll: collection,
	}
}

func (h *hubService) GetHubs(ctx context.Context) (map[string]string, error) {
	var hubs []Hub
	cursor, err := h.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &hubs)
	if err != nil {
		return nil, err
	}

	return h.toMap(hubs), nil
}

func (h *hubService) toMap(hubs []Hub) map[string]string {
	m := make(map[string]string)
	for _, hub := range hubs {
		m[hub.From] = hub.To
	}
	return m
}
