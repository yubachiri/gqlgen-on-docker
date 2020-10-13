package servers

import (
	"errors"
	"fmt"
	"go-graphql-api-sample/graph/model"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Printf("%v\n", "読めてないウホ🦍")
	}

	user := os.Getenv("DB_USER_NAME")
	password := os.Getenv("DB_PASSWORD")
	connectMethod := os.Getenv("DB_CONNECT_METHOD")
	containerName := os.Getenv("DB_CONTAINER_NAME")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	err := errors.New("")

	dataSourceName := user + ":" + password + "@" + connectMethod + "(" + containerName + ":" + port + ")/" + name
	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// db.LogMode(true)

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE app_db")
	db.Exec("USE app_db")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&model.Order{}, &model.Item{})
}
