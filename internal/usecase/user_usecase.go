package usecase

import (
	"chatross-api/internal/entity"
	"chatross-api/internal/helper"
	rerror "chatross-api/internal/helper/error"
	"chatross-api/internal/model"
	"chatross-api/internal/model/converter"
	"chatross-api/internal/repository"
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUsecase struct {
	DB	*gorm.DB
	Log *logrus.Logger
	UserRepository *repository.UserRepository
	Validate *validator.Validate
}

func NewUserUsecase(db *gorm.DB, userRepository *repository.UserRepository, validate *validator.Validate, log *logrus.Logger) *UserUsecase{
	return &UserUsecase{
		DB: db,
		Log: log,
		Validate: validate,
		UserRepository: userRepository,
	}
}

func (uu *UserUsecase) Refresh(ctx context.Context, request *model.TokenRequest) (*model.TokenResponse, error){
	
	if err := uu.Validate.Struct(request); err != nil {
		uu.Log.Error("Error Bad Request")
		return nil, rerror.ErrBadReq
	}

	userID, err := helper.ValidateRefreshToken(request.RefreshToken)
	if err != nil {
		uu.Log.Error("Invalid Refresh Token")
		return nil, err
	}

	access_token, err := helper.GenerateAccessToken(userID)
	if err != nil {
		uu.Log.Error("Failed Generated Access Token")
		return nil, err
	}

	token := &model.TokenResponse{
		AccessToken: access_token,
	}

	return token, nil

}

func (c *UserUsecase) Create(ctx context.Context, request *model.RegisterUserRequest ) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Error("Invalid JSON Structure")
		return nil, rerror.ErrBadReq // Bad Request
	}

	total, err := c.UserRepository.CountById(tx, request.Username)
	if err != nil {
		c.Log.Error("Failed Access Database")
		return nil, rerror.ErrInternalServer // Internal Server Error
	}
	log.Printf("User Count: %d", total)
	if total > 0 {
		c.Log.Error("Conflict In Database")
		return nil, rerror.ErrConflict // Error Conflict
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Error("Failed Hashing Password")
		return nil, rerror.ErrInternalServer// Internal Server Error
	}

	newUser := &entity.User{
		ID: request.Username,
		Email: request.Email,
		Password: string(password),
	}

	if err := c.UserRepository.Create(tx, newUser); err != nil {
		c.Log.Error("Failed To Add User")
		return nil, rerror.ErrInternalServer
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Error("Failed Session Commit")
		return nil, rerror.ErrInternalServer
	}

	return converter.UserToResponse(newUser), nil
}

func (c *UserUsecase) Login(ctx context.Context, request *model.LoginUserRequest) (*model.TokenResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil{
		c.Log.Error("Invalid JSON Structure")
		return nil, rerror.ErrBadReq
	}
	
	user := new(entity.User)
	if err := c.UserRepository.FindById(tx, user, request.Username); err != nil {
		c.Log.WithField("Error", err).Error("Failed Access Database")
		return nil, rerror.ErrNotFound
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil{
		c.Log.WithField("Error", err).Error("Invalid Password")
		return nil, rerror.ErrBadReq
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithField("Error", err).Error("Failed Session Commit")
		return nil, rerror.ErrInternalServer
	}

	access_token, err := helper.GenerateAccessToken(user.ID)
	if err != nil {
		c.Log.WithField("Error", err).Error("Failed Generate Access Token")
		return nil, err
	}

	refresh_token, err := helper.GenerateRefreshToken(user.ID)
	if err != nil {
		c.Log.WithField("Error", err).Error("Failed Generate Refresh Token")
		return nil, err
	}

	
	token := &model.TokenResponse{
		RefreshToken: refresh_token,
		AccessToken: access_token,
	}

	return token, nil
}

func (c *UserUsecase) Verify(ctx context.Context, id string) (*entity.User, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	
	user := new(entity.User)

	if err := c.UserRepository.FindById(tx, user, id); err != nil {
		c.Log.WithField("Error", err).Error("Failed Access Database")
		return nil, rerror.ErrBadReq
	} 

	if err := tx.Commit().Error; err != nil {
		return nil, rerror.ErrInternalServer
	}
	
	return user, nil
}