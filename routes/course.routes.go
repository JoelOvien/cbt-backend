package routes

import (
	"github.com/JoelOvien/cbt-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// CourseRouteController function for all auth routes
type CourseRouteController struct {
	courseRouteController controllers.CourseController
}

// NewCourseRouteController function for all auth routes
func NewCourseRouteController(courseController controllers.CourseController) CourseRouteController {
	return CourseRouteController{courseController}
}

// CourseRoute defines auth routes for admin login
func (rc *CourseRouteController) CourseRoute(micro fiber.Router) {
	micro.Route("/courses", func(router fiber.Router) {
		router.Get("", rc.courseRouteController.FetchAllCourses)
		router.Post("", rc.courseRouteController.CreateCourse)
	})

}
