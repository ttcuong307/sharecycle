package gapi

import (
	"context"
	gapi "sharecycle/internal/gapi/models"
	"sharecycle/internal/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	userInfo, err := s.h.User.GetUserByID(ctx, req.UserId)
	if err != nil {
		s.h.Log.Error(err, "grpc - v1 - GetUserByID - Cannot get User Info")
		return nil, status.Errorf(codes.Internal, "Cannot get User Info")
	}

	resp := &pb.GetUserResponse{User: gapi.ConvertUser(userInfo)}

	return resp, nil
}
