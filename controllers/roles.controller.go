package controllers

import (
	"github.com/JoelOvien/cbt-backend/middleware"
	"github.com/JoelOvien/cbt-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// RoleController is a struct to store our db instance
type RoleController struct {
	DB *gorm.DB
}

// NewRoleController is a constructor for RoleController
func NewRoleController(DB *gorm.DB) RoleController {
	return RoleController{DB}
}

// FetchAllRoles fetches all Roles from the Roles table
func (rc *RoleController) FetchAllRoles(ctx *fiber.Ctx) error {
	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var Roles []models.Role

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	result := rc.DB.Table("ROLES").Limit(intLimit).Offset(offset).Find(&Roles)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(Roles), "Roles": Roles})
}

// CreateRole creates a new Role in the Roles table
func (rc *RoleController) CreateRole(ctx *fiber.Ctx) error {

	var role models.Role

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// Parse body into struct
	err = ctx.BodyParser(&role)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	errors := models.ValidateStruct(role)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// check if college with same ID exists
	existingRole := &models.Role{}
	existingRoleResult := rc.DB.Table("ROLES").Where("RoleName = ?", role.RoleID).First(&existingRole)

	if existingRoleResult.Error == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  fiber.StatusConflict,
			"message": "A role with this name already exists",
		})
	}

	// If the error returned contains Duplicate entry we return a 409 status code

	now := time.Now()

	role.RoleID = uuid.New()
	role.DateCreated = now
	role.DateUpdated = now

	// create new college
	result := rc.DB.Table("ROLES").Create(&role)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "role": role})

}
