package routes

import (
	"github.com/JoelOvien/cbt-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// AuthRouteController function for all auth routes
type AuthRouteController struct {
	authRouteController controllers.AuthController
}

// NewAuthRouteController function for all auth routes
func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

// AuthRoute defines auth routes for admin login
func (rc *AuthRouteController) AuthRoute(micro fiber.Router) {

	micro.Route("/auth", func(router fiber.Router) {
		router.Post("/register", rc.authRouteController.CreateUser)
		router.Post("/login", rc.authRouteController.SignInUser)
		router.Get("/refresh", rc.authRouteController.RefreshAccessToken)
	})

}
