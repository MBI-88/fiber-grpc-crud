package user

import (
	"context"

	pb "grpc/user"
	"server/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type user struct {
	m models.User
	pb.UnimplementedUserServiceServer
}

func (u *user) CreateUser(_ context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	u.m.Name =  req.Name
	u.m.Phone = req.Phone
	u.m.Address = req.Address
	u.m.Dni = req.Dni
	u.m.Website = req.Website
	u.m.Password = req.Password

	if err := u.m.CreateUser(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

 	return &pb.CreateUserResponse{Message: "Operation successful!"}, nil
}

func (u *user) UpdateUser(_ context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserRespose, error) {
	u.m.Name =  req.User.Name
	u.m.Phone = req.User.Phone
	u.m.Address = req.User.Address
	u.m.Dni = req.User.Dni
	u.m.Website = req.User.Website
	u.m.Password = req.User.Password
	u.m.ID = req.User.Id

	if err := u.m.UpdateUser(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateUserRespose{Message: "Operation successful!"}, nil
}

func (u *user) DeleteUser(_ context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.Aborted, "ID empty")
	}
	
	u.m.ID = req.Id
	if err := u.m.DeleteUser(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteUserResponse{Message: "Operation successful!"}, nil
}

func (u user) GetUser(_ context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	users, err := u.m.GetUsers()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ListUserResponse{Users: users}, nil
}
