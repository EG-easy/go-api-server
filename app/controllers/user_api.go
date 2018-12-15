package controllers

import (
	"github.com/revel/revel"
	"go-api-server/app/models"
)

type UserApi struct {
	*revel.Controller
}

type JsonResponse struct {
	Response interface{} `json:"response"`
}

func (c UserApi) GetUsers() revel.Result {

	Users := []models.User{}

	DB.Find(&Users)

	response := JsonResponse{}
	response.Response = Users

	return c.RenderJSON(response)
}