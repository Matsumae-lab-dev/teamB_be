package handler

import (
	"net/http"

	"github.com/Matsumae-lab-dev/teamB_be/db"
	"github.com/Matsumae-lab-dev/teamB_be/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func UpdateTodo(c echo.Context) error {
	type Todo struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		// Deadline    time.Time `json:"deadline"`
		Tag         string `json:"tag"`
		TagColor    string `json:"tag_color"`
		CreaterId   uint   `json:"creater_id"`
		RepeatFlag  bool   `json:"repeat_flag"`
		RepeatSpan  uint   `json:"repeat_span"`
		RepeatCount uint   `json:"repeat_count"`
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	useridFloat := claims["id"].(float64)
	userid := uint(useridFloat)

	id := c.Param("id")
	var todo db.Todo
	if err := db.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		// return 500
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Database Error: " + err.Error(),
		})

	} else {

		// parse json
		obj := new(Todo)
		if err := c.Bind(obj); err != nil {
			// return 400
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Json Format Error: " + err.Error(),
			})
		}

		// check field
		if util.HasEmptyField(obj, "Username", "Title", "Category") {
			// return 400
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Missing Required Field",
			})
		}

		// create todo, return 201
		todo.Title = obj.Title
		todo.Content = obj.Content
		// todo.Deadline = obj.Deadline
		todo.Tag = obj.Tag
		todo.TagColor = obj.TagColor
		todo.CreaterId = userid
		todo.RepeatFlag = obj.RepeatFlag
		todo.RepeatSpan = obj.RepeatSpan
		todo.RepeatCount = obj.RepeatCount
		db.DB.Save(&todo)
		return c.JSON(http.StatusCreated, echo.Map{
			"id":      todo.Id,
			"title":   todo.Title,
			"content": todo.Content,
			// "deadline":     todo.Deadline,
			"tag":          todo.Tag,
			"tag_color":    todo.TagColor,
			"creater_id":   todo.CreaterId,
			"repeat_flag":  todo.RepeatFlag,
			"repeat_span":  todo.RepeatSpan,
			"repeat_count": todo.RepeatCount,
		})

	}
}
