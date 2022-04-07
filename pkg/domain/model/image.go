package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"strings"
	"time"
)

var validPrefixes = []string{
	"data:image/png;base64,",
	"data:image/jpeg;base64,",
}

var contentRegex = "^([A-Za-z0-9+\\/]{4})*([A-Za-z0-9+]{3}=|[A-Za-z0-9+\\/]{2}==)?$"

type Image struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    string             `json:"user_id" bson:"user_id"`
	Content   string             `json:"content" bson:"content"`
	Size      int                `json:"size" bson:"size"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

func checkContent(content string) (string, error) {
	for _, prefix := range validPrefixes {
		if strings.Contains(content, prefix) {
			baseContent := strings.TrimPrefix(content, prefix)
			if !regexp.MustCompile(contentRegex).MatchString(baseContent) {
				return "", ErrContentInvalid
			}
			return prefix, nil
		}
	}
	return "", ErrContentInvalid
}

func NewImage(content, userID string) (*Image, error) {
	extension, err := checkContent(content)
	if err != nil {
		return nil, err
	}
	return &Image{
		ID:        primitive.NewObjectID(),
		UserID:    userID,
		Content:   content,
		Size:      len(strings.TrimPrefix(content, extension)),
		CreatedAt: time.Now().UTC(),
	}, nil
}
