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

// QuestionsBankController struct
type QuestionsBankController struct {
	DB *gorm.DB
}

// NewQuestionsBankController function
func NewQuestionsBankController(DB *gorm.DB) QuestionsBankController {
	return QuestionsBankController{DB}
}

// CreateQuestion creates a new question
func (qc *QuestionsBankController) CreateQuestion(ctx *fiber.Ctx) error {

	var quesions models.QuestionsBank

	err := middleware.DeserializeUserAndCheckUserType(ctx, "EXAMINER")
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// Parse body into struct
	err = ctx.BodyParser(&quesions)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	validationError := models.ValidateStruct(quesions)
	if validationError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(validationError)
	}

	now := time.Now()
	courseCode := quesions.CourseCode
	questionNumber := strconv.Itoa(quesions.QuestionNumber)

	quesions.QuestionID = courseCode + "_" + questionNumber
	quesions.DateCreated = now
	quesions.DateUpdated = now

	// create a new question
	result := qc.DB.Table("QUESTIONS_BANK").Create(&quesions)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status":  fiber.StatusConflict,
				"message": "A question with this QuestionID already exists",
			})
		}

		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "question": quesions})
}
