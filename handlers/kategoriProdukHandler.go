package handlers

import (
	"belajar_fiber/models"
	"belajar_fiber/response"
	"belajar_fiber/services"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type KategoriProductHandler struct {
	KategoriProductService *services.KategoriProductService
}


// GetKategoriProductByID handles the request to fetch a single KategoriProduct by its ID
func (h *KategoriProductHandler) GetAllCategory(c *fiber.Ctx) error {
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
	if nama_kategori := c.Query("nama_kategori"); nama_kategori != "" {
		filters["nama_kategori"] = nama_kategori // Assuming you want to filter by product nama_kategori
	}
	if deskripsi_kategori := c.Query("deskripsi_kategori"); deskripsi_kategori != "" {
		filters["deskripsi_kategori"] = deskripsi_kategori // Assuming you want to filter by deskripsi_kategori
	}

	// Call the service to get paginated and filtered products
	products, totalData, err := h.KategoriProductService.GetAllCategory(page, size, filters)
	if err != nil {
		return response.InternalServerError(c, err, "Cannot retrieve products")
	}

	// Return the paginated response
	return response.SuccessHandlerPaginate(c, fiber.StatusOK, "Products retrieved successfully", products, page, size, int(totalData))
}

//getKategoriProductByUUID handles the request to get the product

func (h *KategoriProductHandler) GetKategoriCategoryByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	KategoriProduct, err := h.KategoriProductService.GetCategoryByID(uuid)
	if err != nil {
		return response.InternalServerError(c, err, "Cannot retrieve KategoriProduct")
	}
	if KategoriProduct == nil {
		return response.BadRequest(c, err, "KategoriProduct not found")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "KategoriProduct retrieved successfully", KategoriProduct)
}

// CreateKategoriProduct handles the request to create a new KategoriProduct
func (h *KategoriProductHandler) CreateCategory(c *fiber.Ctx) error {
	var kategoriproduct models.KategoriProduk
	if err := c.BodyParser(&kategoriproduct); err != nil {
		return response.BadRequest(c, err, "Cannot parse JSON")
	}
	if err := h.KategoriProductService.CreateCategory(&kategoriproduct); err != nil {
		return response.InternalServerError(c, err, "Cannot create product")
	}
	return response.SuccessHandler(c, fiber.StatusCreated, "Product created successfully", kategoriproduct)
}

// UpdateProduct handles the request to update a product by its UID

func (h *KategoriProductHandler) UpdateCategory(c *fiber.Ctx) error {  
	// Retrieve the UUID from the request parameters  
	uid := c.Params("uuid")  
  
	// Parse the JSON body into the updatedkategoriProduct struct  
	var updatedkategoriProduct models.KategoriProduk  
	if err := c.BodyParser(&updatedkategoriProduct); err != nil {  
		return response.BadRequest(c, err, "Cannot parse JSON")  
	}  
  
	// Attempt to update the category in the database  
	if err := h.KategoriProductService.UpdateCategory(uid, &updatedkategoriProduct); err != nil {  
		return response.InternalServerError(c, err, "Cannot update product")  
	}  
  
	// Retrieve the updated category to return in the response  
	KategoriProduct, err := h.KategoriProductService.GetCategoryByID(uid)  
	if err != nil {  
		return response.InternalServerError(c, err, "Cannot retrieve updated product")  
	}  
  
	// Return a success response with the updated category  
	return response.SuccessHandler(c, fiber.StatusOK, "Product updated successfully", KategoriProduct)  
}

// DeleteProduct handles the request to delete a product by its UID
func (h *KategoriProductHandler) DeleteCategory(c *fiber.Ctx) error {
	uid := c.Params("uuid")
	if err := h.KategoriProductService.DeleteCategory(uid); err != nil {
		return response.InternalServerError(c, err, "Cannot delete product")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "Product deleted successfully", nil)
}

// return all request
func (h *KategoriProductHandler) GetAllRequest(c *fiber.Ctx) error {  
	// Retrieve the 'namakategori' parameter from the query string  
	namaKategori := c.Query("namakategori")  
  
	// Validate that 'namakategori' is not empty  
	if namaKategori == "" {   
		return response.ErrorHandler(c,errors.New("error"), fiber.StatusBadRequest, "Parameter 'namakategori' harus diisi")  
	}  
  
	// Retrieve all request data  
	requestData := map[string]interface{}{  
		"method":      c.Method(),  
		"url":         c.OriginalURL(),  
		"headers":     c.GetReqHeaders(),  
		"body":        c.Body(),  
		"body_string": string(c.Body()), // This will give you the raw body of the request  
		"namakategori": namaKategori,    // Include the validated 'namakategori' in the response data  
	}  
  
	// Return the request data in the response  
	return response.SuccessHandler(c, http.StatusOK, "All request", requestData)  
}  
