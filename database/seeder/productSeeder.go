package seeder

import (
	"belajar_fiber/models"

	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) error {
	// Dummy data
	products := []models.Product{
		{
			Nama:       "Produk A",
			Produsen:   "Produsen A",
			KodeBarang: "A001",
			Kategori:   "Kategori A",
			Deskripsi:  "Deskripsi untuk Produk A",
		},
		{
			Nama:       "Produk B",
			Produsen:   "Produsen B",
			KodeBarang: "B001",
			Kategori:   "Kategori B",
			Deskripsi:  "Deskripsi untuk Produk B",
		},
		{
			Nama:       "Produk C",
			Produsen:   "Produsen C",
			KodeBarang: "C001",
			Kategori:   "Kategori C",
			Deskripsi:  "Deskripsi untuk Produk C",
		},
	}

	// Insert dummy data into the database
	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			return err // Return error if insertion fails
		}
	}

	return nil // Return nil if seeding is successful
}
