package servers

import (
	_usersHttp "github.com/rayato159/clean-arch-medium/modules/users/controllers"
	_usersRepository "github.com/rayato159/clean-arch-medium/modules/users/repositories"
	_usersUsecase "github.com/rayato159/clean-arch-medium/modules/users/usecases"

	_authHttp "github.com/rayato159/clean-arch-medium/modules/auth/controllers"
	_authRepository "github.com/rayato159/clean-arch-medium/modules/auth/repositories"
	_authUsecase "github.com/rayato159/clean-arch-medium/modules/auth/usecases"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	// Group a version
	v1 := s.App.Group("/v1")

	//* Users group
	usersGroup := v1.Group("/users")
	usersRepository := _usersRepository.NewUsersRepository(s.Db)
	usersUsecase := _usersUsecase.NewUsersUsecase(usersRepository)
	_usersHttp.NewUsersController(usersGroup, usersUsecase)

	authGroup := v1.Group("/auth")
	authRepository := _authRepository.NewAuthRepository(s.Db)
	authUsecase := _authUsecase.NewAuthUsecase(authRepository, usersRepository)
	_authHttp.NewAuthController(authGroup, s.Cfg, authUsecase)

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil
}
