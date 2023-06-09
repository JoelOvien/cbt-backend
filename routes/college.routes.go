package routes

import (
	"github.com/JoelOvien/cbt-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// CollegeRouteController function for all auth routes
type CollegeRouteController struct {
	collegeRouteController controllers.CollegeController
}

// NewCollegeRouteController function for all auth routes
func NewCollegeRouteController(collegeController controllers.CollegeController) CollegeRouteController {
	return CollegeRouteController{collegeController}
}

// CollegeRoute defines auth routes for admin login
func (rc *CollegeRouteController) CollegeRoute(micro fiber.Router) {
	micro.Route("/college", func(router fiber.Router) {
		router.Get("/colleges", rc.collegeRouteController.FetchAllColleges)
		router.Post("", rc.collegeRouteController.CreateCollege)
	})

}
