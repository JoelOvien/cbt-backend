package controllers

import (
	"github.com/JoelOvien/cbt-backend/middleware"
	"github.com/JoelOvien/cbt-backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

// UserController is a struct to store our db instance
type UserController struct {
	DB *gorm.DB
}

// NewUserController is a constructor for UserController
func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

// FetchAllUsers fetches all Users from the Users table
func (rc *UserController) FetchAllUsers(ctx *fiber.Ctx) error {
	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var Users []models.UserResponse

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	result := rc.DB.Table("USERS").Limit(intLimit).Offset(offset).Find(&Users)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(Users), "Users": Users})
}
