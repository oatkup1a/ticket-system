package handler

import (
	"github.com/gofiber/fiber/v2"
	"oatkup.sys/internal/application/usecase/applications"
)

type ApplicationHandler struct {
	createUC *applications.CreateApplicationUsecase
}

func NewApplicationHandler(createUC *applications.CreateApplicationUsecase) *ApplicationHandler {
	return &ApplicationHandler{createUC: createUC}
}

func (h *ApplicationHandler) Create(c *fiber.Ctx) error {
	var in applications.ApplicationRequest
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	out, err := h.createUC.Apply(c.Context(), &in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(out)
}
