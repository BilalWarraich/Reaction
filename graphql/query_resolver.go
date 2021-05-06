package main

import (
	"context"
	"log"
	"time"
)

type queryResolver struct {
	server *Server
}

func (r *queryResolver) ReactionTypes(ctx context.Context, id *string) ([]*ReactionTypes, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// Get single
	if id != nil {
		r, err := r.server.accountClient.GetReactionTypeID(ctx, *id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return []*ReactionTypes{{
			ID:   r.ID,
			Reacts: r.Reacts,
		}}, nil
	}


	accountList, err := r.server.accountClient.GetReactionTypes(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var accounts []*ReactionTypes
	for _, a := range accountList {
		account := &ReactionTypes{
			ID:   a.ID,
			Reacts: a.Reacts,
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *queryResolver) GetUserByPostID(ctx context.Context, postId *string, reactType *string) ([]*Reactions, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()



	accountList, err := r.server.accountClient.GetUserByPostID(ctx, *postId, *reactType)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var accounts []*Reactions
	for _, a := range accountList {
		account := &Reactions{
			ID:   a.ID,
			UserID: a.UserId,
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *queryResolver) GetReactByUserID(ctx context.Context, userId *string) ([]*Reactions, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()



	accountList, err := r.server.accountClient.GetReactByUserID(ctx, *userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var accounts []*Reactions
	for _, a := range accountList {
		account := &Reactions{
			ID:   a.ID,
			PostID: a.PostId,
			ReactType: a.ReactType,
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *queryResolver) TotalReactionCount(ctx context.Context, postId *string) (*string, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// Get single
		a, err := r.server.accountClient.TotalReactionCount(ctx, *postId)
		if err != nil {
			log.Println(err)
			return nil, err
	
		}
		return &a, nil
	}

	func (r *queryResolver) ReactionCountByType(ctx context.Context, postId *string, reactType *string) (*string, error) {
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
	
		// Get single
			a, err := r.server.accountClient.ReactionCountByType(ctx, *postId, *reactType)
			if err != nil {
				log.Println(err)
				return nil, err
		
			}
			return &a, nil
		}
	

