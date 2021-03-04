package db

import (
	"encoding/base64"
	"errors"
	paginationV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"strings"
)

const (
	MinPageSize     = 1
	DefaultPageSize = 3
	MaxPageSize     = 50
)

type Cursor struct {
	field string // Dgraph field name, e.g., "created_at"
	value string // cursor value, e.g., "2021-03-04 14:15:00"
}

func (c Cursor) encode() string {
	return base64.StdEncoding.EncodeToString([]byte(c.field + ":" + c.value))
}

func newCursor(field, value string) Cursor {
	return Cursor{
		field: field,
		value: value,
	}
}

func parseCursor(in string) (Cursor, error) {
	// step 1: base-64 decode
	var decoded string
	if cursorBytes, err := base64.StdEncoding.DecodeString(in); err != nil {
		return Cursor{}, err
	} else {
		decoded = string(cursorBytes)
	}

	r := strings.Split(decoded, ":")
	if len(r) != 2 {
		return Cursor{}, errors.New("base64-decoded cursor should include 2 elements with a colon delimiter")
	}

	return Cursor{
		field: r[0],
		value: r[1],
	}, nil
}

func getPaginationInfo(in *paginationV1.PaginationRequest) (string, int, bool) {
	if in == nil {
		return "", DefaultPageSize, true
	}
	pageSize := DefaultPageSize
	var count int32
	var cursor string
	var forwardsPagination bool
	if f := in.GetForwardPaginationInfo(); f != nil {
		count = f.First
		cursor = f.After
		forwardsPagination = true

		if cursor == "" {
			cursor = ""
		}
	} else if b := in.GetBackwardPaginationInfo(); b != nil {
		count = b.Last
		cursor = b.Before
		forwardsPagination = false
	}
	if count >= MinPageSize && count <= MaxPageSize {
		pageSize = int(count)
	}
	return cursor, pageSize, forwardsPagination
}

func emptyPageInfo() *paginationV1.PageInfo {
	return &paginationV1.PageInfo{
		EndCursor:   "",
		HasNextPage: false,
	}
}
