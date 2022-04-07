package mongostore

import (
	"context"
	"github.com/mdapathy/imageuploader/pkg/domain/model"
	"github.com/mdapathy/imageuploader/pkg/domain/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ImageRepository struct {
	collection *mongo.Collection
}

func NewImageRepository(db *mongo.Database) *ImageRepository {
	return &ImageRepository{
		collection: db.Collection("images"),
	}
}

func (i *ImageRepository) Create(ctx context.Context, image *model.Image) error {
	_, err := i.collection.InsertOne(ctx, image)
	if err != nil {
		return handleMongoErr(err)
	}

	return nil
}

func (i *ImageRepository) Delete(ctx context.Context, id primitive.ObjectID, userID string) error {
	filter := bson.M{
		"_id":     id,
		"user_id": userID,
	}
	res, err := i.collection.DeleteOne(ctx, filter)
	if err != nil {
		return handleMongoErr(err)
	}

	if res.DeletedCount == 0 {
		return model.ErrImageNotFound
	}
	return nil
}

func (i *ImageRepository) List(ctx context.Context, q *query.List) (*query.Result, error) {
	res := query.Result{}

	count, err := i.count(ctx, q.Filters...)
	if err != nil {
		return nil, err
	}
	res.Total = count

	if q.Sort == nil {
		q.Sort = &sorter{}
	}
	q.Sort.Append("_id", -1)

	opts := options.Find().SetSort(q.Sort.Build())
	opts.SetSkip(int64(q.Offset)).SetLimit(int64(q.Limit))

	findFilters := bson.M{}
	for _, f := range q.Filters {
		findFilters = f(findFilters)
	}

	cursor, err := i.collection.Find(ctx, findFilters, opts)
	if err != nil {
		return nil, err
	}

	objects := make([]model.Image, 0)
	if err := cursor.All(ctx, &objects); err != nil {
		return nil, err
	}

	res.Data = objects
	return &res, nil
}

func (i *ImageRepository) Details(ctx context.Context, q *query.Detail) (*model.Image, error) {
	findFilters := bson.M{}
	for _, f := range q.Filters {
		findFilters = f(findFilters)
	}

	findRes := i.collection.FindOne(ctx, findFilters, options.FindOne())
	if err := findRes.Err(); err != nil {
		return nil, handleMongoErr(err)
	}

	var object model.Image

	if err := findRes.Decode(&object); err != nil {
		return nil, err
	}

	return &object, nil
}

func (i *ImageRepository) count(ctx context.Context, filters ...query.FindFilter) (int64, error) {
	findFilters := make(bson.M)
	for _, f := range filters {
		findFilters = f(findFilters)
	}

	return i.collection.CountDocuments(ctx, findFilters)
}
