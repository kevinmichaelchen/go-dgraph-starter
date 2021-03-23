package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/MyOrg/todo-api/internal/db"
	"reflect"
	"strings"

	paginationV1 "github.com/MyOrg/todo-api/pkg/pb/myorg/todo/v1"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type xidRule struct{}

func newXidRule() xidRule {
	return xidRule{}
}

func (r xidRule) Validate(value interface{}) error {
	value, isNil := validation.Indirect(value)
	if isNil || validation.IsEmpty(value) {
		return nil
	}

	if s, ok := value.(string); ok {
		if s != "" {
			if len(s) != 20 {
				return errors.New("xid field should always be 20 chars")
			}
		}
	} else if !ok {
		return fmt.Errorf("field expected to be a string but was type %t", value)
	}

	return nil
}

type xidListRule struct{}

func newXidListRule() xidListRule {
	return xidListRule{}
}

func getInterface(value reflect.Value) interface{} {
	switch value.Kind() {
	case reflect.Ptr, reflect.Interface:
		if value.IsNil() {
			return nil
		}
		return value.Elem().Interface()
	default:
		return value.Interface()
	}
}

func (r xidListRule) Validate(value interface{}) error {

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			val := getInterface(v.Index(i))
			if s, ok := val.(string); ok {
				if len(s) != 20 {
					return errors.New("xid field should always be 20 chars")
				}
			} else {
				return errors.New("field expected to consist of strings")
			}
		}
	default:
		return errors.New("must be a slice or array")
	}

	return nil
}

type notAllWhiteSpaceRule struct{}

func newNotAllWhiteSpaceRule() notAllWhiteSpaceRule {
	return notAllWhiteSpaceRule{}
}

func (r notAllWhiteSpaceRule) Validate(value interface{}) error {
	if s, ok := value.(string); ok {
		if strings.TrimSpace(s) == "" {
			return errors.New("field cannot be entirely whitespace")
		}
	} else {
		return errors.New("field expected to be a string")
	}
	return nil
}

type SortRule struct {
	fields []string
}

func newSortRule(fields []string) SortRule {
	return SortRule{
		fields: fields,
	}
}

func (r SortRule) Validate(value interface{}) error {
	// it's fine if sort list is nil
	if value != nil {
		if sortList, ok := value.([]string); ok {
			for _, e := range sortList {
				arr := strings.Split(e, "|")
				if len(arr) != 2 {
					return errors.New("found malformed sort list element")
				}

				fieldName := arr[0]
				var fieldAllowed bool
				for _, f := range r.fields {
					if f == fieldName {
						fieldAllowed = true
						break
					}
				}
				if !fieldAllowed {
					return errors.New("field not allowed in sort list")
				}

				sortOrder := strings.ToLower(arr[1])
				if sortOrder != "asc" && sortOrder != "desc" {
					return errors.New("sort order can be either 'asc' or 'desc'")
				}
			}
		}
	}

	return nil
}

func isValidLang(lang string) bool {
	return lang == "en" || lang == "es"
}

func validateCursor(ctx context.Context, r *paginationV1.PaginationRequest) error {
	if f := r.GetForwardPaginationInfo(); f != nil {
		return validation.ValidateStruct(&f,
			validation.Field(&f.First, validation.Min(db.MinPageSize), validation.Max(db.MaxPageSize)),
		)
	} else if b := r.GetBackwardPaginationInfo(); b != nil {
		return validation.ValidateStruct(&b,
			validation.Field(&b.Last, validation.Min(db.MinPageSize), validation.Max(db.MaxPageSize)),
		)
	}
	return nil
}
