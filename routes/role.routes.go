package routes

import (
	"github.com/JoelOvien/cbt-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// RoleRouteController function for all auth routes
type RoleRouteController struct {
	roleRouteController controllers.RoleController
}

// NewRoleRouteController function for all auth routes
func NewRoleRouteController(roleController controllers.RoleController) RoleRouteController {
	return RoleRouteController{roleController}
}

// RoleRoute defines auth routes for admin login
func (rc *RoleRouteController) RoleRoute(micro fiber.Router) {
	micro.Route("/roles", func(router fiber.Router) {
		router.Get("", rc.roleRouteController.FetchAllRoles)
		router.Post("", rc.roleRouteController.CreateRole)
	})

}
