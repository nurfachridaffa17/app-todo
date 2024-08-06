package handler

import (
	"app-todo/internal/task/controllers"
	"app-todo/internal/task/repository"
	"app-todo/internal/task/service"

	categ "app-todo/internal/category/repository"
	taskCateg "app-todo/internal/task-category/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type handlerTask struct {
	Controller controllers.Controller
}

func NewHandlerTask(db *gorm.DB) *handlerTask {
	cr := repository.NewRepository()
	tc := taskCateg.NewRepository()
	cc := categ.NewRepository()
	cs := service.NewService(db, cr, tc, cc)

	return &handlerTask{
		Controller: controllers.NewController(cs),
	}
}

func (h *handlerTask) Route(g *echo.Group) {
	g.POST("/task", h.Controller.Create)
	g.PUT("/task/:id", h.Controller.Update)
}
