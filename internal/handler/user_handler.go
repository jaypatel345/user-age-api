package handler

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	db "github.com/jaypatel/user-age-api/db/sqlc"
	"github.com/jaypatel/user-age-api/internal/logger"
	"github.com/jaypatel/user-age-api/internal/service"
	"go.uber.org/zap"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

type createUserRequest struct {
	Name string `json:"name"`
	DOB  string `json:"dob"`
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req createUserRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Logger.Warn("Invalid create user request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		logger.Logger.Warn("Invalid create user name")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "name is required"})
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		logger.Logger.Warn("Invalid create user DOB", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "dob must be in YYYY-MM-DD format"})
	}

	user, err := h.service.CreateUser(c.UserContext(), db.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		logger.Logger.Error("Failed to create user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Logger.Info("User created", zap.Int32("id", user.ID))
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.service.ListUsers(c.UserContext())
	if err != nil {
		logger.Logger.Error("Failed to list users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Logger.Info("Users listed", zap.Int("count", len(users)))
	return c.JSON(users)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		logger.Logger.Warn("Invalid user id", zap.String("id", c.Params("id")), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	user, err := h.service.GetUserByID(c.UserContext(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Logger.Warn("User not found", zap.Int32("id", int32(id)))
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
		}

		logger.Logger.Error("Failed to get user", zap.Int32("id", int32(id)), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Logger.Info("User fetched", zap.Int32("id", user.ID))
	return c.JSON(user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		logger.Logger.Warn("Invalid user id", zap.String("id", c.Params("id")), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	var req createUserRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Logger.Warn("Invalid update user request body", zap.Int32("id", int32(id)), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		logger.Logger.Warn("Invalid update user name", zap.Int32("id", int32(id)))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "name is required"})
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		logger.Logger.Warn("Invalid update user DOB", zap.Int32("id", int32(id)), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "dob must be in YYYY-MM-DD format"})
	}

	user, err := h.service.UpdateUser(c.UserContext(), db.UpdateUserParams{
		ID:   int32(id),
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Logger.Warn("User not found", zap.Int32("id", int32(id)))
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
		}

		logger.Logger.Error("Failed to update user", zap.Int32("id", int32(id)), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Logger.Info("User updated", zap.Int32("id", user.ID))
	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		logger.Logger.Warn("Invalid user id", zap.String("id", c.Params("id")), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	if _, err := h.service.GetUserByID(c.UserContext(), int32(id)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Logger.Warn("User not found", zap.Int32("id", int32(id)))
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
		}

		logger.Logger.Error("Failed to get user before delete", zap.Int32("id", int32(id)), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.DeleteUser(c.UserContext(), int32(id)); err != nil {
		logger.Logger.Error("Failed to delete user", zap.Int32("id", int32(id)), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Logger.Info("User deleted", zap.Int32("id", int32(id)))
	return c.SendStatus(fiber.StatusNoContent)
}
