package reaction

import (
	"context"

	pb "github.com/Reaction/reaction/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.ReactionServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewReactionServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) PostReactionTypes(ctx context.Context, reacts string) (*ReactionTypes, error) {
	r, err := c.service.PostReactionTypes(
		ctx,
		&pb.PostReactionTypesRequest{Reacts: reacts},
	)
	if err != nil {
		return nil, err
	}
	return &ReactionTypes{
		ID:     r.ReactionTypes.Id,
		Reacts: r.ReactionTypes.Reacts,
	}, nil
}

func (c *Client) GetReactionTypeID(ctx context.Context, id string) (*ReactionTypes, error) {
	r, err := c.service.GetReactionTypeID(
		ctx,
		&pb.GetReactionTypeIDRequest{Id: id},
	)
	if err != nil {
		return nil, err
	}
	return &ReactionTypes{
		ID:     r.ReactionTypes.Id,
		Reacts: r.ReactionTypes.Reacts,
	}, nil
}

func (c *Client) GetReactionTypes(ctx context.Context) ([]ReactionTypes, error) {
	r, err := c.service.GetReactionTypes(
		ctx,
		&pb.GetReactionTypesRequest{},
	)
	if err != nil {
		return nil, err
	}
	reactionTypes := []ReactionTypes{}
	for _, a := range r.ReactionTypes {
		reactionTypes = append(reactionTypes, ReactionTypes{
			ID:     a.Id,
			Reacts: a.Reacts,
		})
	}
	return reactionTypes, nil
}

func (c *Client) PostReactions(ctx context.Context, reactType string, postId string, userId string) (*Reactions, error) {
	r, err := c.service.PostReactions(
		ctx,
		&pb.PostReactionsRequest{ReactType: reactType, PostId: postId, UserId: userId},
	)
	if err != nil {
		return nil, err
	}
	return &Reactions{
		ID:        r.Reactions.Id,
		ReactType: r.Reactions.ReactType,
		PostId:    r.Reactions.PostId,
		UserId:    r.Reactions.UserId,
	}, nil
}

func (c *Client) GetUserByPostID(ctx context.Context, postId string, reactType string) ([]Reactions, error) {
	r, err := c.service.GetUserByPostID(
		ctx,
		&pb.GetUserByPostIDRequest{PostId: postId, ReactType: reactType},
	)
	if err != nil {
		return nil, err
	}
	reactionTypes := []Reactions{}
	for _, a := range r.Reactions {
		reactionTypes = append(reactionTypes, Reactions{
			ID:     a.Id,
			UserId: a.UserId,
		})
	}
	return reactionTypes, nil
}

func (c *Client) GetReactByUserID(ctx context.Context, userId string) ([]Reactions, error) {
	r, err := c.service.GetReactByUserID(
		ctx,
		&pb.GetReactByUserIDRequest{UserId: userId},
	)
	if err != nil {
		return nil, err
	}
	reactionTypes := []Reactions{}
	for _, a := range r.Reactions {
		reactionTypes = append(reactionTypes, Reactions{
			ID:        a.Id,
			PostId:    a.PostId,
			ReactType: a.ReactType,
		})
	}
	return reactionTypes, nil
}

func (c *Client) DeleteReactionByPostID(ctx context.Context, postId string) (string, error) {
	r, err := c.service.DeleteReactionByPostID(
		ctx,
		&pb.DeleteReactionByPostIDRequest{PostId: postId},
	)
	if err != nil {
		return "", err
	}
	return r.Ok, nil
}

func (c *Client) DeleteReactionByUserPostID(ctx context.Context, postId string, userId string) (string, error) {
	r, err := c.service.DeleteReactionByUserPostID(
		ctx,
		&pb.DeleteReactionByUserPostIDRequest{PostId: postId, UserId: userId},
	)
	if err != nil {
		return "", err
	}
	return r.Ok, nil
}

func (c *Client) UpdateReactions(ctx context.Context, reactType string, postId string, userId string) (*Reactions, error) {
	r, err := c.service.UpdateReactions(
		ctx,
		&pb.UpdateReactionsRequest{ReactType: reactType, PostId: postId, UserId: userId},
	)
	if err != nil {
		return nil, err
	}
	return &Reactions{
		ID:        r.Reactions.Id,
		ReactType: r.Reactions.ReactType,
		PostId:    r.Reactions.PostId,
		UserId:    r.Reactions.UserId,
	}, nil
}

func (c *Client) TotalReactionCount(ctx context.Context, postId string) (string, error) {
	r, err := c.service.TotalReactionCount(
		ctx,
		&pb.TotalReactionCountRequest{PostId: postId},
	)
	if err != nil {
		return "", err
	}
	return r.Count, nil
}

func (c *Client) ReactionCountByType(ctx context.Context, postId string, reactType string) (string, error) {
	r, err := c.service.ReactionCountByType(
		ctx,
		&pb.ReactionCountByTypeRequest{PostId: postId, ReactType: reactType},
	)
	if err != nil {
		return "", err
	}
	return r.Count, nil
}
