package dto

import (
	model "app-todo/models/entity"
	res "app-todo/pkg/util/response"
)

type (
	GetTaskRequest struct {
		Page    int    `queryparam:"page"`
		Limit   int    `queryparam:"limit"`
		Search  string `queryparam:"search"`
		Filter  int    `queryparam:"filter"`
		Title   string `queryparam:"title"`
		DueDate string `queryparam:"due_date"`
	}

	CreateTask struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
		Status      string `json:"status"`
	}

	UpdateTask struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
		Status      string `json:"status"`
	}

	DeleteTask struct {
		ID int `json:"id"`
	}
)

type (
	TaskRes struct {
		model.MasterTaskModel
	}

	TaskResponse struct {
		Body struct {
			Meta res.Meta `json:"meta"`
			Data TaskRes  `json:"data"`
		} `json:"body"`
	}
)
