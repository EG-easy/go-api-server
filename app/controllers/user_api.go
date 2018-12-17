package controllers

import (
	"github.com/revel/revel"
	"go-api-server/app/models"
	"strconv"
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

func (c UserApi) GetUser() revel.Result {

	//int型に変換する
	number , _ := strconv.Atoi(c.Params.Get("number"))

	Users := []models.User{}

	DB.Find(&Users)

	response := JsonResponse{}
	response.Response = Users[number]

	return c.RenderJSON(response)
}