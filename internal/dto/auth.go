package dto

// RegisterRequest represents register request
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required,min=3,max=100"`
}

// LoginRequest represents login request
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents login response
type LoginResponse struct {
	Token string  `json:"token"`
	User  UserDTO `json:"user"`
}

// UserDTO represents user data transfer object
type UserDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	PhotoURL string `json:"photo_url,omitempty"`
}

// UpdateProfileRequest represents update profile request
type UpdateProfileRequest struct {
	FullName string `json:"full_name" binding:"omitempty,min=3,max=100"`
	Email    string `json:"email" binding:"omitempty,email"`
	PhotoURL string `json:"photo_url" binding:"omitempty"`
}

// ChangePasswordRequest represents change password request
type ChangePasswordRequest struct {
	OldPassword             string `json:"old_password" binding:"required"`
	NewPassword             string `json:"new_password" binding:"required,min=6"`
	NewPasswordConfirmation string `json:"new_password_confirmation" binding:"required,eqfield=NewPassword"`
}
