package controllers

import (
	"fmt"
	"github.com/JoelOvien/cbt-backend/database"
	"github.com/JoelOvien/cbt-backend/models"
	"github.com/JoelOvien/cbt-backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// CollegeController is a struct to store our db instance
type CollegeController struct {
	DB *gorm.DB
}

// NewCollegeController is a constructor for CollegeController
func NewCollegeController(DB *gorm.DB) CollegeController {
	return CollegeController{DB}
}

// FetchAllColleges fetches all colleges from the COLLEGE table
func (cc *CollegeController) FetchAllColleges(ctx *fiber.Ctx) error {

	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var colleges []models.College

	// if user token from Authorization header is wrong from the DesertializeUser middleware, return error
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
	userResult := database.DB.Table("USERS").First(&user, "UserID = ?", fmt.Sprint(sub))
	if userResult.Error != nil {
		fmt.Println("error", userResult.Error)
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  fiber.StatusForbidden,
			"message": "The user belonging to this token no logger exists",
		})

	}

	ctx.Set("currentUser", user.UserID)

	// get all colleges in COLLEGE table
	result := cc.DB.Table("COLLEGE").Limit(intLimit).Offset(offset).Find(&colleges)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(colleges), "colleges": colleges})

}
