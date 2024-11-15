package ssov1

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (request *LoginRequest) Validate() error {
	if request.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}

	if request.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}

	if request.GetAppID() == 0 {
		return status.Error(codes.InvalidArgument, "app ID is required")
	}

	return nil
}
