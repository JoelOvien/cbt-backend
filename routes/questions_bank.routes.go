package routes

import (
	"github.com/JoelOvien/cbt-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

// QuestionsBankRouteController struct
type QuestionsBankRouteController struct {
	questionsBankRouteController controllers.QuestionsBankController
}

// NewQuestionsBankRouteController function
func NewQuestionsBankRouteController(questionsBankController controllers.QuestionsBankController) QuestionsBankRouteController {
	return QuestionsBankRouteController{questionsBankController}
}

// QuestionsBankRoute defines routes
func (qc *QuestionsBankRouteController) QuestionsBankRoute(micro fiber.Router) {
	micro.Route("/exam-bank", func(router fiber.Router) {
		router.Post("", qc.questionsBankRouteController.CreateQuestion)
	})

}
