package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"strings"
)

func JWTAuthenticationInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Error: %v\n", r)
		}
	}()

	data, ok := metadata.FromIncomingContext(ctx)
	methodName, _ := grpc.Method(ctx)

	if methodName == "/internal.ChatAppService/LoginUser" || methodName == "/internal.ChatAppService/RegisterUser" {
		return handler(ctx, req)
	}

	if !ok || len(data["authorization"]) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "Unauthorization")
	}

	tokenString := strings.TrimPrefix(data["authorization"][0], "Bearer ")

	if tokenString == "" {
		return nil, status.Errorf(codes.Unauthenticated, "Token is required")
	}

	// check jwt token access full
	if _, err := verifyToken(tokenString); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	return handler(ctx, req)
}
