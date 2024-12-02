package usecase

import (
	"context"

	"github.com/AECInfraconnect/AEC-FORM/models"
	"github.com/AECInfraconnect/AEC-FORM/proto/proto_models"
	"github.com/AECInfraconnect/AEC-FORM/services/auth"
)

type authUsecase struct {
	authRepo auth.AuthRepository
}

func NewAuthUsecaseImpl(authRepo auth.AuthRepository) auth.AuthUsecase {
	return &authUsecase{authRepo: authRepo}
}

func (u *authUsecase) RegisterUserProcessing(ctx context.Context, user *models.User) (*proto_models.RegisterResponse, error) {
	return u.authRepo.RegisterUser(ctx, user)
}

func (u *authUsecase) LoginUserProcessing(ctx context.Context, user *proto_models.LoginRequest) (*models.User, error) {
	return u.authRepo.LoginUser(ctx, user)
}
