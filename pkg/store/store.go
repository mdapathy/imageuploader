package mongostore

import (
	"context"
	"errors"
	"github.com/mdapathy/imageuploader/pkg/domain"
	"github.com/mdapathy/imageuploader/pkg/domain/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewClient creates new connection to MongoDB and returns client.
func NewClient(ctx context.Context, uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}

// Store for interactions with database collections.
type Store struct {
	db *mongo.Database

	image *ImageRepository
}

// NewStore returns new database store object.
func NewStore(client *mongo.Client, db string) (*Store, error) {
	s := Store{
		db: client.Database(db),
	}

	return &s, nil
}

func (s *Store) Image() domain.ImageRepository {
	if s.image == nil {
		s.image = NewImageRepository(s.db)
	}

	return s.image
}

func handleMongoErr(err error) error {
	if errors.Is(err, mongo.ErrNoDocuments) {
		return model.ErrImageNotFound
	}

	return err
}
