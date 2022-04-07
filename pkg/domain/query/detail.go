package query

type IDetail interface {
	GetID() string
	GetUserID() string
}

type DetailFilterer interface {
	FilterByID(id string) (FindFilter, error)
	FilterByUserID(userID string) FindFilter
}

type Detail struct {
	Filters []FindFilter
}

func NewDetailFrom(query IDetail, filterer DetailFilterer) (*Detail, error) {
	detail := Detail{}

	idFilter, err := filterer.FilterByID(query.GetID())
	if err != nil {
		return nil, err
	}
	userIDFilter := filterer.FilterByUserID(query.GetUserID())

	return detail.withFilter(idFilter).withFilter(userIDFilter), nil
}

func (d *Detail) withFilter(f FindFilter) *Detail {
	d.Filters = append(d.Filters, f)
	return d
}
