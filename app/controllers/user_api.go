package controllers

import (
	"github.com/revel/revel"
	"go-api-server/app/models"
	"strconv"
	"encoding/json"
	"log"
	"fmt"
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

func (c UserApi) PostUser(age int, name string, description string, info [][]string) revel.Result {

	 var TotalNewData = []models.SiblingInfo{}
	length := len(info)
	fmt.Println(length)

	for i := 0; i < length; i++ {

		NewData := models.SiblingInfo{
			Sibling:info[i][0],
			Name:info[i][1],
		}

		TotalNewData = append(TotalNewData, NewData)
	}

	//structからjson.RawMessage型に変換
	RawData, err := json.Marshal(TotalNewData)
	if err != nil {
		log.Panic(err)
	}

	//挿入するデータの生成
	User := models.User{
		Age: age,
		Name: name,
		Description: description,
		SiblingInfo: RawData,
	}

	//DBに追加
	DB.NewRecord(User)
	DB.Create(&User)

	//結果表示用
	response := JsonResponse{}
	response.Response = User

	return c.RenderJSON(response)
}