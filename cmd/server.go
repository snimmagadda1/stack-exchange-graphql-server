package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	gorilla "github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/snimmagadda1/graphql-api/generated"
	"github.com/snimmagadda1/graphql-api/internal/dal"
	_ "github.com/snimmagadda1/graphql-api/internal/logger"
	graph "github.com/snimmagadda1/graphql-api/internal/resolver"
)

const defaultPort = "8080"

func main() {
	logrus.Info("Loading environment")
	godotenv.Load()
	db, err := dal.Init()

	if err != nil {
		logrus.Error(err)
		logrus.Fatal("Failed to connect database")
	}

	port := os.Getenv("PORT")
	if port == "" {
		logrus.Debugf("using default port %s", defaultPort)
		port = defaultPort
	}

	// srv type handler
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db.DB}}))
	http.Handle("/", gorilla.LoggingHandler(logrus.StandardLogger().Out, playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", srv)

	logrus.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	logrus.Fatal(http.ListenAndServe(":"+port, nil))
}
