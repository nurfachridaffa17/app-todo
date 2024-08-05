package response

import (
	"app-todo/models/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type successConstant struct {
	OK Success
}

type successConstantWithTotal struct {
	OK SuccessWithTotal
}

type successConstantLogin struct {
	OK SuccessLogin
}

type successConstantNil struct {
	OK SuccessNil
}

type successResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
	Code int         `json:"code"`
}

type successResponseLogin struct {
	Meta  Meta        `json:"meta"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
	Token string      `json:"token"`
}

type successResponseNil struct {
	Meta Meta `json:"meta"`
	Code int  `json:"code"`
}

type successResponseWithTotal struct {
	Meta  Meta        `json:"meta"`
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
	Code  int         `json:"code"`
}

type Success struct {
	Response successResponse `json:"response"`
	Code     int             `json:"code"`
}

type SuccessLogin struct {
	Response successResponseLogin `json:"response"`
	Code     int                  `json:"code"`
}

type SuccessNil struct {
	Response successResponseNil `json:"response"`
	Code     int                `json:"code"`
}

type SuccessWithTotal struct {
	Response successResponseWithTotal `json:"response"`
	Code     int                      `json:"code"`
}

var SuccessConstant successConstant = successConstant{
	OK: Success{
		Response: successResponse{
			Meta: Meta{
				Success: true,
				Message: "success",
			},
			Data: nil,
			Code: http.StatusOK,
		},
		Code: http.StatusOK,
	},
}

var SuccessConstantWithTotal successConstantWithTotal = successConstantWithTotal{
	OK: SuccessWithTotal{
		Response: successResponseWithTotal{
			Meta: Meta{
				Success: true,
				Message: "success",
			},
			Data:  nil,
			Total: 0,
			Code:  http.StatusOK,
		},
		Code: http.StatusOK,
	},
}

var SuccessConstantLogin successConstantLogin = successConstantLogin{
	OK: SuccessLogin{
		Response: successResponseLogin{
			Meta: Meta{
				Success: true,
				Message: "success",
			},
			Data:  nil,
			Code:  http.StatusOK,
			Token: "",
		},
		Code: http.StatusOK,
	},
}

var SuccessConstantNil successConstantNil = successConstantNil{
	OK: SuccessNil{
		Response: successResponseNil{
			Meta: Meta{
				Success: true,
				Message: "success",
			},
			Code: http.StatusOK,
		},
		Code: http.StatusOK,
	},
}

func SuccessBuilder(res *Success, data interface{}) *Success {
	res.Response.Data = data
	return res
}

func SuccessBuilderWithTotal(res *SuccessWithTotal, data interface{}, total int64) *SuccessWithTotal {
	res.Response.Data = data
	res.Response.Total = total
	return res
}

func SuccessBuilderLogin(res *SuccessLogin, data interface{}, token string) *SuccessLogin {
	res.Response.Data = data
	res.Response.Token = token
	return res
}

func CustomSuccessBuilder(code int, data interface{}, message string, info *base.PaginationInfo) *Success {
	return &Success{
		Response: successResponse{
			Meta: Meta{
				Success: true,
				Message: message,
				Info:    info,
			},
			Data: data,
		},
		Code: code,
	}
}

func SuccessResponse(data interface{}) *Success {
	return SuccessBuilder(&SuccessConstant.OK, data)
}

func SuccessResponseWithTotal(data interface{}, total int64) *SuccessWithTotal {
	return SuccessBuilderWithTotal(&SuccessConstantWithTotal.OK, data, total)
}

func SuccessResponseLogin(data interface{}, token string) *SuccessLogin {
	return SuccessBuilderLogin(&SuccessConstantLogin.OK, data, token)
}

func SuccessResponseNil() SuccessNil {
	return SuccessConstantNil.OK
}

func (s *Success) Send(c echo.Context) error {
	return c.JSON(s.Code, s.Response)
}

func (s *SuccessWithTotal) Send(c echo.Context) error {
	return c.JSON(s.Code, s.Response)
}

func (s *SuccessLogin) Send(c echo.Context) error {
	return c.JSON(s.Code, s.Response)
}

func (s *SuccessNil) Send(c echo.Context) error {
	return c.JSON(http.StatusOK, s.Response)
}
