package test

import (
	"context"
	pbtest "grpc/user"
	"log"
	"net"
	"server/api/user"
	"server/models"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)



const buffSize = 1024 * 1024
var list *bufconn.Listener

func init() {
	models.Migrate("./../test.db")
	models.DialDb("./../test.db", "./../logs/error.log")

	ts := grpc.NewServer()
	user := user.NewUser()
	list = bufconn.Listen(buffSize)

	pbtest.RegisterUserServiceServer(ts, user)
	go func() {
		if err := ts.Serve(list); err != nil {
			log.Fatalf("%s", err)
		}
	}()

}

func buildDialer(context.Context, string) (net.Conn, error) {
	return list.Dial()
}

func newClient( t *testing.T) (*grpc.ClientConn, pbtest.UserServiceClient) {
	ctx := context.TODO()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(buildDialer), creds)
	if err != nil {
		t.Fatal(err)
	}
	return conn, pbtest.NewUserServiceClient(conn)
}