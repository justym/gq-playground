package main

import (
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	g "github.com/justym/gq-playground/graphql-server/graphql"
)

func main() {
	schema, err := g.NewSchema()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	h := handler.New(&handler.Config{
		Schema:   schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
