package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AllowedImageExtensions defines allowed image file extensions
var AllowedImageExtensions = []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}

// MaxFileSize defines maximum file size (5MB)
const MaxFileSize = 5 * 1024 * 1024

// UploadImage handles image file upload
func UploadImage(c *gin.Context, fieldName, uploadDir string) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return "", fmt.Errorf("failed to get file: %v", err)
	}

	// Check file size
	if file.Size > MaxFileSize {
		return "", fmt.Errorf("file size exceeds maximum limit of 5MB")
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !isAllowedExtension(ext) {
		return "", fmt.Errorf("invalid file type. Allowed types: jpg, jpeg, png, gif, webp")
	}

	// Generate unique filename
	filename := generateUniqueFilename(ext)

	// Create upload directory if not exists
	uploadPath := filepath.Join("uploads", uploadDir)
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Save file
	filePath := filepath.Join(uploadPath, filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	// Return relative path for storing in database
	return filepath.Join(uploadDir, filename), nil
}

// DeleteImage deletes an image file
func DeleteImage(filePath string) error {
	if filePath == "" {
		return nil
	}

	fullPath := filepath.Join("uploads", filePath)
	if _, err := os.Stat(fullPath); err == nil {
		return os.Remove(fullPath)
	}
	return nil
}

// ValidateImageFile validates if the uploaded file is valid
func ValidateImageFile(file *multipart.FileHeader) error {
	// Check file size
	if file.Size > MaxFileSize {
		return fmt.Errorf("file size exceeds maximum limit of 5MB")
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !isAllowedExtension(ext) {
		return fmt.Errorf("invalid file type. Allowed types: jpg, jpeg, png, gif, webp")
	}

	return nil
}

// isAllowedExtension checks if file extension is allowed
func isAllowedExtension(ext string) bool {
	for _, allowedExt := range AllowedImageExtensions {
		if ext == allowedExt {
			return true
		}
	}
	return false
}

// generateUniqueFilename generates unique filename using UUID and timestamp
func generateUniqueFilename(ext string) string {
	timestamp := time.Now().Unix()
	uniqueID := uuid.New().String()[:8]
	return fmt.Sprintf("%d_%s%s", timestamp, uniqueID, ext)
}

// GetImageURL returns full URL for image
func GetImageURL(c *gin.Context, imagePath string) string {
	if imagePath == "" {
		return ""
	}

	// Get base URL from request
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	baseURL := fmt.Sprintf("%s://%s", scheme, c.Request.Host)
	return fmt.Sprintf("%s/uploads/%s", baseURL, imagePath)
}
