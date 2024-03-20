package cmd

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc"
)

var token string

const key = "auth_token"

func SetToken(key string) {
	token = key
}

func UnaryAuthInterceptor(ctx context.Context, method string, req, reply any, conn *grpc.ClientConn, 
invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = metadata.AppendToOutgoingContext(ctx, key,token)
	err := invoker(ctx,method,req,reply,conn,opts...)
	return err
}