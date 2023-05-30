package middleware

import (
	"fmt"
	"github.com/JoelOvien/cbt-backend/database"
	"github.com/JoelOvien/cbt-backend/models"
	"github.com/JoelOvien/cbt-backend/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// DeserializeUser is a middleware to deserialize user
func DeserializeUser(ctx *fiber.Ctx) error {
	var accessToken string
	cookie := new(fiber.Cookie)
	cookie.Name = "access_token"

	authorizationHeader := ctx.Get("Authorization")
	fmt.Println(authorizationHeader)

	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
		fmt.Println(accessToken)
	} else {
		accessToken = cookie.Value
		fmt.Println(accessToken)
	}

	if accessToken == "" {
		fmt.Println("error access token is empty stirng")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "You are not logged in",
		})
	}

	config, _ := database.LoadConfig(".")
	fmt.Println("after setting config", accessToken)

	//called this function to validate the access token and extract the payload (user’s ID) we stored in it
	sub, err := utils.ValidateToken(accessToken, config.AccessTokenPublicKey)
	if err != nil {
		fmt.Println("error", err.Error())
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})

	}

	var user models.Users
	result := database.DB.Table("USERS").First(&user, "UserID = ?", fmt.Sprint(sub))
	if result.Error != nil {
		fmt.Println("error", result.Error)
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  fiber.StatusForbidden,
			"message": "The user belonging to this token no logger exists",
		})

	}

	ctx.Set("currentUser", user.UserID)

	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"status":  fiber.StatusForbidden,
		"message": "The user belonging to this token no logger exists",
	})
}
