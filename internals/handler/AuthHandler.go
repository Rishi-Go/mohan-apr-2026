package handler

import (
	"blog_post/internals/dto"
	"blog_post/internals/service"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v3"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	Service service.AuthService
}

func InitAuthHandler(svc service.AuthService) *AuthHandler {
	return &AuthHandler{Service: svc}
}

var validate = validator.New()

func (h *AuthHandler) SignUpUser(Ctx fiber.Ctx) error {

	var res = dto.SignUpRequest{}

	if err := Ctx.Bind().Body(&res); err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	if err := validate.Struct(&res); err != nil {

		errs := err.(validator.ValidationErrors)
		errMap := make(map[string]string)

		for _, e := range errs {

			var field string
			switch e.StructField() {
			case "FirstName":
				field = "first_name"
			case "LastName":
				field = "last_name"
			case "Password":
				field = "password"
			case "Email":
				field = "email"
			case "UserName":
				field = "user_name"
			}

			switch e.Tag() {
			case "required":
				errMap[field] = "field is required"
			case "email":
				errMap[field] = "must be a valid email address"
			case "min":
				errMap[field] = "must be at least " + e.Param() + " characters long"
			default:
				errMap[field] = "is invalid"
			}
		}
		return Ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": "Validation failed", "Details": errMap, "StatusCode": http.StatusBadRequest})
	}

	result, err := h.Service.SignUpUser(res)
	if err != nil {
		// logger.Log.Error("Validation failed", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": "Validation failed", "Details": err.Error(), "StatusCode": http.StatusBadRequest})
	}

	err = Ctx.Status(fiber.StatusCreated).JSON(&dto.SuccessResponse{
		Message:    "User registered successfully",
		StatusCode: fiber.StatusCreated,
		Data: &dto.BlogUserResponse{
			BlogUsers: result,
		},
	})
	// logger.Log.Info("User registered successfully")
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	return nil
}

func (h *AuthHandler) GetUser(Ctx fiber.Ctx) error {

	username := Ctx.Query("username")
	email := Ctx.Query("email")

	page, err := strconv.Atoi(Ctx.Query("page"))

	if page < 1 {
		page = 1
	} else if err != nil {
		// logger.Log.With(zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	limit, err := strconv.Atoi(Ctx.Query("limit"))

	if limit > 99 || limit < 0 {
		// logger.Log.Debug("Limit should be with in 1 - 99")
		return Ctx.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{Message: "Limit should be with in 1 - 99", StatusCode: http.StatusInternalServerError})
	} else if limit == 0 {

		limit = 10

	} else if err != nil {
		// logger.Log.With(zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	offset := (page - 1) * limit

	result, Page, err := h.Service.GetUser(page, limit, offset, username, email)
	if err != nil {
		// logger.Log.Error("failed to get user details", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: &dto.BlogUsersResponse{
			BlogUsers: result,
		}})

	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "User details retreived successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.UserResponse{
			BlogUsers:  result,
			Pagination: *Page,
		},
	})
	// logger.Log.Info("User details retreived successfully")

	if err != nil {
		// logger.Log.Warn("User Details Not found", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil

}

func (h *AuthHandler) SelectUser(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	UserId, err := uuid.FromString(uuidStr)
	if err != nil {
		// logger.Log.With(zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	ID, err := h.Service.SelectUser(UserId)
	if err != nil {
		// logger.Log.With(zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorMessage{Message: err.Error(), StatusCode: http.StatusNotFound, ID: UserId})
	}

	err = Ctx.Status(http.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "User detail retreived successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.BlogUserResponse{
			BlogUsers: ID,
		}})

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *AuthHandler) UpdateUser(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	UserId, err := uuid.FromString(uuidStr)
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	var res = dto.SignUpRequest{}

	claims := Ctx.Locals("user_id").(jwt.MapClaims)

	userID := claims["id"].(string)

	LogInUser, err := uuid.FromString(userID)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	role := claims["role"].(string)

	if err := Ctx.Bind().Body(&res); err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = h.Service.UpdateUser(res, UserId, LogInUser, role)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: dto.Response{Message: "Failed to Update Record", ID: UserId}})
	}

	err = Ctx.Status(http.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "User Updated successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.Response{
			Message: "Successfully Updated Record", ID: UserId,
		}})

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *AuthHandler) DeleteUser(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	claims := Ctx.Locals("user_id").(jwt.MapClaims)

	userID := claims["id"].(string)
	LogInUser, err := uuid.FromString(userID)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	role := claims["role"].(string)

	UserId, err := uuid.FromString(uuidStr)
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	err = h.Service.DeleteUser(UserId, LogInUser, role)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: dto.Response{Message: "Failed to Delete Record", ID: UserId}})
	}
	err = Ctx.Status(http.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "User Deleted successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.Response{
			Message: "Successfully deleted Record", ID: UserId,
		}})
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *AuthHandler) LogInUser(Ctx fiber.Ctx) error {

	var res dto.LogInRequest

	if err := Ctx.Bind().Body(&res); err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	tokenString, err := h.Service.LogInUser(res)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	Ctx.Cookie(&fiber.Cookie{
		Expires:  time.Now().Add(24 * time.Hour),
		Name:     "auth_token",
		Value:    tokenString,
		HTTPOnly: true,
		Secure:   false,
		Path:     "/",
	})

	err = Ctx.Status(http.StatusOK).JSON(dto.TokenMessage{Message: "Logged In Successfully", Token: tokenString, StatusCode: http.StatusOK})
	if err != nil {
		return Ctx.Status(http.StatusUnauthorized).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusUnauthorized})
	}

	return nil
}
