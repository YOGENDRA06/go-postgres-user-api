package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"Go_Backend_Development_Task/internal/logger"
	"Go_Backend_Development_Task/internal/models"
	"Go_Backend_Development_Task/internal/service"
)

type UserHandler struct {
	service  *service.UserService
	validate *validator.Validate
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service:  s,
		validate: validator.New(),
	}
}

// ---------------- CREATE USER ----------------
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		logger.Log.Warn("invalid request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).
			JSON(models.NewError("INVALID_REQUEST", "invalid request body"))
	}

	if err := h.validate.Struct(req); err != nil {
		logger.Log.Warn("validation failed", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).
			JSON(models.NewError("VALIDATION_ERROR", err.Error()))
	}

	res, err := h.service.CreateUser(c.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidDOB) {
			return c.Status(fiber.StatusBadRequest).
				JSON(models.NewError(
					"VALIDATION_ERROR",
					"dob must be in YYYY-MM-DD format",
				))
		}

		logger.Log.Error("failed to create user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).
			JSON(models.NewError("INTERNAL_ERROR", "failed to create user"))
	}

	logger.Log.Info("user created", zap.Int64("user_id", res.ID))
	return c.Status(fiber.StatusCreated).JSON(res)
}

// ---------------- GET USER BY ID ----------------
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(models.NewError("INVALID_ID", "invalid user id"))
	}

	res, err := h.service.GetUser(c.Context(), id)
	if err != nil {
		logger.Log.Warn("user not found", zap.Int64("user_id", id))
		return c.Status(fiber.StatusNotFound).
			JSON(models.NewError("NOT_FOUND", "user not found"))
	}

	return c.JSON(res)
}

// ---------------- GET USERS (PAGINATED) ----------------
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	users, err := h.service.GetUsers(c.Context(), page, limit)
	if err != nil {
		logger.Log.Error("failed to fetch users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).
			JSON(models.NewError("INTERNAL_ERROR", "failed to fetch users"))
	}

	return c.JSON(fiber.Map{
		"page":  page,
		"limit": limit,
		"data":  users,
	})
}

// ---------------- UPDATE USER ----------------
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(models.NewError("INVALID_ID", "invalid user id"))
	}

	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(models.NewError("INVALID_REQUEST", "invalid request body"))
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(models.NewError("VALIDATION_ERROR", err.Error()))
	}

	res, err := h.service.UpdateUser(c.Context(), id, req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidDOB) {
			return c.Status(fiber.StatusBadRequest).
				JSON(models.NewError(
					"VALIDATION_ERROR",
					"dob must be in YYYY-MM-DD format",
				))
		}

		logger.Log.Error("failed to update user", zap.Int64("user_id", id), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).
			JSON(models.NewError("INTERNAL_ERROR", "failed to update user"))
	}

	logger.Log.Info("user updated", zap.Int64("user_id", id))
	return c.JSON(res)
}

// ---------------- DELETE USER ----------------
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(models.NewError("INVALID_ID", "invalid user id"))
	}

	if err := h.service.DeleteUser(c.Context(), id); err != nil {
		logger.Log.Error("failed to delete user", zap.Int64("user_id", id), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).
			JSON(models.NewError("INTERNAL_ERROR", "failed to delete user"))
	}

	logger.Log.Info("user deleted", zap.Int64("user_id", id))
	return c.SendStatus(fiber.StatusNoContent)
}
