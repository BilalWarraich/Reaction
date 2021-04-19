package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/Reaction/reaction"

)

type Server struct {
	accountClient *reaction.Client

}

func NewGraphQLServer(reactionUrl string) (*Server, error) {
	// Connect to account service
	reactionUrlClient, err := reaction.NewClient(reactionUrl)
	if err != nil {
		return nil, err
	}



	return &Server{
		reactionUrlClient,
	}, nil
}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}
}



func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers:  s,
	
	})
}
