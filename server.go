package main

import (
	"go-graphql-api-sample/graph"
	"go-graphql-api-sample/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"xorm.io/xorm"

	_ "github.com/go-sql-driver/mysql"

	"errors"
	"fmt"
	"go-graphql-api-sample/graph/model"

	"github.com/joho/godotenv"
)

const defaultPort = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

var DB *xorm.Engine

func initDB() {
	envErr := godotenv.Load()
	if envErr != nil {
		// .env読めなかった場合の処理
		//log.Fatal("Error loading .env file")
		fmt.Printf("%v\n", "読めてないウホ🦍")
	}
	env := os.Getenv("ENV")
	fmt.Println(env)

	driverName := "mysql"

	user := os.Getenv("DB_USER_NAME")
	password := os.Getenv("DB_PASSWORD")
	connectMethod := os.Getenv("DB_CONNECT_METHOD")
	containerName := os.Getenv("DB_CONTAINER_NAME")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	DsName := user + ":" + password + "@" + connectMethod + "(" + containerName + ":" + port + ")/" + name

	err := errors.New("")
	DB, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DB.ShowSQL(true)
	DB.SetMaxOpenConns(2)
	DB.Sync2(model.Todo{})
	fmt.Println("init data base ok")
}
