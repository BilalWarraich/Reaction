package main

import (
	"context"
	"errors"
	"log"
	"time"
)

var (
	ErrInvalidParameter = errors.New("invalid parameter")
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) CreateReactionTypes(ctx context.Context, in ReactionTypesInput) (*ReactionTypes, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	a, err := r.server.accountClient.PostReactionTypes(ctx, in.Reacts)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &ReactionTypes{
		ID:   a.ID,
		Reacts: a.Reacts,
	}, nil
}
func (r *mutationResolver) CreateReactions(ctx context.Context, in ReactionsInput) (*Reactions, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	a, err := r.server.accountClient.PostReactions(ctx, in.ReactType, in.PostID, in.UserID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Reactions{
		ID:   a.ID,
		ReactType: a.ReactType,
		PostID: a.PostId,
		UserID: a.UserId,
	}, nil
}

func (r *mutationResolver) UpdateReactions(ctx context.Context, in ReactionsInput) (*Reactions, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	a, err := r.server.accountClient.UpdateReactions(ctx, in.ReactType, in.PostID, in.UserID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Reactions{
		ID:   a.ID,
		ReactType: a.ReactType,
		PostID: a.PostId,
		UserID: a.UserId,
	}, nil
}

func (r *mutationResolver) DeleteReactionByPostID(ctx context.Context, postId *string) (*string, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	a, err := r.server.accountClient.DeleteReactionByPostID(ctx,  *postId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &a, nil
}
func (r *mutationResolver) DeleteReactionByUserPostIDRequest(ctx context.Context, postId *string, userId *string) (*string, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	a, err := r.server.accountClient.DeleteReactionByUserPostID(ctx,  *postId, *userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
     
	return &a, nil
}

