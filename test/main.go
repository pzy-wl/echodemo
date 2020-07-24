package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User1 struct {
	Name string `query:"name" form:"name" json:"name"`
	Sex  string `query:"sex" form:"sex" json:"sex"`
}

func main() {
	e := echo.New()

	e.Any("/", func(ctx echo.Context) error {
		user := new(User1)
		if err := ctx.Bind(user); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, user)
	})

	e.Logger.Fatal(e.Start(":2021"))
}
