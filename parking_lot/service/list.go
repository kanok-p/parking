package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"parking/repository"
)

func (impl *implementation) List() (items []repository.Park, err error) {
	filter := bson.M{"deleted": bson.M{"$ne": "true"}}
	items, err = impl.repo.List(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return
}
