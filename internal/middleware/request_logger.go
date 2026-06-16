package middleware

import (
	"time"

	"ainyx/internal/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestLogger(c *fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	duration := time.Since(start)

	logger.Log.Info(
		"request completed",
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
		zap.Int("status", c.Response().StatusCode()),
		zap.Duration("duration", duration),
	)

	return err
}