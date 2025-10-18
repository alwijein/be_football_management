package service

import (
	"errors"

	"github.com/alwijein/be_football/internal/domain"
	"github.com/alwijein/be_football/internal/dto"
	"github.com/alwijein/be_football/internal/repository"
	"github.com/alwijein/be_football/internal/utils"
	"gorm.io/gorm"
)

// AuthService handles authentication business logic
type AuthService interface {
	Register(req *dto.RegisterRequest) (*domain.User, error)
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	GetUserByID(id uint) (*dto.UserDTO, error)
	UpdateProfile(id uint, req *dto.UpdateProfileRequest) (*dto.UserDTO, error)
	ChangePassword(id uint, req *dto.ChangePasswordRequest) error
}

type authService struct {
	userRepo repository.UserRepository
}

// NewAuthService creates new auth service
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) (*domain.User, error) {
	// Check if username exists
	existingUser, err := s.userRepo.FindByUsername(req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// Check if email exists
	existingUser, err = s.userRepo.FindByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		FullName: req.FullName,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// Find user by username
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	// Check password
	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
		User: dto.UserDTO{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			FullName: user.FullName,
			PhotoURL: user.PhotoURL,
		},
	}, nil
}

func (s *authService) GetUserByID(id uint) (*dto.UserDTO, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &dto.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
		PhotoURL: user.PhotoURL,
	}, nil
}

func (s *authService) UpdateProfile(id uint, req *dto.UpdateProfileRequest) (*dto.UserDTO, error) {
	// Get user
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Check if email already used by another user
	if req.Email != user.Email {
		existingUser, err := s.userRepo.FindByEmail(req.Email)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if existingUser != nil && existingUser.ID != id {
			return nil, errors.New("email already used by another user")
		}
	}

	// Update user
	user.Email = req.Email
	user.FullName = req.FullName
	user.PhotoURL = req.PhotoURL

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
		PhotoURL: user.PhotoURL,
	}, nil
}

func (s *authService) ChangePassword(id uint, req *dto.ChangePasswordRequest) error {
	// Get user
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Verify old password
	if !utils.CheckPassword(req.OldPassword, user.Password) {
		return errors.New("old password is incorrect")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	// Update password
	user.Password = hashedPassword
	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	return nil
}
