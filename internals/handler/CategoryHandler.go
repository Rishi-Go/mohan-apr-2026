package handler

import (
	"blog_post/internals/dto"
	"blog_post/internals/service"
	"blog_post/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type CategoryHandler struct {
	Service service.CategoryService
}

func InitCategoryHandler(svc service.CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: svc}
}

func (h *CategoryHandler) InsertCategory(Ctx fiber.Ctx) error {

	var res = dto.CategoryRequest{}

	if err := Ctx.Bind().Body(&res); err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	result, err := h.Service.InsertCategory(res)
	if err != nil {
		logger.Log.Error("Failed to Insert Category details", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Category register successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.CategoryInsertResponse{
			Category: result,
		},
	})
	logger.Log.Info("Category register successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *CategoryHandler) GetCategory(Ctx fiber.Ctx) error {

	category_name := Ctx.Query("search")

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

	result, Page, err := h.Service.GetCategory(page, limit, offset, category_name)
	if err != nil {
		logger.Log.Error("Failed to retreived Category details", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: &dto.CategoryInsertsResponse{
			Category: result,
		}})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Category details retreived successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.CategoryResponse{
			Category:   result,
			Pagination: *Page,
		},
	})
	logger.Log.Info("Category details retreived successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *CategoryHandler) SelectCategory(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	CategoryId, err := uuid.FromString(uuidStr)
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	ID, err := h.Service.SelectCategory(CategoryId)
	if err != nil {
		logger.Log.Error("Failed to retreived Category detail", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorMessage{Message: err.Error(), StatusCode: http.StatusNotFound, ID: CategoryId})
	}

	err = Ctx.Status(fiber.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Category detail retreived successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.CategoryInsertResponse{
			Category: ID,
		},
	})
	logger.Log.Info("Category detail retreived successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *CategoryHandler) UpdateCategory(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	CategoryId, err := uuid.FromString(uuidStr)
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	var res = dto.CategoryRequest{}

	if err := Ctx.Bind().Body(&res); err != nil {
		logger.Log.With(zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	}

	err = h.Service.UpdateCategory(res, CategoryId)
	if err != nil {
		logger.Log.Error("Failed to Update Category detail", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: dto.Response{Message: "Failed to Update Category Record", ID: CategoryId}})
	}

	err = Ctx.Status(http.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Category Updated successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.Response{
			Message: "Successfully Updated Record", ID: CategoryId,
		}})
	logger.Log.Info("Category Updated successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}

func (h *CategoryHandler) DeleteCategory(Ctx fiber.Ctx) error {

	uuidStr := Ctx.Params("id")

	CategoryId, err := uuid.FromString(uuidStr)
	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}

	err = h.Service.DeleteCategory(CategoryId)
	if err != nil {
		logger.Log.Error("Failed to Delete Category detail", zap.String("Error:", err.Error()))
		return Ctx.Status(http.StatusBadRequest).JSON(dto.SuccessResponse{Message: err.Error(), StatusCode: http.StatusBadRequest, Data: dto.Response{Message: "Failed to Delete Category Record", ID: CategoryId}})
	}

	err = Ctx.Status(http.StatusOK).JSON(&dto.SuccessResponse{
		Message:    "Category Delete successfully",
		StatusCode: fiber.StatusOK,
		Data: &dto.Response{
			Message: "Successfully Delete Record", ID: CategoryId,
		}})
	logger.Log.Info("Category Delete successfully")

	if err != nil {
		return Ctx.Status(http.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	}
	return nil
}
