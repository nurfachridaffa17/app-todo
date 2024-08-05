package controllers

import (
	"app-todo/internal/category/service"
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

// @Summary FindAll
// @Description FindAll Category
// @Tags Category
// @Produce json
// @Param search query string false "search"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param filter query int false "filter"
// @Param name query string false "name"
// @Param user_id query string false "user_id"
// @Success 200 {object} dto.CategoryRes
// @Failure 500 {object} response.Error
// @Failure 400 {object} response.Error
// @Router /categ/categories [get]
func (co *controllerImpl) FindAll(c echo.Context) error {
	search := c.QueryParam("search")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	filter, _ := strconv.Atoi(c.QueryParam("filter"))
	name := c.QueryParam("name")
	userID := c.QueryParam("user_id")

	payload := dto.GetCategoryRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
		Filter: filter,
		Name:   name,
		UserID: userID,
	}

	data, err, total := co.Service.FindAll(c, payload)

	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err, err.Error())
	}

	return res.SuccessResponseWithTotal(data, total).Send(c)
}

// @Summary FindByID
// @Description FindByID
// @Tags Category
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dto.CategoryRes
// @Failure 500 {object} response.Error
// @Failure 400 {object} response.Error
// @Router /categ/categories/{id} [get]
func (co *controllerImpl) FindByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := co.Service.FindByID(c, id)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err, err.Error())
	}

	if data.ID == 0 {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err, "data in id "+c.Param("id")+" not found")
	}

	return res.SuccessResponse(data).Send(c)
}

// @Summary Create
// @Description Create
// @Tags Category
// @Produce json
// @Param request body dto.CreateCategory true "request body"
// @Success 200 {object} dto.CategoryRes
// @Failure 500 {object} response.Error
// @Failure 400 {object} response.Error
// @Router /categ/categories [post]
func (co *controllerImpl) Create(c echo.Context) error {
	var payload dto.CreateCategory
	if err := c.Bind(&payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err, "invalid request body")
	}

	data, err := co.Service.Create(c, payload)
	if err != nil {
		return res.ErrorResponse(err)
	}

	return res.SuccessResponse(data).Send(c)
}

// @Summary Update
// @Description Update
// @Tags Category
// @Produce json
// @Param id path int true "id"
// @Param request body dto.UpdateCategory true "request body"
// @Success 200 {object} dto.CategoryRes
// @Failure 500 {object} response.Error
// @Failure 400 {object} response.Error
// @Router /categ/categories/{id} [put]
func (co *controllerImpl) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var payload dto.UpdateCategory
	if err := c.Bind(&payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err, "invalid request body")
	}

	data, err := co.Service.Update(c, id, payload)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err, err.Error())
	}

	return res.SuccessResponse(data).Send(c)
}

// @Summary Delete
// @Description Delete
// @Tags Category
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dto.CategoryRes
// @Failure 500 {object} response.Error
// @Failure 400 {object} response.Error
// @Router /categ/categories/{id} [delete]
func (co *controllerImpl) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := co.Service.Delete(c, id)
	if err != nil {
		return res.ErrorResponse(err)
	}

	if data.ID == 0 {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err, "data in id "+c.Param("id")+" not found")
	}

	return res.SuccessResponse(data).Send(c)
}
