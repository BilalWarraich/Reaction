//go:generate protoc ./reactionTypes.proto --go_out=./pb
package reaction

import (
	"context"
	"fmt"
	"net"

	pb "github.com/Reaction/reaction/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
	pb.UnimplementedReactionServiceServer
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterReactionServiceServer(serv, &grpcServer{s,pb.UnimplementedReactionServiceServer{}})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostReactionTypes(ctx context.Context, r *pb.PostReactionTypesRequest) (*pb.PostReactionTypesResponse, error) {
	a, err := s.service.PostReactionTypes(ctx, r.Reacts)
	if err != nil {
		return nil, err
	}
	return &pb.PostReactionTypesResponse{ReactionTypes: &pb.ReactionTypes{
		Id:   a.ID,
		Reacts: a.Reacts,
	}}, nil
}

func (s *grpcServer) GetReactionTypeID(ctx context.Context, r *pb.GetReactionTypeIDRequest) (*pb.GetReactionTypeIDResponse, error) {
	a, err := s.service.GetReactionTypeID(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetReactionTypeIDResponse{
		ReactionTypes: &pb.ReactionTypes{
			Id:   a.ID,
			Reacts: a.Reacts,
		},
	}, nil
}

func (s *grpcServer) GetReactionTypes(ctx context.Context, r *pb.GetReactionTypesRequest) (*pb.GetReactionTypesResponse, error) {
	res, err := s.service.GetReactionTypes(ctx)
	if err != nil {
		return nil, err
	}
	reactionTypes := []*pb.ReactionTypes{}
	for _, p := range res {
		reactionTypes = append(
			reactionTypes,
			&pb.ReactionTypes{
				Id:   p.ID,
				Reacts: p.Reacts,
			},
		)
	}
	return &pb.GetReactionTypesResponse{ReactionTypes: reactionTypes}, nil
}

func (s *grpcServer) PostReactions(ctx context.Context, r *pb.PostReactionsRequest) (*pb.PostReactionsResponse, error) {
	a, err := s.service.PostReactions(ctx, r.ReactType, r.PostId, r.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.PostReactionsResponse{Reactions: &pb.Reactions{
		Id:   a.ID,
		ReactType: a.ReactType,
		PostId: a.PostId,
		UserId: a.UserId,
	}}, nil
}

func (s *grpcServer) GetUserByPostID(ctx context.Context, r *pb.GetUserByPostIDRequest) (*pb.GetUserByPostIDResponse, error) {
	res, err := s.service.GetUserByPostID(ctx, r.PostId,r.ReactType)
	if err != nil {
		return nil, err
	}
	reactionTypes := []*pb.Reactions{}
	for _, p := range res {
		reactionTypes = append(
			reactionTypes,
			&pb.Reactions{
				Id:   p.ID,
				UserId: p.UserId,
			},
		)
	}
	return &pb.GetUserByPostIDResponse{Reactions : reactionTypes}, nil
}

func (s *grpcServer) GetReactByUserID(ctx context.Context, r *pb.GetReactByUserIDRequest) (*pb.GetReactByUserIDResponse, error) {
	res, err := s.service.GetReactByUserID(ctx, r.UserId)
	if err != nil {
		return nil, err
	}
	reactionTypes := []*pb.Reactions{}
	for _, p := range res {
		reactionTypes = append(
			reactionTypes,
			&pb.Reactions{
				Id:   p.ID,
				PostId: p.PostId,
				ReactType: p.ReactType,
			},
		)
	}
	return &pb.GetReactByUserIDResponse{Reactions : reactionTypes}, nil
}

func (s *grpcServer) DeleteReactionByPostID(ctx context.Context, r *pb.DeleteReactionByPostIDRequest) (*pb.DeleteReactionByPostIDResponse, error) {
	a, err := s.service.DeleteReactionByPostID(ctx, r.PostId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteReactionByPostIDResponse{Ok: a}, nil
}

func (s *grpcServer) DeleteReactionByUserPostID(ctx context.Context, r *pb.DeleteReactionByUserPostIDRequest) (*pb.DeleteReactionByUserPostIDResponse, error) {
	a, err := s.service.DeleteReactionByUserPostIDRequest(ctx, r.PostId, r.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteReactionByUserPostIDResponse{Ok: a}, nil
}

func (s *grpcServer) UpdateReactions(ctx context.Context, r *pb.UpdateReactionsRequest) (*pb.UpdateReactionsResponse, error) {
	a, err := s.service.UpdateReactions(ctx, r.ReactType, r.PostId, r.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateReactionsResponse{Reactions: &pb.Reactions{
		Id:   a.ID,
		ReactType: a.ReactType,
		PostId: a.PostId,
		UserId: a.UserId,
	}}, nil
}

func (s *grpcServer) TotalReactionCount(ctx context.Context, r *pb.TotalReactionCountRequest) (*pb.TotalReactionCountResponse, error) {
	a, err := s.service.TotalReactionCount(ctx, r.PostId)
	if err != nil {
		return nil, err
	}
	return &pb.TotalReactionCountResponse{Count: a}, nil
}

func (s *grpcServer) ReactionCountByType(ctx context.Context, r *pb.ReactionCountByTypeRequest) (*pb.ReactionCountByTypeResponse, error) {
	a, err := s.service.ReactionCountByType(ctx, r.PostId, r.ReactType)
	if err != nil {
		return nil, err
	}
	return &pb.ReactionCountByTypeResponse{Count: a}, nil
}
