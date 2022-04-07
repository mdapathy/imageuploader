package domain

import (
	"context"
	command "github.com/mdapathy/imageuploader/pkg/domain/cmd"
	"github.com/mdapathy/imageuploader/pkg/domain/model"
	"github.com/mdapathy/imageuploader/pkg/domain/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ImageRepository interface {
	Create(ctx context.Context, image *model.Image) error
	Delete(ctx context.Context, id primitive.ObjectID, userID string) error

	List(ctx context.Context, query *query.List) (*query.Result, error)
	Details(ctx context.Context, query *query.Detail) (*model.Image, error)
}

type ImageService interface {
	Create(ctx context.Context, cmd *command.Create) error
	Delete(ctx context.Context, cmd *command.Delete) error

	List(ctx context.Context, query *query.List) (*query.Result, error)
	Details(ctx context.Context, query *query.Detail) (*model.Image, error)
}
