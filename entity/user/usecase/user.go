package usecase

import (
	"be-golang-fiber/entity/user"
	"be-golang-fiber/entity/user/repository"
	"be-golang-fiber/utils/jwt"
	"be-golang-fiber/utils/password"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserUsecase interface {
	UserRegister(um *user.UserModel) (*user.UserModel, error)
	UserLogin(um *user.UserLoginModel) (string, error)
	GetProfile(id int) (*user.UserProfileModel, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(u repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: u,
	}
}

func (u *userUsecase) UserRegister(um *user.UserModel) (*user.UserModel, error) {
	err := validator.New().Struct(um)
	fmt.Println(um)
	if err != nil {
		return um, user.REQUEST_BODY_NOT_VALID
	}
	um.Active = true
	um.RoleID = 1
	um.Password, _ = password.HashPassword(um.Password)
	um.CreatedAt = time.Now()
	res, err := u.userRepository.UserRegister(um)
	if err != nil {
		return um, err
	}
	return res, nil
}

func (u *userUsecase) UserLogin(um *user.UserLoginModel) (string, error) {
	err := validator.New().Struct(um)
	if err != nil {
		return "", user.REQUEST_BODY_NOT_VALID
	}
	userObj, err := u.userRepository.UserLogin(um)
	if err != nil {
		return "", err
	}
	if !userObj.Active {
		return "", user.LOGIN_FAILED_INACTIVE
	}
	loggedIn, err := password.VerifyPassword(um.Password, userObj.Password)
	if err != nil || !loggedIn {
		return "", user.LOGIN_FAILED
	}
	token, err := jwt.CreateJWTToken(userObj)

	return token, nil
}

func (u *userUsecase) GetProfile(id int) (*user.UserProfileModel, error) {
	res, err := u.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	profile := new(user.UserProfileModel)
	profile.Email = res.Email
	profile.Firstname = res.Firstname
	profile.Lastname = res.Lastname
	profile.RoleID = res.RoleID
	profile.ImageName = res.ImageName

	return profile, nil
}
