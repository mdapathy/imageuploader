package command

import (
	"github.com/mdapathy/imageuploader/pkg/domain/model"
)

type Create struct {
	Content string
	UserID  string
}

func (c *Create) Prepare() {}

func (c *Create) Validate() error {
	if c.UserID == "" {
		return model.ErrUserIDRequired
	}

	if c.Content == "" {
		return model.ErrContentRequired
	}
	return nil
}
