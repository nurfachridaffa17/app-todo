package controllers

import (
	"app-todo/internal/task/service"
	"app-todo/models/dto"
	"strconv"

	res "app-todo/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type controllerImpl struct {
	Service service.Service
}

func NewController(service service.Service) Controller {
	return &controllerImpl{
		Service: service,
	}
}

// @Summary Create
// @Description Create Task
// @Tags Task
// @Produce json
// @Param payload body dto.CreateTask true "payload"
// @Success 200 {object} dto.CreateTaskRes
// @Failure 500 {object} response.Error
// @Failure 400 {object} response.Error
// @Router /task/task [post]
func (co *controllerImpl) Create(c echo.Context) error {
	var payload dto.CreateTask

	if err := c.Bind(&payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	data, err := co.Service.Create(c, payload)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err, err.Error()).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

// @Summary Update
// @Description Update Task
// @Tags Task
// @Produce json
// @Param id path int true "id"
// @Param payload body dto.UpdateTask true "payload"
// @Success 200 {object} dto.CreateTaskRes
// @Failure 500 {object} response.Error
// @Failure 400 {object} response.Error
// @Router /task/task/{id} [put]
func (co *controllerImpl) Update(c echo.Context) error {
	var payload dto.UpdateTask

	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	data, err := co.Service.Update(c, id, payload)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err, err.Error()).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
