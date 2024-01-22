package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Matsumae-lab-dev/teamB_be/db"
	"github.com/Matsumae-lab-dev/teamB_be/handler"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	godotenv.Load(".env")
	e := echo.New()
	db.Connect()
	db.Migrate()

	config := middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
		ParseTokenFunc: func(tokenString string, c echo.Context) (interface{}, error) {
			keyFunc := func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("JWT_SECRET_KEY")), nil
			}

			token, err := jwt.Parse(tokenString, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	}
	e.Use(middleware.CORS())

	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)

	r := e.Group("/auth")
	r.Use(middleware.JWTWithConfig(config))

	r.GET("", handler.Auth)

	r.GET("/user/:id", handler.GetUser)

	r.POST("/todo", handler.CreateTodo)
	r.GET("/todo", handler.GetTodoList)
	r.PUT("/todo/:id", handler.UpdateTodo)
	r.DELETE("/todo/:id", handler.DeleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
