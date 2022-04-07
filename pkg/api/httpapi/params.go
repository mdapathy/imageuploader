package httpapi

import (
	"net/http"
	"strconv"
)

func Limit(r *http.Request, defaultLimit uint64) uint64 {
	limit, err := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 64)
	if err != nil {
		return defaultLimit
	}

	if limit > defaultLimit || limit == 0 {
		return defaultLimit
	}

	return limit
}

// Offset returns offset from the the request.
func Offset(r *http.Request) uint64 {
	offset, err := strconv.ParseUint(r.URL.Query().Get("offset"), 10, 64)
	if err != nil {
		return 0
	}
	return offset
}
