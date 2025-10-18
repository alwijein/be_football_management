package handler

import (
	"net/http"

	"github.com/alwijein/be_football/internal/dto"
	"github.com/alwijein/be_football/internal/service"
	"github.com/alwijein/be_football/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	loginResp, err := h.authService.Login(&req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	loginResp.User.PhotoURL = utils.GetImageURL(c, loginResp.User.PhotoURL)

	utils.SuccessResponse(c, http.StatusOK, "Login successful", loginResp)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	utils.SuccessResponse(c, http.StatusOK, "Logout successful", nil)
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	user, err := h.authService.GetUserByID(userID.(uint))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	response := dto.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
		PhotoURL: utils.GetImageURL(c, user.PhotoURL),
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile retrieved", response)
}

func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	existingUser, err := h.authService.GetUserByID(userID.(uint))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	// Parse form data
	var req dto.UpdateProfileRequest

	// Get form values (all optional)
	if fullName := c.PostForm("full_name"); fullName != "" {
		req.FullName = fullName
	}
	if email := c.PostForm("email"); email != "" {
		req.Email = email
	}

	// Handle photo upload if provided
	_, err = c.FormFile("photo_url")
	if err == nil {
		// New photo file provided, upload it
		photoPath, err := utils.UploadImage(c, "photo_url", "profiles")
		if err != nil {
			utils.ValidationErrorResponse(c, "Photo upload failed: "+err.Error())
			return
		}

		// Delete old photo if exists
		if existingUser.PhotoURL != "" {
			utils.DeleteImage(existingUser.PhotoURL)
		}

		req.PhotoURL = photoPath
	}

	user, err := h.authService.UpdateProfile(userID.(uint), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response := dto.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
		PhotoURL: utils.GetImageURL(c, user.PhotoURL),
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile updated successfully", response)
}

// ChangePassword changes user password
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	err := h.authService.ChangePassword(userID.(uint), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Password changed successfully", nil)
}
