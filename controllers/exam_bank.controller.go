package controllers

import (
	"errors"
	"strconv"
	"time"

	"github.com/JoelOvien/cbt-backend/middleware"
	"github.com/JoelOvien/cbt-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ExamBankController struct
type ExamBankController struct {
	DB *gorm.DB
}

// NewExamBankController function
func NewExamBankController(DB *gorm.DB) ExamBankController {
	return ExamBankController{DB}
}

// CheckSubmittedAnswer checks if the AnswerProvided in the response is equal to the CorrectAnswer for the given QuestionID.
func CheckSubmittedAnswer(response models.ExamBank, question models.ExamBank) bool {
	return response.AnswerProvided == question.CorrectAnswer
}

// SubmitAnswer submits a user's answers for multiple questions
func (qc *ExamBankController) SubmitAnswer(ctx *fiber.Ctx) error {
	var answerDetails []models.ExamBank

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// Parse body into slice of structs
	err = ctx.BodyParser(&answerDetails)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Perform validation and other processing for each answerDetails entry
	for i, answer := range answerDetails {
		validationError := models.ValidateStruct(answer)
		if validationError != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(validationError)
		}

		isAnswerCorrect := CheckSubmittedAnswer(answer, answer) // Replace the second argument with the appropriate reference to the correct answer for this question

		if isAnswerCorrect {
			answerDetails[i].AnswerMark = 1 // Update the value in the slice directly
		} else {
			answerDetails[i].AnswerMark = 0 // Update the value in the slice directly
		}

		now := time.Now()
		answerDetails[i].DateCreated = now // Update the value in the slice directly
		answerDetails[i].ID = uuid.New()   // Update the value in the slice directly
		answerDetails[i].DateUpdated = now // Update the value in the slice directly
	}

	// Create new questions in the database
	result := qc.DB.Table("EXAM_BANK").Create(&answerDetails)

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

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "questions": answerDetails})
}

// SubmitSingleAnswer submits a user's answer for a question
func (qc *ExamBankController) SubmitSingleAnswer(ctx *fiber.Ctx) error {

	var answerDetails models.ExamBank

	err := middleware.DeserializeUserAndCheckUserType(ctx, "STUDENT")
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// Parse body into struct
	err = ctx.BodyParser(&answerDetails)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	validationError := models.ValidateStruct(answerDetails)
	if validationError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(validationError)
	}

	isAnswerCorrect := CheckSubmittedAnswer(answerDetails, answerDetails)

	now := time.Now()
	if isAnswerCorrect {
		answerDetails.AnswerMark = 1
	} else {
		answerDetails.AnswerMark = 0
	}
	answerDetails.ID = uuid.New()
	answerDetails.DateCreated = now
	answerDetails.DateUpdated = now

	// create a new question
	result := qc.DB.Table("EXAM_BANK").Create(&answerDetails)

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

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "question": answerDetails})
}

// FindAll questions in QUESTIONSBANK table
func (qc *ExamBankController) FindAll(ctx *fiber.Ctx) error {
	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")
	var courseCode = ctx.Query("courseCode")
	var questionID = ctx.Query("questionID")
	var semester = ctx.Query("semester")
	var session = ctx.Query("session")
	var userID = ctx.Query("userID")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var submittedAnswers []models.ExamBank

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// Build the query directly in the database based on the presence of courseCode and questionID parameters
	query := qc.DB.Table("EXAM_BANK")
	if courseCode != "" {
		query = query.Where("CourseCode = ?", courseCode)
	}
	if questionID != "" {
		query = query.Where("AnswerTypeID = ?", questionID)
	}
	if semester != "" {
		query = query.Where("Semester = ?", semester)
	}
	if session != "" {
		query = query.Where("Session = ?", session)
	}
	if userID != "" {
		query = query.Where("UserID = ?", userID)
	}

	// Fetch the filtered questions from the database
	result := query.Limit(intLimit).Offset(offset).Find(&submittedAnswers)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success",
		"results":   len(submittedAnswers),
		"questions": submittedAnswers})
}
