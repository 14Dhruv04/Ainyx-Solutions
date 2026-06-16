package handler

import (
	"ainyx/internal/logger"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"

	"ainyx/internal/middleware"
	"ainyx/internal/models"
	"ainyx/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(
	service *service.UserService,
) *UserHandler {

	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(
	c *fiber.Ctx,
) error {

	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid request",
			},
		)
	}

	if err := middleware.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	user, err := h.service.CreateUser(
		c.Context(),
		req.Name,
		req.DOB,
	)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	logger.Log.Info(
		"user created",
		zap.String("name", user.Name),
	)

	return c.Status(201).JSON(
		fiber.Map{
			"id":   user.ID,
			"name": user.Name,
			"dob":  user.Dob.Format("2006-01-02"),
		},
	)
}

func (h *UserHandler) GetUser(
	c *fiber.Ctx,
) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid id",
			},
		)
	}

	user, err := h.service.GetUser(
		c.Context(),
		int32(id),
	)

	if err == sql.ErrNoRows {
		return c.Status(404).JSON(
			fiber.Map{
				"error": "user not found",
			},
		)
	}

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(user)
}

func (h *UserHandler) GetAllUsers(
	c *fiber.Ctx,
) error {

	users, err := h.service.GetAllUsers(
		c.Context(),
	)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(users)
}

func (h *UserHandler) UpdateUser(
	c *fiber.Ctx,
) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid id",
			},
		)
	}

	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid request",
			},
		)
	}

	if err := middleware.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	user, err := h.service.UpdateUser(
		c.Context(),
		int32(id),
		req.Name,
		req.DOB,
	)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	logger.Log.Info(
		"user updated",
		zap.Int32("id", user.ID),
	)

	return c.JSON(
		fiber.Map{
			"id":   user.ID,
			"name": user.Name,
			"dob":  user.Dob.Format("2006-01-02"),
		},
	)
}

func (h *UserHandler) DeleteUser(
	c *fiber.Ctx,
) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid id",
			},
		)
	}

	err = h.service.DeleteUser(
		c.Context(),
		int32(id),
	)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	logger.Log.Info(
		"user deleted",
		zap.Int32("id", int32(id)),
	)

	return c.SendStatus(fiber.StatusNoContent)
}
