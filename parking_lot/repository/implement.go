package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repository) List(ctx context.Context, filter bson.M) (items []Park, err error) {
	var item Park
	cursor, err := repo.Coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	for cursor.Next(ctx) {
		err = cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (repo *Repository) Create(ctx context.Context, ent interface{}) (err error) {
	_, err = repo.Coll.InsertOne(ctx, ent)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) Read(ctx context.Context, filters bson.M, out *Park) (err error) {
	return repo.Coll.FindOne(ctx, filters, options.FindOne().SetSort(bson.M{"index": 1})).Decode(out)
}

func (repo *Repository) Update(ctx context.Context, filters bson.M, ret *Park) (err error) {
	_, err = repo.Coll.UpdateOne(ctx, filters, bson.M{"$set": ret})
	return err
}

func (repo *Repository) UpdateMany(ctx context.Context, filters bson.M) (err error) {
	_, err = repo.Coll.UpdateMany(ctx, filters, bson.M{"$set": bson.M{"deleted": "true"}})
	return err
}
