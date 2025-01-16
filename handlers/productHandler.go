package handlers

import (
	"belajar_fiber/models"
	"belajar_fiber/response"
	"belajar_fiber/services"
	"strings"

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
  
	size := c.QueryInt("size", 10) // Default to size 10  
	if size < 1 {  
		size = 10 // Ensure size is at least 1  
	}  
  
	// Get filters from query parameters  
	filters := make(map[string]string)  
	if name := c.Query("nama"); name != "" {  
		filters["nama"] = name // Assuming you want to filter by product name  
	}  
	if producer := c.Query("produsen"); producer != "" {  
		filters["produsen"] = producer // Assuming you want to filter by producer  
	}  
	if category := c.Query("kategori"); category != "" {  
		filters["kategori"] = category // Assuming you want to filter by category  
	}  
  
	// Call the service to get paginated and filtered products  
	products, totalData, err := h.ProductService.GetAllProduct(page, size, filters)  
	if err != nil {  
		return response.InternalServerError(c, err, "Cannot retrieve products")  
	}  
  
	// Return the paginated response  
	return response.SuccessHandlerPaginate(c, fiber.StatusOK, "Products retrieved successfully", products, page, size, int(totalData))  
}

func (h *ProductHandler) GetAllProductsQuery(c *fiber.Ctx) error {  
	// Get pagination parameters from query  
	page := c.QueryInt("page", 1) // Default to page 1  
	if page < 1 {  
		page = 1 // Ensure page is at least 1  
	}  
  
	size := c.QueryInt("size", 10) // Default to size 10  
	if size < 1 {  
		size = 10 // Ensure size is at least 1  
	}  
  
	// Get filters from query parameters  
	filters := make([]string, 0)  
	args := make([]interface{}, 0)  
  
	if name := c.Query("nama"); name != "" {  
		filters = append(filters, "nama LIKE ?")  
		args = append(args, "%"+name+"%") // Use LIKE for partial matching  
	}  
	if producer := c.Query("produsen"); producer != "" {  
		filters = append(filters, "produsen LIKE ?")  
		args = append(args, "%"+producer+"%")  
	}  
	if category := c.Query("kategori"); category != "" {  
		filters = append(filters, "kategori LIKE ?")  
		args = append(args, "%"+category+"%")  
	}  
  
	// Construct the SQL query  
	query := "SELECT * FROM products" // Assuming your table name is 'products'  
	if len(filters) > 0 {  
		query += " WHERE " + strings.Join(filters, " AND ")  
	}  
  
	// Calculate total data count  
	var totalData int64  
	if err := h.ProductService.DB.Raw("SELECT COUNT(*) FROM products" +   
		(strings.Join(filters, " AND "))).Scan(&totalData).Error; err != nil {  
		return response.InternalServerError(c, err, "Cannot count products")  
	}  
  
	// Calculate limit and offset  
	offset := (page - 1) * size  
  
	// Execute the paginated query  
	var products []models.Product  
	if err := h.ProductService.DB.Raw(query+" LIMIT ? OFFSET ?", append(args, size, offset)...).Scan(&products).Error; err != nil {  
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
