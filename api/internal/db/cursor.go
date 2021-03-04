package db

import (
	"encoding/base64"
	paginationV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
)

const (
	DefaultPageSize = 20
	MinPageSize     = 1
	MaxPageSize     = 100
)

func newCursor() string {

}

func parseCursor(in string) (Cursor, error) {
	// step 1: base-64 decode
	var decoded string
	if cursorBytes, err := base64.StdEncoding.DecodeString(in); err != nil {
		return Cursor{}, err
	} else {
		decoded = string(cursorBytes)
	}

	// step 2: parse id:
}

type Cursor struct {
	ascending bool
	field string // Dgraph field name, e.g., created_at
	value string
}

func getCursorAndPageSize(in *paginationV1.PaginationRequest) (string, int) {
	if in == nil {
		return "", DefaultPageSize
	}
	pageSize := DefaultPageSize
	var count int32
	var cursor string
	if f := in.GetForwardPaginationInfo(); f != nil {
		count = f.First
		cursor = f.After
	} else if b := in.GetBackwardPaginationInfo(); b != nil {
		count = b.Last
		cursor = b.Before
	}
	if count >= MinPageSize && count <= MaxPageSize {
		pageSize = int(count)
	}
	return cursor, pageSize
}

func emptyCursor(prevCursor string) *paginationV1.PageInfo {
	return &paginationV1.PageInfo{
		EndCursor:   "",
		HasNextPage: false,
	}
}
