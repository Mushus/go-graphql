package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	go_graphql "github.com/Mushus/go-graphql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := gorm.Open("sqlite3", "test.db")
	db = db.Debug()
	db.AutoMigrate(&WebSite{}, &Document{}, &Category{})

	resolver := Resolver{
		db: db,
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(go_graphql.NewExecutableSchema(go_graphql.Config{Resolvers: &resolver})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
