package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Repo interface {
	List(ctx context.Context, filter bson.M) (items []Park, err error)
	Create(ctx context.Context, ent interface{}) (err error)
	Read(ctx context.Context, filters bson.M, out *Park) (err error)
	Update(ctx context.Context, filters bson.M, ent *Park) (err error)
	UpdateMany(ctx context.Context, filters bson.M) (err error)
	//Delete(ctx context.Context, filters bson.M) (err error)
	//Count(ctx context.Context, filters []string) (total int, err error)

	//Push(ctx context.Context, param *domain.SetOpParam) (err error)
	//Pop(ctx context.Context, param *domain.SetOpParam) (err error)
	//IsFirst(ctx context.Context, param *domain.SetOpParam) (is bool, err error)
	//CountArray(ctx context.Context, param *domain.SetOpParam) (total int, err error)
	//ClearArray(ctx context.Context, param *domain.SetOpParam) (err error)
}
