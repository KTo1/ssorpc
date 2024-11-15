package ssov1

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (request *LoginRequest) Validate() error {
	if request.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "Email is required")
	}

	return nil
}
