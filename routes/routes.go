package routes

import (
	"belajar_fiber/config"
	"belajar_fiber/handlers"
	"belajar_fiber/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	produkService := &services.ProductService{DB: db}
	kategoriProductService := &services.KategoriProductService{DB: db}

	productHandler := &handlers.ProductHandler{ProductService: produkService}
	kategoriProductHandler := &handlers.KategoriProductHandler{KategoriProductService: kategoriProductService}

	api := app.Group("/api")
	api.Static("/public", "/public")

	//Produk Route
	productRoutes := api.Group("/produk")
	productRoutes.Get("/", productHandler.GetAllProducts)
	productRoutes.Get("/:uuid", productHandler.GetProductByUID)
	productRoutes.Post("/", productHandler.CreateProduct)
	productRoutes.Put("/:uuid", productHandler.UpdateProduct)
	productRoutes.Delete("/:uuid", productHandler.DeleteProduct)

	//Kategori Produk Route
	kategoriProductRoutes := api.Group("/kategori_produk")
	kategoriProductRoutes.Get("/headers", kategoriProductHandler.GetAllRequest)
	kategoriProductRoutes.Get("/", kategoriProductHandler.GetAllCategory)
	kategoriProductRoutes.Get("/:uuid", kategoriProductHandler.GetKategoriCategoryByUUID)
	kategoriProductRoutes.Post("/", kategoriProductHandler.CreateCategory)
	kategoriProductRoutes.Put("/:uuid", kategoriProductHandler.UpdateCategory)
	kategoriProductRoutes.Delete("/:uuid", kategoriProductHandler.DeleteCategory)

}