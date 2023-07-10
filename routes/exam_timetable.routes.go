package routes

import (
	"github.com/JoelOvien/cbt-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// ExamTimetableRouteController function for all auth routes
type ExamTimetableRouteController struct {
	examTimetableRouteController controllers.ExamTimetableController
}

// NewExamTimetableRouteController function for all auth routes
func NewExamTimetableRouteController(examTimetableController controllers.ExamTimetableController) ExamTimetableRouteController {
	return ExamTimetableRouteController{examTimetableController}
}

// ExamTimetableRoute defines auth routes for admin login
func (ec *ExamTimetableRouteController) ExamTimetableRoute(micro fiber.Router) {
	micro.Route("/exam/timetable", func(router fiber.Router) {
		router.Get("", ec.examTimetableRouteController.FetchExamTimetable)
		router.Post("", ec.examTimetableRouteController.CreateExamTimetable)
	})

	micro.Route("/exam/timetable/:Semester", func(router fiber.Router) {
		router.Get("", ec.examTimetableRouteController.FetchExamTimetableBySemester)
	})

	micro.Route("/exam/timetable", func(router fiber.Router) {
		router.Get("/?CourseCode:CourseCode", ec.examTimetableRouteController.FetchExamTimetableByCourseCode)
	})

}
