package middleware

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"strings"
)

func JWTAuthenticationStreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Error: %v\n", r)
		}
	}()

	data, _ := metadata.FromIncomingContext(ss.Context())

	authorization, err := data["authorization"]
	if !err {
		return status.Errorf(codes.Unauthenticated, "Invalid or expired token")
	}

	token := strings.TrimPrefix(authorization[0], "Bearer ")

	if token == "" {
		return status.Errorf(codes.Unauthenticated, "Token is required")
	}

	// check jwt token access full
	if _, err := verifyToken(token); err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}

	return handler(srv, ss)
}
