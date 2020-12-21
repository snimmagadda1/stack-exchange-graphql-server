package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	gorilla "github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/snimmagadda1/graphql-api/graph"
	"github.com/snimmagadda1/graphql-api/graph/generated"
	_ "github.com/snimmagadda1/graphql-api/graph/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormlog "gorm.io/gorm/logger"
)

const defaultPort = "8080"

var db *gorm.DB

func initDB() {
	logrus.Info("Initializing datasource ...")
	var err error
	cnx := os.Getenv("UNAME") + ":" + os.Getenv("PASS")
	schema, exist := os.LookupEnv("SCHEMA")
	if !exist {
		schema = "stacke"
	}
	dataSourceName := cnx + "@tcp(" + os.Getenv("SERVER") + ":3306)/" + schema + "?parseTime=true"

	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: gormlog.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      gormlog.Info,
				Colorful:      true,
			},
		)})

	if err != nil {
		fmt.Println(err)
		logrus.Fatal("Failed to connect database")
	}
}

func main() {
	logrus.Info("Loading environment")
	godotenv.Load()
	initDB()

	port := os.Getenv("PORT")
	if port == "" {
		logrus.Debugf("using default port %s", defaultPort)
		port = defaultPort
	}

	// srv type handler
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))
	http.Handle("/", gorilla.LoggingHandler(logrus.StandardLogger().Out, playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", srv)

	logrus.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	logrus.Fatal(http.ListenAndServe(":"+port, nil))
}
