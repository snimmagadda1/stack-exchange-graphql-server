package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	gorilla "github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/snimmagadda1/graphql-api/graph"
	"github.com/snimmagadda1/graphql-api/graph/generated"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const defaultPort = "8080"

var db *gorm.DB

func initDB() {
	var err error
	cnx := os.Getenv("UNAME") + ":" + os.Getenv("PASS")
	dataSourceName := cnx + "@tcp(" + os.Getenv("SERVER") + ":3306)/stacke?parseTime=true"
	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		log.Panic("failed to connect database")
	}

	// db.AutoMigrate(&model.User{})
}

func main() {
	godotenv.Load()
	initDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// srv type handler
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))
	logHandler := gorilla.LoggingHandler(os.Stdout, srv)
	http.Handle("/", gorilla.LoggingHandler(os.Stdout, playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", logHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
