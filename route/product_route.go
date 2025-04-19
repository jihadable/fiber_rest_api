package route

import (
	"database/sql"
	"fiber_rest_api/service"

	"github.com/gofiber/fiber/v2"
)

func ProductRoute(api fiber.Router, db *sql.DB) {
	productController := service.NewProductService(db)
	productRoute := api.Group("/products")

	productRoute.Get("/:id", productController.GetProduct)
	productRoute.Get("/", productController.GetProducts)
	productRoute.Post("/", productController.StoreProduct)
	productRoute.Patch("/:id", productController.UpdateProduct)
	productRoute.Delete("/:id", productController.DeleteProduct)
}
