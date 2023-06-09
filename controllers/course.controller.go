package controllers

import (
	"github.com/JoelOvien/cbt-backend/middleware"
	"github.com/JoelOvien/cbt-backend/models"

	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// CourseController is a struct to store our db instance
type CourseController struct {
	DB *gorm.DB
}

// NewCourseController is a constructor for CourseController
func NewCourseController(DB *gorm.DB) CourseController {
	return CourseController{DB}
}

// FetchAllCourses fetches all DCourses from the Course table
func (cc *CourseController) FetchAllCourses(ctx *fiber.Ctx) error {

	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var Courses []models.Course

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

	// get all DCourses in DCourse table
	result := cc.DB.Table("COURSES").Limit(intLimit).Offset(offset).Find(&Courses)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(Courses), "Courses": Courses})
}

// CreateCourse creates a new Course to the COURSES Table
func (cc *CourseController) CreateCourse(ctx *fiber.Ctx) error {

	var course models.Course

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// Parse body into struct
	err = ctx.BodyParser(&course)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	errors := models.ValidateStruct(course)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// check if course with same CourseCode exists
	existingCourse := &models.Course{}
	existingCourseResult := cc.DB.Table("COURSES").Where("CourseCode = ?", course.CourseCode).First(&existingCourse)

	if existingCourseResult.Error == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  fiber.StatusConflict,
			"message": "A course with this ID already exists",
		})
	}

	var time = time.Now()

	course.DateCreated = time
	course.DateUpdated = time

	// create new Course
	result := cc.DB.Table("COURSES").Create(&course)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "Course": course})

}
