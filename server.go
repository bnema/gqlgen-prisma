package main

import (
	"graphql-prisma-api/graph"
	"graphql-prisma-api/prisma/db"
	"log"
	"net/http"
	"os"

	// read the .env file
	_ "github.com/joho/godotenv/autoload"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	// Instantiate a new Prisma client
	prisma := db.NewClient()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	// Setup protocol scheme for the Prisma client

	// Create a new resolver, passing the Prisma client to its Prisma field
	resolver := &graph.Resolver{Prisma: prisma}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	// get the errors from the Prisma client
	if err := prisma.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
