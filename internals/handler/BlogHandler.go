package handler

import (
	"blog_post/internals/dto"
	"blog_post/internals/service"
	"blog_post/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type BlogHandler struct {
	Service service.BlogService
}

func InitBlogHandler(svc service.BlogService) *BlogHandler {
	return &BlogHandler{Service: svc}
}

func (h *BlogHandler) InsertBlog(Ctx fiber.Ctx) error {

	var res = dto.BlogRequest{}

	claims := Ctx.Locals("user_id").(jwt.MapClaims)

	userID := claims["id"].(string)

	ID, err := uuid.FromString(userID)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	res.AuthorID = ID

	if err := Ctx.Bind().Body(&res); err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	result, err := h.Service.InsertBlog(res)
	if err != nil {
		logger.Log.Error("Failed to Post Blog", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Blog Posted successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.BlogPostResponse{
			Blog: result,
		},
	})
	logger.Log.Info("Blog Posted successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil

}

func (h *BlogHandler) GetBlog(Ctx fiber.Ctx) error {

	title := Ctx.Query("title")

	categoryStr := Ctx.Query("category-id")

	categoryId := uuid.FromStringOrNil(categoryStr)

	AuthorStr := Ctx.Query("author-id")

	authorId := uuid.FromStringOrNil(AuthorStr)

	page, err := strconv.Atoi(Ctx.Query("page"))

	if page < 1 {
		page = 1
	} else if err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	limit, err := strconv.Atoi(Ctx.Query("limit"))

	if limit > 99 || limit < 0 {

		return Ctx.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{Message: "Limit should be with in 1 - 99", StatusCode: http.StatusInternalServerError})
	} else if limit == 0 {

		limit = 10

	} else if err != nil {

		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	offset := (page - 1) * limit

	result, Page, err := h.Service.GetBlog(page, limit, offset, title, categoryId, authorId)
	if err != nil {
		logger.Log.Error("Failed to retreive Blog details", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: dto.BlogPostResponses{Blog: result}})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Blog details retreived successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.BlogResponse{
			Blog:       result,
			Pagination: *Page,
		},
	})
	logger.Log.Info("Blog details retreived successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil

}

func (h *BlogHandler) SelectBlog(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	BlogId, err := uuid.FromString(uuidStr)

	if err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	ID, err := h.Service.SelectBlog(BlogId)
	if err != nil {
		logger.Log.Error("Failed to retreive Blog detail", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorMessage{Message: err.Error(), StatusCode: http.StatusNotFound, ID: BlogId})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Blog detail retreived successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.BlogPostResponse{
			Blog: ID,
		},
	})
	logger.Log.Info("Blog detail retreived successfully")
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *BlogHandler) UpdateBlog(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	BlogId, err := uuid.FromString(uuidStr)
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	var res = dto.BlogRequest{}

	claims := Ctx.Locals("user_id").(jwt.MapClaims)

	userID := claims["id"].(string)
	role := claims["role"].(string)

	ID, err := uuid.FromString(userID)
	if err != nil {
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}
	res.AuthorID = ID

	if err := Ctx.Bind().Body(&res); err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = h.Service.UpdateBlog(res, BlogId, role)
	if err != nil {
		logger.Log.Error("Failed to Update Blog Record", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: dto.Response{Message: "Failed to Update Blog Record", ID: BlogId}})
	}

	err = Ctx.Status(http.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Blog Post Updated successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.Response{
			Message: "Successfully Updated Record", ID: BlogId,
		}})
	logger.Log.Info("Blog Post Updated successfully")
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *BlogHandler) DeleteBlog(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	BlogId, err := uuid.FromString(uuidStr)
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	claims := Ctx.Locals("user_id").(jwt.MapClaims)

	userID := claims["id"].(string)

	role := claims["role"].(string)

	TokenAuthorID, err := uuid.FromString(userID)
	if err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = h.Service.DeleteBlog(BlogId, TokenAuthorID, role)
	if err != nil {
		logger.Log.Error("Failed to Delete Blog Record",zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: dto.Response{Message: "Failed to Delete Blog Record", ID: BlogId}})
	}

	err = Ctx.Status(http.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Blog Post Deleted successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.Response{
			Message: "Successfully Deleted Record", ID: BlogId,
		}})
	logger.Log.Info("Blog Post Deleted successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}
