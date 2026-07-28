package middleware

import (
	"blog_post/internals/dto"
	"blog_post/pkg/models"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(res dto.LogInRequest, users models.BlogUsers) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": users.ID,
		"exp": time.Now().Add(time.Hour * 24 * 15).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", errors.New("Failed to create Token")
	}

	return tokenString, nil

}

func RequestToken(Ctx fiber.Ctx) error {

	cookie := new(fiber.Cookie)
	tokenString := cookie.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return os.Getenv("SECRET_KEY"), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		log.Fatal(err)
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["sub"], claims["exp"])
	} else {
		fmt.Println(err)
		return err
	}

	Ctx.Next()
	return nil
}
