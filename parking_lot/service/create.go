package service

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"parking/repository"
)

func (impl *implementation) Create(length int) (err error) {
	_ = impl.repo.UpdateMany(context.Background(), bson.M{})

	for i := 1; i <= length; i++ {
		rand.Seed(time.Now().UnixNano())
		err := impl.repo.Create(context.Background(), repository.Park{
			ID:     strconv.Itoa(int(rand.Int63()) + i),
			Index:  i,
			Car:    "",
			Colour: "",
			Status: "available",
		})
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}
