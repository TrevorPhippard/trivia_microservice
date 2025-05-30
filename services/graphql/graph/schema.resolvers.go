package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.73

import (
	"context"
	"fmt"
	"trivia/services/graphql/graph/model"
)

// CreateQuiz is the resolver for the createQuiz field.
func (r *mutationResolver) CreateQuiz(ctx context.Context, input model.NewQuiz) (*model.Quiz, error) {
	panic(fmt.Errorf("not implemented: CreateQuiz - createQuiz"))
}

// CreateQuestion is the resolver for the createQuestion field.
func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.NewQuestion) (*model.Question, error) {
	panic(fmt.Errorf("not implemented: CreateQuestion - createQuestion"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	r.users = append(r.users, user)
	return user, nil
}

// Quizes is the resolver for the quizes field.
func (r *queryResolver) Quizes(ctx context.Context) ([]*model.Quiz, error) {
	panic(fmt.Errorf("not implemented: Quizes - quizes"))
}

// Questions is the resolver for the questions field.
func (r *queryResolver) Questions(ctx context.Context) ([]*model.Question, error) {
	panic(fmt.Errorf("not implemented: Questions - questions"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// Options is the resolver for the options field.
func (r *questionResolver) Options(ctx context.Context, obj *model.Question) (string, error) {
	panic(fmt.Errorf("not implemented: Options - options"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Question returns QuestionResolver implementation.
func (r *Resolver) Question() QuestionResolver { return &questionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type questionResolver struct{ *Resolver }
