package controllers

import (
	"github.com/JoelOvien/cbt-backend/middleware"
	"github.com/JoelOvien/cbt-backend/models"

	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

// DepartmentController is a struct to store our db instance
type DepartmentController struct {
	DB *gorm.DB
}

// NewDepartmentController is a constructor for DepartmentController
func NewDepartmentController(DB *gorm.DB) DepartmentController {
	return DepartmentController{DB}
}

// FetchAllDepartments fetches all Departments from the Department table
func (cc *DepartmentController) FetchAllDepartments(ctx *fiber.Ctx) error {

	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var Departments []models.Department

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	currentUser := ctx.Get("currentUser")
	userID := currentUser
	// Use the userID value
	fmt.Println("user id is ", userID)

	// get all Departments in Department table
	result := cc.DB.Table("DEPARTMENT").Limit(intLimit).Offset(offset).Find(&Departments)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(Departments), "Departments": Departments})

}
