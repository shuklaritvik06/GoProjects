package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Data struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Raw struct {
	Data []Data
}

func ReadData() ([]Data, error) {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		return nil, err
	}
	var payload Raw
	err = json.Unmarshal(data, &payload.Data)
	if err != nil {
		return nil, err
	}
	return payload.Data, nil
}

func CreatePost(c echo.Context) error {
	data := json.NewDecoder(c.Request().Body)
	var newpost Data
	data.Decode(&newpost)
	jsondata, err := ReadData()
	if err != nil {
		fmt.Println(err)
	}
	jsondata = append(jsondata, Data{
		UserId: newpost.UserId,
		Id:     newpost.Id,
		Title:  newpost.Title,
		Body:   newpost.Body,
	})
	file, _ := json.MarshalIndent(jsondata, "", " ")
	_ = ioutil.WriteFile("data.json", file, 0644)
	return c.JSON(http.StatusOK, jsondata)
}

func DeletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	jsondata, err := ReadData()
	if err != nil {
		fmt.Println(err)
	}
	for idx, v := range jsondata {
		if v.Id == id {
			jsondata = append(jsondata[:idx], jsondata[idx+1:]...)
			break
		}
	}
	file, _ := json.MarshalIndent(jsondata, "", " ")
	_ = ioutil.WriteFile("data.json", file, 0644)
	return c.JSON(http.StatusOK, jsondata)
}

func DeletePosts(c echo.Context) error {
	file, _ := json.MarshalIndent(Data{}, "", " ")
	_ = ioutil.WriteFile("data.json", file, 0644)
	return c.JSON(http.StatusOK, "Deleted All Posts")
}

func GetPost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	jsondata, err := ReadData()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range jsondata {
		if v.Id == id {
			return c.JSON(http.StatusOK, v)
		}
	}
	return nil
}

func GetPosts(c echo.Context) error {
	jsondata, err := ReadData()
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, jsondata)
}

func UpdatePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data := json.NewDecoder(c.Request().Body)
	var newpost Data
	data.Decode(&newpost)
	jsondata, err := ReadData()
	if err != nil {
		fmt.Println(err)
	}
	for idx, v := range jsondata {
		if v.Id == id {
			jsondata[idx].UserId = newpost.UserId
			jsondata[idx].Id = newpost.Id
			jsondata[idx].Title = newpost.Title
			jsondata[idx].Body = newpost.Body
		}
	}
	file, _ := json.MarshalIndent(jsondata, "", " ")
	_ = ioutil.WriteFile("data.json", file, 0644)
	return c.JSON(http.StatusOK, "Updated")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API IS UP!")
	})
	e.POST("/create", CreatePost)
	e.PATCH("/update/:id", UpdatePost)
	e.DELETE("/delete/:id", DeletePost)
	e.DELETE("/delete", DeletePosts)
	e.GET("/posts", GetPosts)
	e.GET("/posts/:id", GetPost)
	e.Logger.Fatal(e.Start(":5000"))
}
