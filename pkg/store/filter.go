package mongostore

import (
	"github.com/mdapathy/imageuploader/pkg/domain/model"
	"github.com/mdapathy/imageuploader/pkg/domain/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func NewSorter() query.SorterFactory {
	return &SorterFactory{}
}

func NewFilterer() query.Filterer {
	return &filters{}
}

type SorterFactory struct{}

func (f *SorterFactory) New() query.Sorter {
	return new(sorter)
}

type sorter struct {
	sorts bson.D
}

func (s *sorter) Append(column string, order query.SortingOrder) query.Sorter {
	if s.sorts == nil {
		s.sorts = make(bson.D, 0)
	}

	s.sorts = append(s.sorts, bson.E{
		Key:   column,
		Value: order,
	})

	return s
}

func (s *sorter) Build() interface{} {
	if s.sorts == nil {
		s.sorts = make(bson.D, 0)
	}

	return s.sorts
}

type filters struct{}

func (f filters) FilterByID(id string) (query.FindFilter, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, model.ErrIDInvalid
	}

	return func(filters map[string]interface{}) map[string]interface{} {
		filters["_id"] = objectID

		return filters
	}, nil
}

func (f filters) FilterByUserID(userID string) query.FindFilter {
	return func(filters map[string]interface{}) map[string]interface{} {
		filters["user_id"] = userID

		return filters
	}
}

func (f filters) FilterBySizeLessThan(size uint64) query.FindFilter {
	return func(filters map[string]interface{}) map[string]interface{} {
		filters["size"] = bson.M{"$lt": size}

		return filters
	}
}

func (f filters) FilterBySizeGreaterThanOrEqual(size uint64) query.FindFilter {
	return func(filters map[string]interface{}) map[string]interface{} {
		filters["size"] = bson.M{"$gte": int(size)}

		return filters
	}
}

func (f filters) FilterByCreatedAtEarlierThan(time time.Time) query.FindFilter {
	return func(filters map[string]interface{}) map[string]interface{} {
		filters["created_at"] = bson.M{"$lt": time}

		return filters
	}
}

func (f filters) FilterByCreatedAtLaterOrEqual(time time.Time) query.FindFilter {
	return func(filters map[string]interface{}) map[string]interface{} {
		filters["created_at"] = bson.M{"$gte": time}

		return filters
	}
}
