package routes

import (
	"app-todo/pkg/constant"
	"app-todo/pkg/database"
	"app-todo/pkg/util/env"
	"app-todo/routes/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	db, err := database.GetConnection(constant.DB_NAME)
	if err != nil {
		panic("Failed init db, connection is undefined")
	}

	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to "+env.NewEnv().GetString("APP")+"! version "+env.NewEnv().GetString("VERSION")+" in mode "+env.NewEnv().GetString("ENV"))
	})

	handler.NewHandlerCategory(db).Route(g.Group("/categ"))
	handler.NewHandlerTask(db).Route(g.Group("/task"))
}
