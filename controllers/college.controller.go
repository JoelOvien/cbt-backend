package controllers

import (
	"github.com/JoelOvien/cbt-backend/middleware"
	"github.com/JoelOvien/cbt-backend/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
	"time"
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

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

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

// CreateCollege creates a new college in the COLLEGE table
func (cc *CollegeController) CreateCollege(ctx *fiber.Ctx) error {

	var collegePayload models.College

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// Parse body into struct
	err = ctx.BodyParser(&collegePayload)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	errors := models.ValidateStruct(collegePayload)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// check if college with same ID exists
	existingCollege := &models.College{}
	existingCollegeResult := cc.DB.Table("COLLEGE").Where("CollegeID = ?", collegePayload.CollegeID).First(&existingCollege)

	if existingCollegeResult.Error == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  fiber.StatusConflict,
			"message": "A college with this ID already exists",
		})
	}

	now := time.Now()

	collegePayload.CollegeStatus = 1
	collegePayload.DateCreated = now
	collegePayload.DateUpdated = now

	// create new college
	result := cc.DB.Table("COLLEGE").Create(&collegePayload)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "college": collegePayload})

}
