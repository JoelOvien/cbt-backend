package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/JoelOvien/cbt-backend/database"
	"github.com/JoelOvien/cbt-backend/models"
	"github.com/JoelOvien/cbt-backend/utils"
	"github.com/gofiber/fiber/v2"
)

// DeserializeUser is a middleware to deserialize user
func DeserializeUser(ctx *fiber.Ctx) error {
	var accessToken string
	cookie := new(fiber.Cookie)
	cookie.Name = "access_token"

	authorizationHeader := ctx.Get("Authorization")

	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	} else {
		accessToken = cookie.Value
	}

	if accessToken == "" {
		return fmt.Errorf("Unauthorized")

	}

	config, _ := database.LoadConfig(".")

	sub, err := utils.ValidateToken(accessToken, config.AccessTokenPublicKey)
	if err != nil {
		return fmt.Errorf(err.Error())

	}

	var user models.Users
	result := database.DB.Table("USERS").First(&user, "UserID = ?", fmt.Sprint(sub))
	if result.Error != nil {
		return fmt.Errorf(result.Error.Error())

	}

	fmt.Println("user is from deserialize ", user.UserType)
	ctx.Set("currentUser", user.UserType)

	currentUser := ctx.Get("currentUser")
	// Use the userID value
	fmt.Println("user id is ", currentUser)

	return nil
}

// DeserializeUserAndCheckUserType is a middleware to deserialize user
func DeserializeUserAndCheckUserType(ctx *fiber.Ctx, role string) error {
	var accessToken string
	cookie := new(fiber.Cookie)
	cookie.Name = "access_token"

	authorizationHeader := ctx.Get("Authorization")

	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	} else {
		accessToken = cookie.Value
	}

	if accessToken == "" {
		return fmt.Errorf("Unauthorized")

	}

	config, _ := database.LoadConfig(".")

	sub, err := utils.ValidateToken(accessToken, config.AccessTokenPublicKey)
	if err != nil {
		return fmt.Errorf(err.Error())

	}

	var user models.Users
	result := database.DB.Table("USERS").First(&user, "UserID = ?", fmt.Sprint(sub))
	if result.Error != nil {
		return fmt.Errorf(result.Error.Error())

	}

	// fmt.Println("user is from deserialize ", user.UserType)
	// ctx.Set("currentUser", user.UserType)

	// currentUser := ctx.Get("currentUser")
	// // Use the userID value
	// fmt.Println("user id is ", currentUser)

	userType := user.Role
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err

}
