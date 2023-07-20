package routes

import (
	"github.com/JoelOvien/cbt-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

// ExamBankRouteController struct
type ExamBankRouteController struct {
	examRouteController controllers.ExamBankController
}

// NewExamBankRouteController function
func NewExamBankRouteController(examBankController controllers.ExamBankController) ExamBankRouteController {
	return ExamBankRouteController{examBankController}
}

// ExamBankRoute defines routes
func (qc *ExamBankRouteController) ExamBankRoute(micro fiber.Router) {
	micro.Route("/exam-bank", func(router fiber.Router) {
		router.Post("/", qc.examRouteController.SubmitAnswer)
		router.Post("/single", qc.examRouteController.SubmitSingleAnswer)
		router.Get("", qc.examRouteController.FindAll)
	})
}
