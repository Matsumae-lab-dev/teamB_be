package handler

import (
	"net/http"
	"sort"

	"github.com/Matsumae-lab-dev/teamB_be/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetTodoList(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["id"].(float64)
	var usertodo db.User
	if err := db.DB.Preload("Todos").First(&usertodo, userid).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			// return 404
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "User Not Found",
			})

		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}

	} else {
		sort.Slice(usertodo.Todos, func(i, j int) bool {
			return usertodo.Todos[i].Tag < usertodo.Todos[j].Tag
		})

		groupedTodos := make(map[string][]db.Todo)
		for _, todo := range usertodo.Todos {
			groupedTodos[todo.Tag] = append(groupedTodos[todo.Tag], todo)
		}

		var sortedTodos [][]db.Todo
		for _, todos := range groupedTodos {
			sortedTodos = append(sortedTodos, todos)
		}

		// return 200
		return c.JSON(http.StatusOK, echo.Map{
			"todos": sortedTodos,
		})
	}
}
