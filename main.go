package main

import (
	"log"

	"github.com/JoelOvien/cbt-backend/controllers"
	"github.com/JoelOvien/cbt-backend/database"
	"github.com/JoelOvien/cbt-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	// AuthController export
	AuthController controllers.AuthController
	// AuthRouteController export
	AuthRouteController routes.AuthRouteController

	// CollegeController export
	CollegeController controllers.CollegeController
	// CollegeRouteController export
	CollegeRouteController routes.CollegeRouteController

	// DepartmentController export
	DepartmentController controllers.DepartmentController
	// DepartmentRouteController export
	DepartmentRouteController routes.DepartmentRouteController

	// CourseController export
	CourseController controllers.CourseController
	// CourseRouteController export
	CourseRouteController routes.CourseRouteController

	// RegisteredCourseController export
	RegisteredCourseController controllers.RegisteredCourseController
	// RegisteredCourseRouteController export
	RegisteredCourseRouteController routes.RegisteredCourseRouteController

	// RoleController export
	RoleController controllers.RoleController
	// RoleRouteController export
	RoleRouteController routes.RoleRouteController

	// UserController export
	UserController controllers.UserController
	// UserRouteController export
	UserRouteController routes.UserRouteController
)

func init() {
	config, err := database.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load config \n", err.Error())
	} else {
		log.Println(" 🚀 Config loaded successfully")
	}

	database.ConnectToDB(&config)

	AuthController = controllers.NewAuthController(database.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	CollegeController = controllers.NewCollegeController(database.DB)
	CollegeRouteController = routes.NewCollegeRouteController(CollegeController)

	DepartmentController = controllers.NewDepartmentController(database.DB)
	DepartmentRouteController = routes.NewDepartmentRouteController(DepartmentController)

	CourseController = controllers.NewCourseController(database.DB)
	CourseRouteController = routes.NewCourseRouteController(CourseController)

	RegisteredCourseController = controllers.NewRegisteredCourseController(database.DB)
	RegisteredCourseRouteController = routes.NewRegisteredCourseRouteController(RegisteredCourseController)

	RoleController = controllers.NewRoleController(database.DB)
	RoleRouteController = routes.NewRoleRouteController(RoleController)

	UserController = controllers.NewUserController(database.DB)
	UserRouteController = routes.NewUserRouteController(UserController)

}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE, PUT",
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to my CBT project",
		})
	})

	AuthRouteController.AuthRoute(micro)
	CollegeRouteController.CollegeRoute(micro)
	DepartmentRouteController.DepartmentRoute(micro)
	CourseRouteController.CourseRoute(micro)
	RegisteredCourseRouteController.RegisteredCourseRoute(micro)
	RoleRouteController.RoleRoute(micro)
	UserRouteController.UserRoute(micro)

	log.Fatal(app.Listen(":8000"))
}
