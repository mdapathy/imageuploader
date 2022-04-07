package query

import (
	"github.com/mdapathy/imageuploader/pkg/domain/model"
	"strconv"
	"strings"
	"time"
)

type FindFilter func(filters map[string]interface{}) map[string]interface{}

type ListFilterer interface {
	FilterByUserID(userID string) FindFilter

	FilterBySizeLessThan(size uint64) FindFilter
	FilterBySizeGreaterThanOrEqual(size uint64) FindFilter

	FilterByCreatedAtEarlierThan(time time.Time) FindFilter
	FilterByCreatedAtLaterOrEqual(time time.Time) FindFilter
}

type Result struct {
	Total int64 `json:"total"`
	Data  []model.Image `json:"data"`
}

// IList is a contract for query object (e.g. GRPC structure).
type IList interface {
	GetOffset() uint64
	GetLimit() uint64
	GetUserID() string
	GetQueries() map[string]string
}

type List struct {
	Offset  uint64
	Limit   uint64
	Sort    Sorter
	Filters []FindFilter
}

func NewListFrom(query IList, sf SorterFactory, filterer ListFilterer) (*List, error) {
	list := List{
		Limit:  query.GetLimit(),
		Offset: query.GetOffset(),
	}

	queryStrings := query.GetQueries()
	if queryStrings == nil {
		return &list, nil
	}

	userID := query.GetUserID()
	if userID == "" {
		return nil, model.ErrUserIDRequired
	}
	list.WithFilter(filterer.FilterByUserID(userID))

	for q, v := range queryStrings {
		if err := parseTimeParams(q, v, filterer, &list); err != nil {
			return nil, err
		}

		switch q {
		case "size_from":
			size, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, model.ErrSizeFromInvalid
			}
			list.WithFilter(filterer.FilterBySizeGreaterThanOrEqual(size))
		case "size_to":
			size, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil,  model.ErrSizeToInvalid
			}
			list.WithFilter(filterer.FilterBySizeLessThan(size))
		default:
			continue
		}
	}

	sorter := sf.New()
	if sorts, present := queryStrings["sort"]; present {
		cols := strings.Split(sorts, ",")

		for _, c := range cols {
			order := Ascending

			if strings.HasPrefix(c, "-") {
				order = Descending
				c = strings.TrimPrefix(c, "-")
			}

			sorter.Append(c, order)
		}
	}

	list.WithSorting(sorter)

	return &list, nil
}

func (l *List) WithSorting(s Sorter) *List {
	l.Sort = s
	return l
}

func (l *List) WithFilter(f FindFilter) *List {
	l.Filters = append(l.Filters, f)
	return l
}

func parseTimeParams(q, v string, filterer ListFilterer, list *List) error {
	if q == "created_from" {
		t, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return model.ErrCreatedFromInvalid
		}
		list.WithFilter(filterer.FilterByCreatedAtLaterOrEqual(t))
		return nil
	}

	if q == "created_to" {
		t, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return model.ErrCreatedToInvalid
		}
		list.WithFilter(filterer.FilterByCreatedAtEarlierThan(t))
		return nil
	}

	return nil
}
