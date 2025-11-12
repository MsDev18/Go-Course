package user

import (
	"E-04/entity"
	"E-04/pkg/phonenumber"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(userID uint) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	auth AuthGenerator
	repo Repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserInfo struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}

type RegisterResponse struct {
	User UserInfo `json:"user"`
}

func New(authGenerator AuthGenerator, repo Repository) Service {
	return Service{
		repo: repo,
		auth: authGenerator,
	}
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {

	// TODO - we should verify phone number with verification code
	// Validate Phone Number
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}
	// Check Uniqueness of phone number
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unExpected error : %w", err)
		}

		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("number is not unique")
		}
	}
	// Validate Name
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name length should be greatest than 3")
	}

	// TODO - Check The Password With Regex Pattern
	// Validate Password
	if len(req.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf("password len should be greater than 8")
	}

	user := entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    getMD5Hash(req.Password),
	}
	// Create New User In Storage
	// TODO - assign anonymose struct
	createdUser, err := s.repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unExpected error : %w", err)
	}
	// Return Created User

	return RegisterResponse{UserInfo{
		ID:          createdUser.ID,
		PhoneNumber: createdUser.PhoneNumber,
		Name:        createdUser.Name,
	}}, nil
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
	User   UserInfo `json:"user"`
	Tokens Tokens   `json:"tokens"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	// TODO - it would be better to user two separate method for existence check and getUserByPhoneNumber
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error : %w", err)
	}
	if !exist {
		return LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}

	// 3. Compare user.password with the req.password
	if user.Password != getMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error : %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error : %w", err)
	}

	return LoginResponse{
		User: UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
		},
		Tokens: Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type ProfileRequest struct {
	UserID uint
}
type ProfileResponse struct {
	Name string `json:"name"`
}

// All Request inputes for interactor/service should be sanitized
func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	// 1. GetUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		// I don't expect repository call return "record not found" error ,
		//  because I assume the interactor input is sanitized
		// TODO - we can use Reich Error
		return ProfileResponse{}, fmt.Errorf("unexpected error : %w", err)
	}
	return ProfileResponse{Name: user.Name}, nil
}
