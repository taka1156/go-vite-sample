package controller

import (
	dbconnect "app/database"
	"app/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func TodoCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := dbconnect.Connect()
		defer db.Close()

		var todo model.Todo
		err := c.Bind(&todo)
		if err != nil {
			return err
		}
		fmt.Printf("result = [%s]", todo.TaskName)

		err = db.Create(&todo).Error
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, todo)
	}
}

func TodoRead() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := dbconnect.Connect()
		defer db.Close()

		var todos []model.Todo

		err := db.Find(&todos).Error
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, todos)
	}
}

func TodoUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := dbconnect.Connect()
		defer db.Close()

		newPost := new(model.Todo)
		if err := c.Bind(newPost); err != nil {
			return err
		}

		todoId := c.Param("id")
		if todoId != "" {
			post := model.Todo{}
			db.First(&post, "id = ?", todoId).Update(newPost)
			return c.JSON(http.StatusOK, post)
		} else {
			return c.JSON(http.StatusNotFound, nil)
		}

	}
}

func TodoDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := dbconnect.Connect()
		defer db.Close()

		newDelete := new(model.Todo)
		todoId := c.Param("id")
		db.Delete(&newDelete, "id = ?", todoId)

		return c.JSON(http.StatusOK, newDelete)
	}
}
