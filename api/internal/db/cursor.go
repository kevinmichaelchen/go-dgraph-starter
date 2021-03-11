package db

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/MyOrg/go-dgraph-starter/internal/obs"
	paginationV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"strings"
)

const (
	MinPageSize     = 1
	DefaultPageSize = 3
	MaxPageSize     = 50

	delimiter = ":::"
)

var (
	errInvalidCursor = errors.New("invalid cursor")
)

type Cursor struct {
	field string // Dgraph field name, e.g., "created_at"
	value string // cursor value, e.g., "2021-03-04T14:15:00"
}

func (c Cursor) encode() string {
	return base64.StdEncoding.EncodeToString([]byte(c.field + delimiter + c.value))
}

func newCursor(field, value string) Cursor {
	return Cursor{
		field: field,
		value: value,
	}
}

func parseCursor(ctx context.Context, in string) (Cursor, error) {
	logger := obs.ToLogger(ctx)

	if in == "" {
		return Cursor{
			// if client doesn't specify a cursor, we'll default to using creation time
			field: "created_at",
			value: "0001-01-01T00:00:00",
		}, nil
	}

	// step 1: base-64 decode
	var decoded string
	if cursorBytes, err := base64.StdEncoding.DecodeString(in); err != nil {
		return Cursor{}, err
	} else {
		decoded = string(cursorBytes)
	}

	r := strings.Split(decoded, delimiter)
	if len(r) != 2 {
		logger.Error().Msgf("Encoded: %s, Decoded: %s", in, decoded)
		return Cursor{}, errInvalidCursor
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
		StartCursor: "",
		EndCursor:   "",
		HasNextPage: false,
	}
}
