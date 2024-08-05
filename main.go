package main

import (
	"app-todo/pkg/database"
	"app-todo/pkg/migration"
	"app-todo/pkg/util"
	"app-todo/pkg/util/env"
	"app-todo/routes"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	_ "app-todo/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger To Do App API
// @version 1.0
// @description This is an swagger for To Do App API
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @basePath /api/v1
func main() {
	e := echo.New()

	env := env.NewEnv()
	env.Load()

	database.Init()

	migration.Init()

	e.Validator = &util.CustomValidation{Validator: validator.New()}

	logrus.SetOutput(os.Stdout)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-Auth-Token"},
	}))

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	routes.Init(e.Group("/api/v1"))
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":" + env.GetString("PORT")))

}
