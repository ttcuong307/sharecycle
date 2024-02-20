package gapi

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"sharecycle/internal/models"
	"sharecycle/internal/pb"
)

func ConvertUser(data *models.User) *pb.User {
	return &pb.User{
		Id:    data.ID,
		Name:  data.Name,
		Email: data.Email,
		//Password: data.Password,
		Gender:    data.Gender,
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
	}
}
