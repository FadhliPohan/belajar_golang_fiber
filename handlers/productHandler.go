package handlers

import (
	"belajar_fiber/models"
	"belajar_fiber/response"
	"belajar_fiber/services"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	ProductService *services.ProductService
}

// GetAllProducts handles the request to fetch all products
func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.ProductService.GetAllProduct()
    if err!= nil {
        return response.InternalServerError(c, err, "Cannot retrieve products")
    }
    return response.SuccessHandler(c, fiber.StatusOK, "Products retrieved successfully", products)
}

// GetProductByUID handles the request to fetch a product by its UID

func (h *ProductHandler) GetProductByUID(c *fiber.Ctx) error {
    uid := c.Params("uuid")
    product, err := h.ProductService.GetProductByID(uid)
    if err != nil {
        return response.InternalServerError(c, err, "Cannot retrieve product")
    }
    if product == nil {
        return response.BadRequest(c, nil, "Product not found")
    }
    return response.SuccessHandler(c, fiber.StatusOK, "Product retrieved successfully", product)
}

// CreateProduct handles the request to create a new product

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
    var product models.Product
    if err := c.BodyParser(&product); err != nil {
        return response.BadRequest(c, err, "Cannot parse JSON")
    }
    if err := h.ProductService.CreateProduct(&product); err != nil {
        return response.InternalServerError(c, err, "Cannot create product")
    }
    return response.SuccessHandler(c, fiber.StatusCreated, "Product created successfully", product)
}

// UpdateProduct handles the request to update a product by its UID

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
    uid := c.Params("uuid")
    var updatedProduct models.Product
    if err := c.BodyParser(&updatedProduct); err != nil {
        return response.BadRequest(c, err, "Cannot parse JSON")
    }
    if err := h.ProductService.UpdateProduct(uid, &updatedProduct); err != nil {
        return response.InternalServerError(c, err, "Cannot update product")
    }
    return response.SuccessHandler(c, fiber.StatusOK, "Product updated successfully", updatedProduct)
}

// DeleteProduct handles the request to delete a product by its UID

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
    uid := c.Params("uuid")
    if err := h.ProductService.DeleteProduct(uid); err != nil {
        return response.InternalServerError(c, err, "Cannot delete product")
    }
    return response.SuccessHandler(c, fiber.StatusOK, "Product deleted successfully", nil)
}