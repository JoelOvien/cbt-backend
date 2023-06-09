package routes

import (
	"github.com/JoelOvien/cbt-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// DepartmentRouteController function for all auth routes
type DepartmentRouteController struct {
	departmentRouteController controllers.DepartmentController
}

// NewDepartmentRouteController function for all auth routes
func NewDepartmentRouteController(departmentController controllers.DepartmentController) DepartmentRouteController {
	return DepartmentRouteController{departmentController}
}

// DepartmentRoute defines auth routes for admin login
func (rc *DepartmentRouteController) DepartmentRoute(micro fiber.Router) {
	micro.Route("/department", func(router fiber.Router) {
		router.Get("/departments", rc.departmentRouteController.FetchAllDepartments)
		router.Post("", rc.departmentRouteController.CreateDepartment)
	})

}
