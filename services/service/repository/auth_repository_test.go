package repository_test

import (
	"context"
	"os"
	"testing"

	"github.com/AECInfraconnect/AEC-FORM/models"
	"github.com/AECInfraconnect/AEC-FORM/proto/proto_models"
	"github.com/AECInfraconnect/AEC-FORM/server"
	"github.com/AECInfraconnect/AEC-FORM/services/auth/mocks"
	_auth_repo "github.com/AECInfraconnect/AEC-FORM/services/auth/repository"
	"github.com/AECInfraconnect/AEC-FORM/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestRegisterUserUnit(t *testing.T) {
	var (
		objIdStr = "66e03c4e1ef34f18c0b150a1"
		user     *models.User
		mt       *mtest.T
		objId    primitive.ObjectID
		err      error
	)

	if objId, err = primitive.ObjectIDFromHex(objIdStr); err != nil {
		t.Errorf("Error while converting hex to ObjectID: %v", err)
	}

	user = &models.User{
		Id:       objId,
		Username: "test",
		Password: utils.HashPassword("test"),
	}

	mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	// go test -run ^TestRegisterUserUnit$/^success$ -v

	mt.Run("success", func(mt *mtest.T) {
		var (
			authRepoMock *mocks.AuthRepository
			response     *proto_models.RegisterResponse
			err          error
		)
		authRepoMock = new(mocks.AuthRepository)

		mt.AddMockResponses(mtest.CreateSuccessResponse()) // mock response

		authRepo := _auth_repo.NewAuthRepositoryImpl(mt.DB)

		if response, err = authRepo.RegisterUser(context.Background(), user); err != nil {
			t.Errorf("Error while registering user: %v", err)
		}

		assert.NotNil(t, response)
		assert.Nil(t, err)

		assert.Equal(t, user.Id.Hex(), response.XId)
		assert.Equal(t, user.Username, response.Username)

		authRepoMock.AssertExpectations(t)
	})

	// go test -run ^TestRegisterUserUnit$/^error$ -v

	mt.Run("error", func(mt *mtest.T) {
		var (
			authRepoMock *mocks.AuthRepository
			response     *proto_models.RegisterResponse
			err          error
		)

		authRepoMock = new(mocks.AuthRepository)

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Message: "mock insert error",
			Code:    11000,
		}))

		authRepo := _auth_repo.NewAuthRepositoryImpl(mt.DB)

		response, err = authRepo.RegisterUser(context.Background(), user)

		assert.Error(t, err)
		assert.Nil(t, response)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "mock insert error")

		authRepoMock.AssertExpectations(t)

	})

}

func TestRegisterUserIntg(t *testing.T) {
	var (
		objId = primitive.NewObjectID()
		mdb   *mongo.Database
		err   error
	)

	if err = os.Setenv("MONGODB_CONNECTION_URI", "mongodb://root:123456@0.0.0.0:1114/"); err != nil {
		t.Errorf("Error while setting .env: %v", err)
	}

	serv := server.Server{
		MONGODB_CONNECTION_URI: os.Getenv("MONGODB_CONNECTION_URI"),
	}

	mdb = serv.SetUpMongoDBTest(context.Background())

	// go test -run ^TestRegisterUserIntg$/^success$ -v
	t.Run("success", func(t *testing.T) {
		var (
			user     = new(models.User)
			ctx      = context.Background()
			response *proto_models.RegisterResponse
			err      error
		)
		user = &models.User{
			Id:       objId,
			Username: "test",
			Password: utils.HashPassword("test"),
		}
		authRepo := _auth_repo.NewAuthRepositoryImpl(mdb)

		if response, err = authRepo.RegisterUser(ctx, user); err != nil {
			t.Errorf("Error while registering user: %v", err)
		}

		assert.Nil(t, err)
		assert.NotNil(t, response)

		assert.Equal(t, user.Id.Hex(), response.XId)
		assert.Equal(t, user.Username, response.Username)

	})

	// go test -run ^TestRegisterUserIntg$/^error$ -v
	t.Run("error", func(t *testing.T) {
		var (
			user     = new(models.User)
			ctx      = context.Background()
			objIdStr = "66dbf9fb5ec964676c63ce04"
			objId    primitive.ObjectID
			err      error
		)
		objId, _ = primitive.ObjectIDFromHex(objIdStr)
		user = &models.User{
			Id:       objId,
			Username: "test",
			Password: utils.HashPassword("test"),
		}
		authRepo := _auth_repo.NewAuthRepositoryImpl(mdb)

		_, err = authRepo.RegisterUser(ctx, user)

		assert.Error(t, err)
		assert.NotNil(t, err)

		assert.Contains(t, err.Error(), "E11000 duplicate key error collection: aec-form.users index: _id_ dup key: { _id: ObjectId('66e1208887efe36606197967') }")

	})

}
