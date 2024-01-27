package internal

import (
	"main/model"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service actions
}

type actions interface {
	PostBinaryTreeSum(tree *model.Tree) (*model.BinaryTreeSum, error)
}

func NewHandler(service actions) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/postBinaryTreeSum", h.PostBinaryTreeSum)
}

func (h *Handler) PostBinaryTreeSum(ctx *fiber.Ctx) error {
	var tree model.Tree

	err := ctx.BodyParser(&tree)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	binaryTreeSum, err := h.service.PostBinaryTreeSum(&tree)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.Status(fiber.StatusOK).JSON(binaryTreeSum)
}
