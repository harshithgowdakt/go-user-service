package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	users = map[int]*user{}
	seq   = 1
)

func creatUsers(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.JSON(http.StatusOK, id)
}

func initHttpHandlers(e *echo.Echo) {
	//Routes
	e.POST("/api/v1/users", creatUsers)
	e.GET("/api/v1/users/:id", getUser)
	e.DELETE("/api/v1/users/:id", deleteUser)
}
