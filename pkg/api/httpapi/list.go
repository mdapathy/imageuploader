package httpapi

import (
	"net/http"
)

type QueryList struct {
	Limit   uint64
	UserID  string
	Offset  uint64
	Queries map[string]string
}

func (q *QueryList) GetLimit() uint64 {
	return q.Limit
}

func (q *QueryList) GetUserID() string {
	return q.UserID
}

func (q *QueryList) GetQueries() map[string]string {
	return q.Queries
}

func (q *QueryList) GetOffset() uint64 {
	return q.Offset
}

func NewQueryList(r *http.Request, defaultLimit uint64) *QueryList {
	list := make(map[string]string)
	for k, v := range r.URL.Query() {
		list[k] = v[0]
	}

	return &QueryList{
		Limit:   Limit(r, defaultLimit),
		Offset:  Offset(r),
		UserID:  UserIDFromContext(r.Context()),
		Queries: list,
	}
}
