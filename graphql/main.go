//go:generate go run github.com/99designs/gqlgen
package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

type AppConfig struct {
	ReactionURL string `envconfig:"REACTION_SERVICE_URL"`

}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	s, err := NewGraphQLServer(cfg.ReactionURL)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/graphql", handler.GraphQL(s.ToExecutableSchema()))
	http.Handle("/graphiql", handler.Playground("reaction", "/graphql"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
