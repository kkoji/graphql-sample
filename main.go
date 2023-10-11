package main

import (
	"database/sql"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/kkoji/gqlgen-todos/graph"
	"github.com/kkoji/gqlgen-todos/graph/generated"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open(
		"postgres",
		"host=localhost port=5432 user=postgres dbname=db sslmode=disable",
	)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{DB: db},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
