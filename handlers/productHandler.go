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
	// Get pagination parameters from query  
	page := c.QueryInt("page", 1) // Default to page 1  
	if page < 1 {  
		page = 1 // Ensure page is at least 1  
	}  
  
	size := c.QueryInt("size", 2) // Default to size 10  
	if size < 1 {  
		size = 10 // Ensure size is at least 1  
	}  
  
	// Call the service to get paginated products  
	products, totalData, err := h.ProductService.GetAllProduct(page, size)  
	if err != nil {  
		return response.InternalServerError(c, err, "Cannot retrieve products")  
	}  
  
	// Return the paginated response  
	return response.SuccessHandlerPaginate(c, fiber.StatusOK, "Products retrieved successfully", products, page, size, int(totalData))  
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
