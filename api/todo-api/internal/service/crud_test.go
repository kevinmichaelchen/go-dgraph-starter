package service

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/MyOrg/todo-api/internal/db"
	todoV1 "github.com/MyOrg/todo-api/pkg/pb/myorg/todo/v1"
	"github.com/rs/zerolog/log"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestPagination(t *testing.T) {
	Convey("CREATE", t, func(c C) {

		ctx := context.TODO()

		Reset(func() {
			// Drop all data and schema
			log.Info().Msg("Dropping all Dgraph data...")
			if err := db.NukeDataButNotSchema(context.Background(), dgraphClient); err != nil {
				log.Fatal().Err(err).Msg("failed to nuke data")
			}
		})

		for i := 1; i <= 10; i++ {
			res, err := svc.CreateTodo(ctx, &todoV1.CreateTodoRequest{
				Title: fmt.Sprintf("Todo %d", i),
			})
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
		}

		Convey("GET PAGE", func() {
			res, err := svc.GetTodos(ctx, &todoV1.GetTodosRequest{
				PaginationRequest: &todoV1.PaginationRequest{
					Request: &todoV1.PaginationRequest_ForwardPaginationInfo{
						ForwardPaginationInfo: &todoV1.ForwardPaginationRequest{
							First: 10,
							After: "",
						},
					},
				},
			})
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.PageInfo, ShouldNotBeNil)
			So(res.PageInfo.HasNextPage, ShouldBeFalse)
		})

	})
}

func TestCrud(t *testing.T) {
	Convey("CREATE", t, func(c C) {

		ctx := context.TODO()

		Reset(func() {
			// Drop all data and schema
			log.Info().Msg("Dropping all Dgraph data...")
			if err := db.NukeDataButNotSchema(context.Background(), dgraphClient); err != nil {
				log.Fatal().Err(err).Msg("failed to nuke data")
			}
		})

		var id string

		res, err := svc.CreateTodo(ctx, &todoV1.CreateTodoRequest{
			Title: "Todo 1",
		})
		So(err, ShouldBeNil)
		So(res, ShouldNotBeNil)
		So(res.Todo, ShouldNotBeNil)
		So(res.Todo.Id, ShouldNotBeEmpty)
		So(res.Todo.Title, ShouldEqual, "Todo 1")
		So(res.Todo.Done, ShouldBeFalse)

		id = res.Todo.Id

		Convey("GET", func() {
			res, err := svc.GetTodo(ctx, &todoV1.GetTodoRequest{
				Id: id,
			})
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.Todo, ShouldNotBeNil)
			So(res.Todo.Id, ShouldEqual, id)
			So(res.Todo.CreatedAt, ShouldNotBeNil)
			So(res.Todo.Title, ShouldEqual, "Todo 1")
			So(res.Todo.Done, ShouldBeFalse)

			Convey("UPDATE", func() {
				res, err := svc.UpdateTodo(ctx, &todoV1.UpdateTodoRequest{
					Id: id,
					FieldMask: &fieldmaskpb.FieldMask{
						Paths: []string{"title"},
					},
					Title: "New Title",
				})
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
				So(res.Todo, ShouldNotBeNil)
				So(res.Todo.Id, ShouldEqual, id)
				So(res.Todo.CreatedAt, ShouldNotBeNil)
				So(res.Todo.Title, ShouldEqual, "New Title")
				So(res.Todo.Done, ShouldBeFalse)

				Convey("DELETE", func() {
					res, err := svc.DeleteTodo(ctx, &todoV1.DeleteTodoRequest{
						Id: id,
					})
					So(err, ShouldBeNil)
					So(res, ShouldNotBeNil)

					Convey("GET", func() {
						res, err := svc.GetTodo(ctx, &todoV1.GetTodoRequest{
							Id: id,
						})
						So(res, ShouldBeNil)
						So(err, ShouldNotBeNil)
						So(errors.Is(err, db.ErrNotFound), ShouldBeTrue)
					})

				})
			})
		})

	})

}
