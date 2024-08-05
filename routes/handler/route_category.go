package handler

import (
	"app-todo/internal/category/controllers"
	"app-todo/internal/category/repository"
	"app-todo/internal/category/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type handlerCategory struct {
	Controller controllers.Controller
}

func NewHandlerCategory(db *gorm.DB) *handlerCategory {
	cr := repository.NewRepository()
	cs := service.NewService(db, cr)

	return &handlerCategory{
		Controller: controllers.NewController(cs),
	}
}

func (h *handlerCategory) Route(g *echo.Group) {
	g.GET("/categories", h.Controller.FindAll)
	g.GET("/categories/:id", h.Controller.FindByID)
	g.POST("/categories", h.Controller.Create)
	g.PUT("/categories/:id", h.Controller.Update)
	g.DELETE("/categories/:id", h.Controller.Delete)
}
