package services

import (
	"belajar_fiber/models"
	"belajar_fiber/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KategoriProductService struct {
	DB *gorm.DB
}

// GetAllKategoriProduct fetches all KategoriProduct records
func (s *KategoriProductService) GetAllCategory(page int, size int, filters map[string]string) ([]models.KategoriProduk, int64, error) {
	var products []models.KategoriProduk
	var totalData int64

	// Apply filters using the GetFilter function from utils
	query := s.DB.Model(&models.KategoriProduk{})
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

// GetProductByID fetches a single KategoriProduct record by ID
func (s *KategoriProductService) GetCategoryByID(id string) (*models.KategoriProduk, error) {
	var product models.KategoriProduk
	if err := s.DB.Where("uuid = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// CreateProduct creates a new KategoriProduct record
func (s *KategoriProductService) CreateCategory(varkategori *models.KategoriProduk) error {
    varkategori.UUID = uuid.New().String();
    return s.DB.Create(varkategori).Error
}

// UpdateProduct updates an existing KategoriProduct record by ID
func (s *KategoriProductService) UpdateCategory(UUID string, varproduct *models.KategoriProduk) error {
	var existingProduct models.KategoriProduk
	if err := s.DB.Where("uuid =?", UUID).First(&existingProduct).Error; err != nil {
		return err
	}
	existingProduct.NamaKategori = varproduct.NamaKategori
	existingProduct.DeskripsiKategori = varproduct.DeskripsiKategori
    return s.DB.Save(&existingProduct).Error
}

// DeleteProduct deletes a KategoriProduct record by ID
func (s *KategoriProductService) DeleteCategory(id string) error {
	var product models.KategoriProduk
	if err := s.DB.Where("uuid = ?", id).First(&product).Error; err != nil {
		return err
	}
	if err := s.DB.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}
