package reaction

import (
	"context"

	"github.com/segmentio/ksuid"
)

type Service interface {
	PostReactionTypes(ctx context.Context, reacts string) (*ReactionTypes, error)
	GetReactionTypeID(ctx context.Context, id string) (*ReactionTypes, error)
	GetReactionTypes(ctx context.Context) ([]ReactionTypes, error)
	PostReactions(ctx context.Context, reactType string, postId string, userId string) (*Reactions, error)
	GetUserByPostID(ctx context.Context, postId string, reactType string) ([]Reactions, error)
	DeleteReactionByPostID(ctx context.Context, postId string) (string, error)
	DeleteReactionByUserPostIDRequest(ctx context.Context, postId string, userId string) (string, error)
	UpdateReactions(ctx context.Context, reactType string, postId string, userId string) (*Reactions, error)
	TotalReactionCount(ctx context.Context, postId string) (string, error)
	ReactionCountByType(ctx context.Context, postId string,reactType string, ) (string, error)

	




}

type ReactionTypes struct {
	ID   string `json:"id"`
	Reacts string `json:"reacts"`
}

type Reactions struct {
	ID   string `json:"id"`
	ReactType string `json:"reactType"`
	PostId string `json:"postId"`
	UserId string `json:"userId"`

}

type reactionTypeService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &reactionTypeService{r}
}

func (s *reactionTypeService) PostReactionTypes(ctx context.Context, reacts string) (*ReactionTypes, error) {
	a := &ReactionTypes{
		Reacts: reacts,
		ID:   ksuid.New().String(),
	}
	if err := s.repository.PostReactionTypes(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *reactionTypeService) GetReactionTypeID(ctx context.Context, id string) (*ReactionTypes, error) {
	return s.repository.GetReactionTypeID(ctx, id)
}

func (s *reactionTypeService) GetReactionTypes(ctx context.Context) ([]ReactionTypes, error) {

	return s.repository.GetReactionTypes(ctx)
}

func (s *reactionTypeService) PostReactions(ctx context.Context, reactType string, postId string, userId string) (*Reactions, error) {
	a := &Reactions{
		ID:   ksuid.New().String(),
		ReactType: reactType,
		PostId: postId,
		UserId: userId,

	}
	if err := s.repository.PostReactions(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *reactionTypeService) GetUserByPostID(ctx context.Context, postId string, reactType string) ([]Reactions, error) {
	return s.repository.GetUserByPostID(ctx, postId, reactType)
}

func (s *reactionTypeService) DeleteReactionByPostID(ctx context.Context,postId string) (string, error) {

	if err := s.repository.DeleteReactionByPostID(ctx, postId); err != nil {
		return "", err
	}
	return "Deleted", nil
}

func (s *reactionTypeService) DeleteReactionByUserPostIDRequest(ctx context.Context,postId string ,userId string) (string, error) {

	if err := s.repository.DeleteReactionByUserPostIDRequest(ctx, postId, userId); err != nil {
		return "", err
	}
	return "Deleted", nil
}

func (s *reactionTypeService) UpdateReactions(ctx context.Context, reactType string, postId string, userId string) (*Reactions, error) {
	a := &Reactions{
		ReactType: reactType,
		PostId: postId,
		UserId: userId,

	}
	if err := s.repository.UpdateReactions(ctx, *a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *reactionTypeService) TotalReactionCount(ctx context.Context,postId string) (string, error) {

	count, err := s.repository.TotalReactionCount(ctx, postId); 
	if err != nil {
		return "", err
	}
	return count, nil
}

func (s *reactionTypeService) ReactionCountByType(ctx context.Context,postId string , reactType string) (string, error) {

	count, err := s.repository.ReactionCountByType(ctx, postId ,reactType); 
	if err != nil {
		return "", err
	}
	return count, nil
}