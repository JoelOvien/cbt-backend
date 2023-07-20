package controllers

import (
	"errors"
	"strconv"
	"time"

	"github.com/JoelOvien/cbt-backend/middleware"
	"github.com/JoelOvien/cbt-backend/models"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ExamTimetableController is a constructor for DepartmentController
type ExamTimetableController struct {
	DB *gorm.DB
}

// NewExamTimetableController is a constructor for ExamTimetableController
func NewExamTimetableController(DB *gorm.DB) ExamTimetableController {
	return ExamTimetableController{DB}
}

// CreateExamTimetable registered a new course for a student
func (ec *ExamTimetableController) CreateExamTimetable(ctx *fiber.Ctx) error {

	var examTimetable models.ExamTimetable

	err := middleware.DeserializeUserAndCheckUserType(ctx, "ADMIN")
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// Parse body into struct
	err = ctx.BodyParser(&examTimetable)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	errors := models.ValidateStruct(examTimetable)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	now := time.Now()

	examTimetable.ID = uuid.New()
	examTimetable.DateCreated = now
	examTimetable.DateUpdated = now

	// create a new registered course
	result := ec.DB.Table("EXAM_TIMETABLE").Create(&examTimetable)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "time_table": examTimetable})

}

// FindAllExamTimetable fetches all timetables from the db
func (ec *ExamTimetableController) FindAllExamTimetable(ctx *fiber.Ctx) error {

	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var ExamTimetables []models.ExamTimetable

	err := middleware.DeserializeUserAndCheckUserType(ctx, "ADMIN")
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// get all ExamTimetables in Department table
	result := ec.DB.Table("EXAM_TIMETABLE").Limit(intLimit).Offset(offset).Find(&ExamTimetables)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(ExamTimetables), "exam_timetable": ExamTimetables})
}

// FindAllExamTimetableBySemester fetches timetable by semester
func (ec *ExamTimetableController) FindAllExamTimetableBySemester(ctx *fiber.Ctx) error {

	semester := ctx.Params("Semester")

	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var ExamTimetable []models.ExamTimetable

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// get all ExamTimetable in Department table
	result := ec.DB.Table("EXAM_TIMETABLE").Find(&ExamTimetable, "Semester = ?", semester).Limit(intLimit).Offset(offset)

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

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(ExamTimetable), "exam_timetable": ExamTimetable})
}

// FindAllExamTimetableByCourseCode fetches timetable by semester
func (ec *ExamTimetableController) FindAllExamTimetableByCourseCode(ctx *fiber.Ctx) error {

	courseCode := ctx.Params("CourseCode")

	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var ExamTimetable []models.ExamTimetable

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// get all ExamTimetable in Department table
	result := ec.DB.Table("EXAM_TIMETABLE").Find(&ExamTimetable, "CourseCode = ?", courseCode).Limit(intLimit).Offset(offset)

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

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(ExamTimetable), "exam_timetable": ExamTimetable})
}
