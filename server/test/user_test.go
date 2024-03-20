package test

import (
	"context"
	"testing"

	pbtest "grpc/user"
)


func TestCreateUser(t *testing.T) {
	conn, c := newClient(t)
	defer conn.Close()

	req := &pbtest.CreateUserRequest{
		Name: "testUser",
		Password: "Testuser@1*",
		Phone: "+59892346512",
		Website: "https://www.test.com",
		Dni: "12345678B",
		Address: "Test address",
	}

	res, err := c.CreateUser(context.TODO(), req)
	if err != nil {
		t.Fatalf("Ouput: %s", err)
	}

	if res.GetMessage() != "Operation successful!" {
		t.Fatal("Unexpected message")
	}
}

func TestUpdateUser(t *testing.T) {
	conn, c := newClient(t)
	defer conn.Close()
	req := &pbtest.ListUserRequest{}

	res, err := c.GetUser(context.TODO(),req)
	if err != nil {
		t.Fatalf("Output %s", err)
	}

	id := res.GetUsers()[0].Id
	user := &pbtest.User{
		Id: uint32(id),
		Name: "TestUpdatedUser",
	}
	req2 := &pbtest.UpdateUserRequest{
		User: user,
	}

	res2, err := c.UpdateUser(context.TODO(), req2)
	if err != nil {
		t.Fatalf("Ouput: %s", err)
	}
 
	if res2.GetMessage() != "Operation successful!" {
		t.Fatal("Unexpected message")
	}
}

func TestDeleteUser(t *testing.T) {
	conn, c := newClient(t)
	defer conn.Close()
	req := &pbtest.ListUserRequest{}

	res, err := c.GetUser(context.TODO(),req)
	if err != nil {
		t.Fatalf("Output %s", err)
	}

	id := res.GetUsers()[0].Id

	req2 := &pbtest.DeleteUserRequest{
		Id: uint32(id),
	}

	res2, err := c.DeleteUser(context.TODO(), req2)
	if err != nil {
		t.Fatalf("Ouput: %s", err)
	}

	if res2.GetMessage() != "Operation successful!" {
		t.Fatal("Unexpected message")
	}
}

func TestGetUsers(t *testing.T) {
	conn, c := newClient(t)
	defer conn.Close()
	req := &pbtest.ListUserRequest{}

	res, err := c.GetUser(context.TODO(),req)
	if err != nil {
		t.Fatalf("Output %s", err)
	}

	if len(res.GetUsers()) > 0 {
		t.Fatalf("GetUsers must be %d > 0", len(res.GetUsers()))
	}

}