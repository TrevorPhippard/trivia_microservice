package main

import (
	"trivia/services/account"
	"trivia/services/game"
	"trivia/services/quiz"

	"github.com/99designs/gqlgen/graphql"
)

type Server struct {
	accountClient *account.Client
	quizClient    *quiz.Client
	gameClient    *game.Client
}

func NewGraphQLServer(accountUrl, quizUrl, gameUrl string) (*Server, error) {
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}

	quizClient, err := quiz.NewClient(quizUrl)
	if err != nil {
		return nil, err
	}

	gameClient, err := game.NewClient(gameUrl)
	if err != nil {
		return nil, err
	}

	return &Server{
		accountClient,
		quizClient,
		gameClient,
	}, nil
}

type mutationResolver struct {
	server *Server
}

// MutationResolver defines the interface for mutation resolvers.
type MutationResolver interface {
	// Add your mutation methods here, e.g., CreateGame(ctx context.Context, input GameInput) (*Game, error)
}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

type queryResolver struct {
	server *Server
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}
}

type accountResolver struct {
	server *Server
}

func (s *Server) Account() AccountResolver {
	return &accountResolver{
		server: s,
	}
}

type NewExecutableSchema struct {
	server *Server
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
