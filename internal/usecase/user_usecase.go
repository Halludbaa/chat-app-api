package usecase

import (
	"chatross-api/internal/entity"
	"chatross-api/internal/helper"
	rerror "chatross-api/internal/helper/error"
	"chatross-api/internal/model"
	"chatross-api/internal/model/converter"
	"chatross-api/internal/repository"
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserUsecase struct {
	DB	*gorm.DB
	UserRepository *repository.UserRepository
	Validate *validator.Validate
}

func NewUserUsecase(db *gorm.DB, userRepository *repository.UserRepository, validate *validator.Validate) *UserUsecase{
	return &UserUsecase{
		DB: db,
		Validate: validate,
		UserRepository: userRepository,
	}
}

func (uu *UserUsecase) Refresh(ctx context.Context, request *model.TokenRequest) (*model.TokenResponse, error){
	
	if err := uu.Validate.Struct(request); err != nil {
		return nil, rerror.ErrBadReq
	}

	userID, err := helper.ValidateRefreshToken(request.RefreshToken)
	if err != nil {
		return nil, err
	}

	access_token, err := helper.GenerateAccessToken(userID)
	if err != nil {
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
		return nil, rerror.ErrBadReq // Bad Request
	}

	total, err := c.UserRepository.CountByEmail(tx, request.Email)
	if err != nil {
		return nil, rerror.ErrInternalServer // Internal Server Error
	}

	if total > 0 {
		return nil, rerror.ErrConflict // Error Conflict
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, rerror.ErrInternalServer// Internal Server Error
	}

	newUser := &entity.User{
		Email: request.Email,
		Name: request.Name,
		Password: string(password),
	}

	if err := c.UserRepository.Create(tx, newUser); err != nil {
		return nil, rerror.ErrInternalServer
	}

	if err := tx.Commit().Error; err != nil {
		return nil, rerror.ErrInternalServer
	}

	return converter.UserToResponse(newUser), nil
}

func (c *UserUsecase) Login(ctx context.Context, request *model.LoginUserRequest) (*model.TokenResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil{
		return nil, rerror.ErrBadReq
	}
	
	user := new(entity.User)
	if err := c.UserRepository.FindByEmail(tx, user, request.Email); err != nil {
		return nil, rerror.ErrNotFound
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil{
		return nil, rerror.ErrBadReq
	}

	if err := tx.Commit().Error; err != nil {
		return nil, rerror.ErrInternalServer
	}

	access_token, err := helper.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refresh_token, err := helper.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	
	token := &model.TokenResponse{
		RefreshToken: refresh_token,
		AccessToken: access_token,
	}

	return token, nil
}

func (c *UserUsecase) Verify(ctx context.Context, id int64) (*entity.User, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	
	user := new(entity.User)

	if err := c.UserRepository.FindById(tx, user, id); err != nil {
		return nil, rerror.ErrBadReq
	} 

	if err := tx.Commit().Error; err != nil {
		return nil, rerror.ErrInternalServer
	}
	
	return user, nil
}