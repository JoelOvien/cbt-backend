package controllers

import (
	"fmt"

	"github.com/JoelOvien/cbt-backend/database"
	"github.com/JoelOvien/cbt-backend/middleware"
	"github.com/JoelOvien/cbt-backend/models"
	"github.com/JoelOvien/cbt-backend/utils"

	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// AuthController is a struct to store db
type AuthController struct {
	DB *gorm.DB
}

// NewAuthController is constructor for AuthController
func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

// SignInUser is a function to sign in a user
func (ac *AuthController) SignInUser(ctx *fiber.Ctx) error {
	var payload *models.Users

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	var user models.Users
	result := ac.DB.Table("USERS").First(&user, "UserID = ?", payload.UserID)

	// If user does not exist
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "UserID does not exist",
			"Error":   result.Error,
		})
	}

	// If user exists, check password
	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid Password",
		})
	}

	// Update user's last access date
	if err := UpdateLastAccessDate(user.UserID); err != nil {
		// Handle the error
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to update last access date",
			"Error":   err.Error(),
		})
	}

	config, _ := database.LoadConfig(".") // Load database config

	// Generate JWT access token
	accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, user.UserID, config.AccessTokenPrivateKey)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Generate JWT refresh token
	refreshToken, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.UserID, config.RefreshTokenPrivateKey)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// COOKIES
	accessTokenMaxAge := time.Duration(config.AccessTokenMaxAge) * time.Minute
	accessTokenExpiry := time.Now().Add(accessTokenMaxAge)

	refreshTokenMaxAge := time.Duration(config.RefreshTokenMaxAge) * time.Minute
	refreshTokenExpiry := time.Now().Add(refreshTokenMaxAge)

	cookie := new(fiber.Cookie)
	// Set access token cookie
	cookie.Name = "access_token"
	cookie.Value = accessToken
	cookie.Expires = accessTokenExpiry
	cookie.HTTPOnly = true
	cookie.Secure = false
	cookie.Path = "/"
	cookie.Path = "localhost"
	ctx.Cookie(cookie)

	// Set refresh token cookie
	cookie.Name = "refresh_token"
	cookie.Value = refreshToken
	cookie.Expires = refreshTokenExpiry
	cookie.HTTPOnly = true
	cookie.Secure = false
	cookie.Path = "/"
	cookie.Path = "localhost"
	ctx.Cookie(cookie)

	// set login token
	cookie.Name = "logged_in"
	cookie.Value = "true"
	cookie.Expires = accessTokenExpiry
	cookie.HTTPOnly = false
	cookie.Secure = false
	cookie.Path = "/"
	cookie.Path = "localhost"
	ctx.Cookie(cookie)

	userResponse := &models.UserResponse{
		UserID:         user.UserID,
		Surname:        user.Surname,
		FirstName:      user.FirstName,
		EmailAddress:   user.EmailAddress,
		UserStatus:     "Active",
		UserType:       user.UserType,
		DateCreated:    user.DateCreated,
		DateUpdated:    user.DateUpdated,
		LastAccessDate: user.LastAccessDate,
		DepartmentID:   user.DepartmentID,
		Role:           user.Role,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":        "success",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"data":          userResponse})

}

// CreateUser is a function to create a new user
func (ac *AuthController) CreateUser(ctx *fiber.Ctx) error {
	var payload *models.Users

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	err := middleware.DeserializeUserAndCheckUserType(ctx, "ADMIN")
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	if payload.UserType != "STAFF" && payload.UserType != "STUDENT" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid UserType",
		})
	}

	// Check if a user exists with the same UserID
	existingUser := &models.Users{}
	existingUserResult := ac.DB.Table("USERS").Where("UserID = ?", payload.UserID).First(existingUser)

	// Check if there was an error
	if existingUserResult.Error == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  fiber.StatusConflict,
			"message": "User with this UserID already exists",
		})
	}

	if existingUserResult.Error != nil {
		if existingUserResult.Error == gorm.ErrRecordNotFound {
			// User does not exist, proceed with creating a new user
			now := time.Now()

			payload.Password = hashedPassword
			payload.DateCreated = now
			payload.DateUpdated = now
			payload.LastAccessDate = now

			newUser := models.Users{
				UserID:         payload.UserID,
				Surname:        payload.Surname,
				FirstName:      payload.FirstName,
				EmailAddress:   payload.EmailAddress,
				Password:       payload.Password,
				UserStatus:     "Active",
				UserType:       payload.UserType,
				DateCreated:    payload.DateCreated,
				DateUpdated:    payload.DateUpdated,
				LastAccessDate: payload.LastAccessDate,
				DepartmentID:   payload.DepartmentID,
				Role:           payload.Role,
			}

			result := ac.DB.Table("USERS").Create(&newUser)

			if result.Error != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  fiber.StatusInternalServerError,
					"message": result.Error.Error(),
				})
			}

		} else {

			// Other error occurred
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status":  fiber.StatusConflict,
				"message": "User with this UserID already exists",
				"Error":   existingUserResult.Error,
			})
		}
	}

	userResponse := &models.UserResponse{
		UserID:         payload.UserID,
		Surname:        payload.Surname,
		FirstName:      payload.FirstName,
		EmailAddress:   payload.EmailAddress,
		UserStatus:     "Active",
		UserType:       payload.UserType,
		DateCreated:    payload.DateCreated,
		DateUpdated:    payload.DateUpdated,
		LastAccessDate: payload.LastAccessDate,
		DepartmentID:   payload.DepartmentID,
		Role:           payload.Role,
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"message": userResponse}})

}

// RefreshAccessToken is a function to refresh access token
func (ac *AuthController) RefreshAccessToken(ctx *fiber.Ctx) error {

	cookie := new(fiber.Cookie)
	config, _ := database.LoadConfig(".")

	sub, err := utils.ValidateToken(cookie.Value, config.RefreshTokenPublicKey)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusForbidden,
			"message": err.Error(),
		})

	}

	var user models.Users
	result := ac.DB.First(&user, "UserID = ?", fmt.Sprint(sub))

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, user.UserID, config.AccessTokenPrivateKey)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusForbidden,
			"message": err.Error(),
		})
	}
	accessTokenMaxAge := time.Duration(config.AccessTokenMaxAge) * time.Minute
	accessTokenExpiry := time.Now().Add(accessTokenMaxAge)

	// Set access token cookie
	cookie.Name = "access_token"
	cookie.Value = accessToken
	cookie.Expires = accessTokenExpiry
	cookie.HTTPOnly = true
	cookie.Secure = false
	cookie.Path = "/"
	cookie.Path = "localhost"
	ctx.Cookie(cookie)

	// set login token
	cookie.Name = "logged_in"
	cookie.Value = "true"
	cookie.Expires = accessTokenExpiry
	cookie.HTTPOnly = false
	cookie.Secure = false
	cookie.Path = "/"
	cookie.Path = "localhost"
	ctx.Cookie(cookie)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":       fiber.StatusOK,
		"access_token": accessToken,
	})

}

// UpdateLastAccessDate is a function to update the last access date of a user
func UpdateLastAccessDate(userID string) error {
	// Get the current time
	currentTime := time.Now()

	// Update the user's last access date in the database
	err := database.DB.Model(&models.Users{}).Where("UserID = ?", userID).Update("LastAccessDate", currentTime).Error
	if err != nil {
		// Handle the error
		return err
	}

	return nil
}
