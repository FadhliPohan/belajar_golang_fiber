package services

import (
	"belajar_fiber/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

func (s *ProductService) GetAllProduct()([]models.Product,error) {
	var products []models.Product
	if err := s.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
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