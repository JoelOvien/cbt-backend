package routes

import (
	"github.com/JoelOvien/cbt-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// UserRouteController function for all auth routes
type UserRouteController struct {
	userRouteController controllers.UserController
}

// NewUserRouteController function for all auth routes
func NewUserRouteController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

// UserRoute defines auth routes for admin login
func (rc *UserRouteController) UserRoute(micro fiber.Router) {
	micro.Route("/users", func(router fiber.Router) {
		router.Get("", rc.userRouteController.FetchAllUsers)
	})

}
