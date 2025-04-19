package controller

import (
	"context"
	"database/sql"
	"fiber_rest_api/model"
	"fiber_rest_api/repo"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type productController struct {
	Repo repo.ProductRepo
}

func NewProductController(db *sql.DB) *productController {
	return &productController{Repo: repo.NewProductRepo(db)}
}

func (controller *productController) GetProduct(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	productId, err := strconv.Atoi(paramId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid product ID",
		})
	}

	product, err := controller.Repo.FindById(context.Background(), productId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Product not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Product retrieved successfully",
		"data":    product,
	})
}

func (controller *productController) GetProducts(ctx *fiber.Ctx) error {
	products, err := controller.Repo.FindAll(context.Background())
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Products not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Products retrieved successfully",
		"data":    products,
	})
}

func (controller *productController) StoreProduct(ctx *fiber.Ctx) error {
	product := model.Product{}
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	newProduct, err := controller.Repo.Save(context.Background(), product)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Product created successfully",
		"data":    newProduct,
	})
}

func (controller *productController) UpdateProduct(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	productId, err := strconv.Atoi(paramId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid product ID",
		})
	}

	product := model.Product{}
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid body request",
		})
	}

	product.Id = productId

	updatedProduct, err := controller.Repo.Update(context.Background(), product)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Product not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Product updated successfully",
		"data":    updatedProduct,
	})
}

func (controller *productController) DeleteProduct(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	productId, err := strconv.Atoi(paramId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid product ID",
		})
	}

	if err = controller.Repo.Delete(context.Background(), productId); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Product not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Product deleted successfully",
	})
}
