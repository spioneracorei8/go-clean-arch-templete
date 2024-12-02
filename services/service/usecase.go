package auth

import (
	"context"

	"github.com/AECInfraconnect/AEC-FORM/models"
	"github.com/AECInfraconnect/AEC-FORM/proto/proto_models"
)

type AuthUsecase interface {
	RegisterUserProcessing(ctx context.Context, user *models.User) (*proto_models.RegisterResponse, error)
	LoginUserProcessing(ctx context.Context, user *proto_models.LoginRequest) (*models.User, error)
}
