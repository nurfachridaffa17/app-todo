package response

import "app-todo/models/base"

type Meta struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Info    *base.PaginationInfo `json:"info"`
}
