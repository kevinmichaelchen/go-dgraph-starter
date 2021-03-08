package service

import (
	"context"
	"errors"
	"testing"

	"github.com/MyOrg/go-dgraph-starter/internal/db"
	todoV1 "github.com/MyOrg/go-dgraph-starter/pkg/pb/myorg/todo/v1"
	"github.com/rs/zerolog/log"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCrud(t *testing.T) {
	Convey("CREATE", t, func(c C) {

		ctx := context.TODO()

		Reset(func() {
			// Drop all data and schema
			log.Info().Msg("Dropping all Dgraph data...")
			if err := db.NukeData(context.Background(), dgraphClient); err != nil {
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
			So(res.Todo.Title, ShouldEqual, "Todo 1")
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

}
