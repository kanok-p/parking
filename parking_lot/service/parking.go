package service

import (
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"

	"parking/repository"
)

func (impl *implementation) Parking(input *repository.Park) (ID string, err error) {
	filter := bson.M{"$and":
	bson.A{
		bson.M{"deleted": bson.M{"$ne": "true"}},
		bson.M{"status": "available"},
	}}
	// update
	p := repository.Park{}
	if err := impl.repo.Read(context.Background(), filter, &p); err != nil {
		return "", err
	}

	err = impl.repo.Update(context.Background(), filter, &repository.Park{
		ID:     p.ID,
		Index:  p.Index,
		Car:    input.Car,
		Colour: input.Colour,
		Status: "unavailable",
	})

	if err != nil {
		return "", err
	}
	ID = strconv.Itoa(p.Index)
	return ID, nil
}
