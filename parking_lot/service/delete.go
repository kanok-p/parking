package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"parking/repository"
)

func (impl *implementation) Delete(input int) (err error) {
	filter := bson.M{"$and":
	bson.A{
		bson.M{"deleted": bson.M{"$ne": "true"}},
		bson.M{"index": input},
	}}
	p := repository.Park{}
	if err := impl.repo.Read(context.Background(), filter, &p); err != nil {
		return err
	}

	err = impl.repo.Update(context.Background(), filter, &repository.Park{
		ID:     p.ID,
		Index:  p.Index,
		Car:    "",
		Colour: "",
		Status: "available",
	})
	if err != nil {
		return err
	}

	return nil
}
