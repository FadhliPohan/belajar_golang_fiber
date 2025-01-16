package services

import (
	"belajar_fiber/models"
	"belajar_fiber/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

// GetAllProduct retrieves paginated and filtered products  
func (s *ProductService) GetAllProduct(page int, size int, filters map[string]string) ([]models.Product, int64, error) {  
	var products []models.Product  
	var totalData int64  
  
	// Apply filters using the GetFilter function from utils  
	query := s.DB.Model(&models.Product{})  
	query = utils.GetFilter(filters, query) // Use utils.GetFilter  
  
	// Count total number of products after applying filters  
	if err := query.Count(&totalData).Error; err != nil {  
		return nil, 0, err  
	}  
  
	// Get limit and offset for pagination  
	limit, offset := utils.GetLimitOffset(size, page) // Use utils.GetLimitOffset  
  
	// Retrieve paginated products after applying filters  
	if err := query.Offset(offset).Limit(limit).Find(&products).Error; err != nil {  
		return nil, 0, err  
	}  
  
	return products, totalData, nil  
}  

func (s *ProductService) GetProductByID(id string)(*models.Product,error) {
	var product models.Product
	if err := s.DB.Where("uuid = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *ProductService) CreateProduct(varproduct *models.Product) error {
	varproduct.UUID = uuid.New().String();
	    return s.DB.Create(varproduct).Error
}

func (s *ProductService) UpdateProduct(UUID string, varproduct *models.Product) error {
	var existingProduct models.Product
    if err := s.DB.Where("uuid =?", UUID).First(&existingProduct).Error; err!= nil {
        return err
    }
    existingProduct.Nama = varproduct.Nama
    existingProduct.Produsen = varproduct.Produsen
    existingProduct.KodeBarang = varproduct.KodeBarang
    existingProduct.Kategori = varproduct.Kategori
    existingProduct.Deskripsi = varproduct.Deskripsi
    return s.DB.Save(&existingProduct).Error
}

func (s *ProductService) DeleteProduct(UUID string) error {
	var product models.Product
    if err := s.DB.Where("uuid =?", UUID).First(&product).Error; err!= nil {
        return err
    }
    return s.DB.Delete(&product).Error
}