package handler

import (
	"time"

	"github.com/AECInfraconnect/AEC-FORM/models"
	"github.com/AECInfraconnect/AEC-FORM/proto/proto_models"
	"github.com/AECInfraconnect/AEC-FORM/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FormatRegisterUserReq(request *proto_models.RegisterRequest) *models.User {
	return &models.User{
		Id:        primitive.NewObjectID(),
		Username:  request.Username,
		Password:  utils.HashPassword(request.Password),
		CreatedAt: utils.TimeNow(time.Now()),
	}
}
