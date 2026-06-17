package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jaypatel/user-age-api/internal/logger"
	"go.uber.org/zap"
)

const requestIDKey = "request_id"

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := uuid.New().String()

		c.Locals(requestIDKey, id)
		c.Set("X-Request-ID", id)

		return c.Next()
	}
}

func Logging() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		fields := []zap.Field{
			zap.String("request_id", requestID(c)),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", time.Since(start)),
			zap.String("ip", c.IP()),
		}

		if err != nil {
			fields = append(fields, zap.Error(err))
			logger.Logger.Error("request completed with error", fields...)
			return err
		}

		logger.Logger.Info("request completed", fields...)
		return nil
	}
}

func requestID(c *fiber.Ctx) string {
	id, ok := c.Locals(requestIDKey).(string)
	if !ok {
		return ""
	}

	return id
}
