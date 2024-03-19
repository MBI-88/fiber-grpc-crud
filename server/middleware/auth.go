package middleware

import (
	"context"
	
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var token string
const key = "auth_token"


func SetToken(key string) {
	token = key
}

func ValidateAuthToken(ctx context.Context) (context.Context, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.DataLoss, "Not token sent")
	}

	if t, ok := meta[key]; ok {
		if len(t) != 1 {
			return nil, status.Error(codes.InvalidArgument, "auth_token should contain only 1 value")
		}
		if t[0] != token {
			return nil, status.Error(codes.Unauthenticated, "Incorrect auth_token")
		}
	}else {
		return nil, status.Error(codes.Unauthenticated, "Failed to get auth_token")
	}

	return ctx, nil
}