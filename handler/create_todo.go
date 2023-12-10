package handler

import (
	"net/http"

	"github.com/Matsumae-lab-dev/teamB_be/db"
	"github.com/Matsumae-lab-dev/teamB_be/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateTodo(c echo.Context) error {
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

	obj := new(Todo)
	if err := c.Bind(obj); err != nil {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}

	// check field
	if util.HasEmptyField(obj, "Title") {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Missing Required Field",
		})
	}

	// create todo, return 201
	new := db.Todo{
		Title:   obj.Title,
		Content: obj.Content,
		// Deadline:    obj.Deadline,
		Tag:         obj.Tag,
		TagColor:    obj.TagColor,
		CreaterId:   userid,
		RepeatFlag:  obj.RepeatFlag,
		RepeatSpan:  obj.RepeatSpan,
		RepeatCount: obj.RepeatCount,
	}
	db.DB.Create(&new)
	return c.JSON(http.StatusCreated, echo.Map{
		"id":      new.Id,
		"title":   new.Title,
		"content": new.Content,
		// "deadline":     new.Deadline,
		"tag":          new.Tag,
		"tag_color":    new.TagColor,
		"creater_id":   new.CreaterId,
		"repeat_flag":  new.RepeatFlag,
		"repeat_span":  new.RepeatSpan,
		"repeat_count": new.RepeatCount,
	})

}
