package handler

import (
	"context"

	"github.com/AECInfraconnect/AEC-FORM/constant"
	"github.com/AECInfraconnect/AEC-FORM/models"
	"github.com/AECInfraconnect/AEC-FORM/proto/proto_models"
	"github.com/AECInfraconnect/AEC-FORM/services/auth"
	"github.com/AECInfraconnect/AEC-FORM/services/token"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authUcase  auth.AuthUsecase
	tokenUcase token.TokenUsecase
}

func NewAuthHandlerImpl(authUcase auth.AuthUsecase, tokenUcase token.TokenUsecase) auth.AuthHandler {
	return &authHandler{
		authUcase:  authUcase,
		tokenUcase: tokenUcase,
	}

}

func (h *authHandler) RegisterUser(c *fiber.Ctx) error {
	var (
		user     = new(models.User)
		ctx      = context.Background()
		request  *proto_models.RegisterRequest
		response *proto_models.RegisterResponse
		err      error
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "status": constant.BAD_REQUEST})
	}
	user = FormatRegisterUserReq(request)
	if response, err = h.authUcase.RegisterUserProcessing(ctx, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "status": constant.INTERNAL_SERVER_ERROR})
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *authHandler) LoginUser(c *fiber.Ctx) error {
	var (
		request  *proto_models.LoginRequest
		response *proto_models.LoginResponse
		userData *models.User
		token    *string
		err      error
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "status": constant.BAD_REQUEST})
	}
	if userData, err = h.authUcase.LoginUserProcessing(c.Context(), request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if token, err = h.tokenUcase.SignToken(userData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	response = &proto_models.LoginResponse{
		Token:  *token,
		Status: "success",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
