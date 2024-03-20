package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	pbuser "grpc/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateUser(name, phone, website, password, dni, address string, conn pbuser.UserServiceClient) string {
	ctx := context.Background()
	req := &pbuser.CreateUserRequest{
		Name: name,
		Phone: phone,
		Website: website,
		Password: password,
		Dni: dni,
		Address: address,
	}

	res, err := conn.CreateUser(ctx, req)
	if err != nil {
		checkStatusCode(err)
	}
	return res.GetMessage()
}

func UpdateUser(id uint32, name, phone, website, password, dni, address string, conn pbuser.UserServiceClient) string {
	ctx := context.Background()
	user := &pbuser.User{
		Id: id,
		Name: name,
		Phone: phone,
		Website: website,
		Password: password,
		Dni: dni,
		Address: address,

	}
	req := &pbuser.UpdateUserRequest{
		User: user,
	}

	res, err := conn.UpdateUser(ctx, req)
	if err != nil {
		checkStatusCode(err)
	}
	return res.GetMessage()
}

func DeletUser(id uint32, conn pbuser.UserServiceClient) string {
	ctx := context.Background()
	req := &pbuser.DeleteUserRequest{
		Id: id,
	}

	res, err := conn.DeleteUser(ctx, req)
	if err != nil {
		checkStatusCode(err)
	}
	return res.GetMessage()
}

func GetUsers(conn pbuser.UserServiceClient) {
	req := &pbuser.ListUserRequest{}
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	res, err := conn.GetUser(ctx, req)
	if err != nil {
		checkStatusCode(err)
	}

	for c, user := range res.GetUsers() {
		fmt.Printf("%d: {\n",c)
		fmt.Printf("\t%d\n\t%s\n\t%s\n\t%s\n\t%s\n", user.Id,user.Name,user.Phone,user.Address,user.Website)
		fmt.Println("}")
	}
}


func checkStatusCode(er error) {
	if s, ok := status.FromError(er); ok {
			switch s.Code() {
				case codes.InvalidArgument, codes.Aborted, codes.Internal, codes.Unauthenticated:
					log.Fatalf("Output -> %s: %s", s.Code(), s.Message())
				default:
					log.Fatal(s.Message())
				
			}
	}else {
		panic(er)
	}

}