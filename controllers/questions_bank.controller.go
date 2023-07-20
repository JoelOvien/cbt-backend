package controllers

import (
	"errors"
	"math/rand"
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

// Define the linear congruential random number generator function
func linearCongruentialRandom(seed int64, n int) int {
	const (
		m = 2147483647 // Modulus
		a = 48271      // Multiplier
		c = 0          // Increment
	)

	seed = (a*seed + c) % m
	return int(seed % int64(n))
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

// FindAll questions in QUESTIONSBANK table
func (qc *QuestionsBankController) FindAll(ctx *fiber.Ctx) error {
	var page = ctx.Query("page", "1")
	var limit = ctx.Query("limit", "10")
	var courseCode = ctx.Query("courseCode")
	var answerTypeID = ctx.Query("answerTypeID")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var filteredQuestions []models.QuestionsBank

	err := middleware.DeserializeUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": err.Error(),
		})
	}

	// Build the query directly in the database based on the presence of courseCode and answerTypeID parameters
	query := qc.DB.Table("QUESTIONS_BANK")
	if courseCode != "" {
		query = query.Where("CourseCode = ?", courseCode)
	}
	if answerTypeID != "" {
		query = query.Where("AnswerTypeID = ?", answerTypeID)
	}

	// Fetch the filtered questions from the database
	result := query.Limit(intLimit).Offset(offset).Find(&filteredQuestions)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  fiber.StatusBadGateway,
			"message": result.Error,
		})
	}

	// Randomize the order of the filtered questions
	seed := time.Now().UnixNano()
	randomizer := rand.New(rand.NewSource(seed))
	for i := 0; i < len(filteredQuestions); i++ {
		j := randomizer.Intn(len(filteredQuestions)-i) + i
		filteredQuestions[i], filteredQuestions[j] = filteredQuestions[j], filteredQuestions[i]
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success",
		"results":   len(filteredQuestions),
		"questions": filteredQuestions})
}
