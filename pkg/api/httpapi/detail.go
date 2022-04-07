package httpapi

type QueryDetail struct {
	ID     string
	UserID string
}

func (q *QueryDetail) GetID() string {
	return q.ID
}

func (q *QueryDetail) GetUserID() string {
	return q.UserID
}
