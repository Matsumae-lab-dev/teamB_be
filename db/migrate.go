package db

import (
	"log"
	"os"

	"github.com/Matsumae-lab-dev/teamB_be/util"
)

func Migrate() {

	DB.Exec("DROP TABLE IF EXISTS todos_users")
	DB.Exec("DROP TABLE IF EXISTS users")
	DB.Exec("DROP TABLE IF EXISTS todos")

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Todo{})

	// test user 作成
	hashedPass, _ := util.HashPassword(os.Getenv("TESTUSER_PASSWORD"))
	user := User{
		Username: "user",
		Email:    "user@teamb.com",
		Password: hashedPass,
	}
	DB.Create(&user)

	log.Print("[INFO] DB Migrated!")
}
