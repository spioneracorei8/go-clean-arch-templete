package repository

import (
	"context"
	"errors"

	"github.com/AECInfraconnect/AEC-FORM/models"
	"github.com/AECInfraconnect/AEC-FORM/proto/proto_models"
	"github.com/AECInfraconnect/AEC-FORM/services/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type authRepository struct {
	db *mongo.Database
}

func NewAuthRepositoryImpl(db *mongo.Database) auth.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) RegisterUser(ctx context.Context, user *models.User) (*proto_models.RegisterResponse, error) {
	var (
		response *proto_models.RegisterResponse
		result   *mongo.InsertOneResult
		err      error
	)
	if result, err = r.db.Collection("users").InsertOne(ctx, user); err != nil {
		return nil, err
	}
	response = &proto_models.RegisterResponse{
		XId:      result.InsertedID.(primitive.ObjectID).Hex(),
		Username: user.Username,
	}
	return response, nil
}

func (r *authRepository) LoginUser(ctx context.Context, user *proto_models.LoginRequest) (*models.User, error) {
	var (
		userData models.User
		err      error
	)
	if err = r.db.Collection("users").FindOne(ctx, bson.M{"username": user.Username}).Decode(&userData); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return &userData, nil
}
