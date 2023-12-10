package handler

import (
	"net/http"

	"github.com/Matsumae-lab-dev/teamB_be/db"
	"github.com/Matsumae-lab-dev/teamB_be/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Signup(c echo.Context) error {
	type Body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	obj := new(Body)
	if err := c.Bind(obj); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}

	if util.HasEmptyField(obj, "Username", "Email", "Password") {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "missing request field",
		})
	}

	var user db.User
	if err := db.DB.Where("email = ?", obj.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			hashedPass, err := util.HashPassword(obj.Password)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Password Hashing Error",
				})
			}
			new := db.User{
				Username: obj.Username,
				Email:    obj.Email,
				Password: hashedPass,
			}
			db.DB.Create(&new)
			return c.JSON(http.StatusCreated, echo.Map{
				"id":         new.Id,
				"username":   new.Username,
				"email":      new.Email,
				"created_at": user.CreatedAt,
				"updated_at": user.UpdatedAt,
			})
		} else {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	} else {
		return c.JSON(http.StatusConflict, echo.Map{
			"message": "email conflict",
		})
	}
}
