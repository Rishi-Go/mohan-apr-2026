package middleware

import (
	"blog_post/internals/dto"
	"blog_post/pkg/models"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(res dto.LogInRequest, users models.BlogUsers) (string, error) {

	claims := jwt.MapClaims{
		"sub":  users.ID,
		"role": users.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", errors.New("Failed to create Token")
	}

	return tokenString, nil
}

func VerifyToken(Ctx fiber.Ctx) error {

	tokenString := Ctx.Cookies("jwt_token")

	if tokenString == "" {
		return Ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authentication Token"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil

	})

	if err != nil || !token.Valid {
		return Ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err})
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return Ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token Expired"})
	}
	return Ctx.Next()
}

func RoleAuthorizeMiddleware(allowedRoles ...string) fiber.Handler {
	return func(Ctx fiber.Ctx) error {

		TokenString := Ctx.Cookies("jwt_token")

		token, _ := jwt.Parse(TokenString, func(token *jwt.Token) (interface{}, error) {

			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return Ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error": "Invalid token claims"})
		}

		userRole, exists := claims["role"].(string)
		if !exists {
			return Ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Error": "Role claim not found in token"})
		}

		if userRole == "" {
			return Ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error": "Unauthorized: No role cookie found", "StatusCode": fiber.StatusUnauthorized})
		}

		for _, role := range allowedRoles {
			if userRole == role {
				return Ctx.Next()
			}
		}

		return Ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"Role": userRole,"Error": "Forbidden: Insufficient permissions", "StatusCode": fiber.StatusForbidden})
	}
}
