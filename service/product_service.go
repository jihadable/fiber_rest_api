package service

import (
	"context"
	"database/sql"
	"fiber_rest_api/model"
	"fiber_rest_api/repo"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductService interface {
	GetProduct(ctx *fiber.Ctx) error
	GetProducts(ctx *fiber.Ctx) error
	StoreProduct(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
	DeleteProduct(ctx *fiber.Ctx) error
}

type productServiceImpl struct {
	Repo repo.ProductRepo
}

func NewProductService(db *sql.DB) ProductService {
	return &productServiceImpl{Repo: repo.NewProductRepo(db)}
}

func (service *productServiceImpl) GetProduct(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	productId, err := strconv.Atoi(paramId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid product ID",
		})
	}

	product, err := service.Repo.FindById(context.Background(), productId)
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

func (service *productServiceImpl) GetProducts(ctx *fiber.Ctx) error {
	products, err := service.Repo.FindAll(context.Background())
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

func (service *productServiceImpl) StoreProduct(ctx *fiber.Ctx) error {
	product := model.Product{}
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	newProduct, err := service.Repo.Save(context.Background(), product)
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

func (service *productServiceImpl) UpdateProduct(ctx *fiber.Ctx) error {
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

	updatedProduct, err := service.Repo.Update(context.Background(), product)
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

func (service *productServiceImpl) DeleteProduct(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")
	productId, err := strconv.Atoi(paramId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid product ID",
		})
	}

	if err = service.Repo.Delete(context.Background(), productId); err != nil {
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
