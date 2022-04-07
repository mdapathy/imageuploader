package command

import (
	"github.com/mdapathy/imageuploader/pkg/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Delete struct {
	ID     primitive.ObjectID
	UserID string
}

func (c *Delete) Prepare() {}

func (c *Delete) Validate() error {
	if c.ID.IsZero() {
		return model.ErrIDRequired
	}

	if c.UserID == "" {
		return model.ErrUserIDRequired
	}

	return nil
}
