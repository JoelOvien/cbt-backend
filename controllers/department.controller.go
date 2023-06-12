package controllers

import (
	"github.com/JoelOvien/cbt-backend/middleware"
	"github.com/JoelOvien/cbt-backend/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
	"time"
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
func (dc *DepartmentController) FetchAllDepartments(ctx *fiber.Ctx) error {

	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var Departments []models.Department

	err := middleware.DeserializeUserAndCheckUserType(ctx, "ADMIN")
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// get all Departments in Department table
	result := dc.DB.Table("DEPARTMENT").Limit(intLimit).Offset(offset).Find(&Departments)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(Departments), "Departments": Departments})

}

// CreateDepartment creates a new college in the COLLEGE table
func (dc *DepartmentController) CreateDepartment(ctx *fiber.Ctx) error {

	var department models.Department

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// Parse body into struct
	err = ctx.BodyParser(&department)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	errors := models.ValidateStruct(department)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// check if college with same ID exists
	existingDepartment := &models.Department{}
	existingDepartmentResult := dc.DB.Table("DEPARTMENT").Where("DepartmentID = ?", department.DepartmentID).First(&existingDepartment)

	if existingDepartmentResult.Error == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  fiber.StatusConflict,
			"message": "A department with this ID already exists",
		})
	}

	now := time.Now()

	department.DepartmentStatus = 1
	department.DateCreated = now
	department.DateUpdated = now

	// create new college
	result := dc.DB.Table("DEPARTMENT").Create(&department)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "department": department})

}
