package controllers

import (
	"github.com/JoelOvien/cbt-backend/middleware"
	"github.com/JoelOvien/cbt-backend/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
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

	middleware.DeserializeUser(ctx)

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
