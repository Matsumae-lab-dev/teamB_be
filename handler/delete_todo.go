package handler

import (
	"net/http"

	"github.com/Matsumae-lab-dev/teamB_be/db"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DeleteTodo(c echo.Context) error {

	id := c.Param("id")
	var todo db.Todo
	if err := db.DB.Where("id = ?", id).First(&todo).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			// return 404
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Todo Not Found",
			})

		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}

	}

	// delete todo, return 200
	db.DB.Delete(&db.Todo{}, id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Deletion Successful",
	})
}
