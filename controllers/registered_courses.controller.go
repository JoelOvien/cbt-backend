package controllers

import (
	"errors"
	"strconv"
	"time"

	"github.com/JoelOvien/cbt-backend/middleware"
	"github.com/JoelOvien/cbt-backend/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// RegisteredCourseController is a constructor for DepartmentController
type RegisteredCourseController struct {
	DB *gorm.DB
}

// NewRegisteredCourseController is a constructor for RegisteredCourseController
func NewRegisteredCourseController(DB *gorm.DB) RegisteredCourseController {
	return RegisteredCourseController{DB}
}

// RegisterCourseForStudent registered a new course for a student
func (rcc *RegisteredCourseController) RegisterCourseForStudent(ctx *fiber.Ctx) error {

	var registeredCourse models.RegisteredCourses

	err := middleware.DeserializeUserAndCheckUserType(ctx, "ADMIN")
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// Parse body into struct
	err = ctx.BodyParser(&registeredCourse)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	errors := models.ValidateStruct(registeredCourse)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	now := time.Now()

	registeredCourse.DateCreated = now
	registeredCourse.DateUpdated = now

	// create a new registered course
	result := rcc.DB.Table("REGISTERED_COURSES").Create(&registeredCourse)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "registered_courses": registeredCourse})

}

// FetchAllRegisteredCourses fetches all reg courses from the db
func (rcc *RegisteredCourseController) FetchAllRegisteredCourses(ctx *fiber.Ctx) error {

	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var RegisteredCourses []models.RegisteredCourses

	err := middleware.DeserializeUserAndCheckUserType(ctx, "ADMIN")
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// get all RegisteredCourses in Department table
	result := rcc.DB.Table("REGISTERED_COURSES").Limit(intLimit).Offset(offset).Find(&RegisteredCourses)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(RegisteredCourses), "registered_courses": RegisteredCourses})
}

// FetchStudentRegisteredCourses fetches all reg courses of a student from the db
func (rcc *RegisteredCourseController) FetchStudentRegisteredCourses(ctx *fiber.Ctx) error {

	studentID := ctx.Params("UserID")

	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var RegisteredCourses []models.RegisteredCourses

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// get all RegisteredCourses in Department table
	result := rcc.DB.Table("REGISTERED_COURSES").Find(&RegisteredCourses, "UserID = ?", studentID).Limit(intLimit).Offset(offset)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"status":  fiber.StatusBadGateway,
				"message": "No registration record found",
			})
		}

		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(RegisteredCourses), "registered_courses": RegisteredCourses})
}

// FetchStudentRegisteredCoursesBySemester fetches all reg courses of a student from the db
func (rcc *RegisteredCourseController) FetchStudentRegisteredCoursesBySemester(ctx *fiber.Ctx) error {

	studentID := ctx.Params("UserID")
	semester := ctx.Params("Semester")

	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var RegisteredCourses []models.RegisteredCourses

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// get all RegisteredCourses in Department table
	result := rcc.DB.Table("REGISTERED_COURSES").Where("UserID = ? AND Semester = ?", studentID, semester).Limit(intLimit).Offset(offset).Find(&RegisteredCourses)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"status":  fiber.StatusBadGateway,
				"message": "No registration record found",
			})
		}

		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(RegisteredCourses), "registered_students": RegisteredCourses})
}

// FetchStudentRegisteredForCourseByCourseCode fetches all students registred for a course
func (rcc *RegisteredCourseController) FetchStudentRegisteredForCourseByCourseCode(ctx *fiber.Ctx) error {

	courseCode := ctx.Query("CourseCode")

	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var RegisteredCourses []models.RegisteredCourses

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// get all RegisteredCourses in Department table
	result := rcc.DB.Table("REGISTERED_COURSES").Where("CourseCode = ?", courseCode).Limit(intLimit).Offset(offset).Find(&RegisteredCourses)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"status":  fiber.StatusBadGateway,
				"message": "No registration record found",
			})
		}

		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(RegisteredCourses), "registered_students": RegisteredCourses})
}
