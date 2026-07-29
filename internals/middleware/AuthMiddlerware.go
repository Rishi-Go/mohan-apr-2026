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
		"exp":  time.Now().Add(time.Hour * 24 * 15).Unix(),
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
		return Ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Token"})
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

	Ctx.Next()

	return nil
}

func AdminOnly(c fiber.Ctx) error {

	role := c.Cookies("role")
	
	if role != "Admin"{
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"Message": "Access Denied"})
	}
	return c.Next()
}
