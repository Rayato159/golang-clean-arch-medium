package middlewares

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		if accessToken == "" {
			log.Println("error, authorization header is empty.")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":      "Unauthorized",
				"status_code": fiber.StatusUnauthorized,
				"message":     "error, unauthorized",
				"result":      nil,
			})
		}

		secretKey := os.Getenv("JWT_SECRET_KEY")
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("error, unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(secretKey), nil
		})
		if err != nil {
			log.Println(err.Error())
			return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
				"status":      fiber.ErrUnauthorized.Message,
				"status_code": fiber.ErrUnauthorized.Code,
				"message":     "error, unauthorized",
				"result":      nil,
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			c.Locals("user_id", claims["user_id"])
			c.Locals("username", claims["username"])
			return c.Next()
		}
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"status":      fiber.ErrUnauthorized.Message,
			"status_code": fiber.ErrUnauthorized.Code,
			"message":     "error, unauthorized",
			"result":      nil,
		})
	}
}
