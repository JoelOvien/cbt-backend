package routes

import (
	"github.com/JoelOvien/cbt-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// RegisteredCourseRouteController function for all auth routes
type RegisteredCourseRouteController struct {
	registeredCourseRouteController controllers.RegisteredCourseController
}

// NewRegisteredCourseRouteController function for all auth routes
func NewRegisteredCourseRouteController(registeredCourseController controllers.RegisteredCourseController) RegisteredCourseRouteController {
	return RegisteredCourseRouteController{registeredCourseController}
}

// RegisteredCourseRoute defines auth routes for admin login
func (rc *RegisteredCourseRouteController) RegisteredCourseRoute(micro fiber.Router) {
	micro.Route("/registered/course", func(router fiber.Router) {
		router.Get("/courses", rc.registeredCourseRouteController.FetchAllRegisteredCourses)

		router.Post("/register", rc.registeredCourseRouteController.RegisterCourseForStudent)
	})

	micro.Route("/registered/course/:UserID", func(router fiber.Router) {
		router.Get("", rc.registeredCourseRouteController.FetchStudentRegisteredCourses)
	})

	micro.Route("/registered/course/:UserID/:Semester", func(router fiber.Router) {
		router.Get("", rc.registeredCourseRouteController.FetchStudentRegisteredCoursesBySemester)
	})

	micro.Route("/registered/course", func(router fiber.Router) {
		router.Get("", rc.registeredCourseRouteController.FetchStudentRegisteredForCourseByCourseCode)
	})

}
