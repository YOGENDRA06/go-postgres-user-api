package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"Go_Backend_Development_Task/internal/logger"
)

const RequestIDKey = "requestId"

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := uuid.New().String()

		// Set response header
		c.Set("X-Request-Id", requestID)

		// Store in context locals
		c.Locals(RequestIDKey, requestID)

		// Log request id
		logger.Log.Info("request received",
			zap.String("requestId", requestID),
		)

		return c.Next()
	}
}
