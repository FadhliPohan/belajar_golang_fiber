package models  
  
import (  
	"crypto/rand"  
	"encoding/hex"  
	"gorm.io/gorm"  
)  
  
// Product mendefinisikan struktur untuk produk  
type Product struct {  
    UUID       string `gorm:"type:char(36);unique;not null" json:"uuid"` // UUID sebagai kolom unik  
	Nama       string `json:"nama"`                                       // Nama produk  
	Produsen   string `json:"produsen"`                                   // Nama produsen  
	KodeBarang string `json:"kode_barang"`                                // Kode unik untuk barang  
	Kategori   string `json:"kategori"`                                   // Kategori produk  
	Deskripsi  string `json:"deskripsi"`                                  // Deskripsi produk  
	gorm.Model // Menyertakan gorm.Model untuk otomatisasi timestamp dan soft delete  
}  
  
// BeforeCreate hook untuk mengatur UUID  
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {  
	id := make([]byte, 16)  
	_, err = rand.Read(id)  
	if err != nil {  
		return err  
	}  
	p.UUID = hex.EncodeToString(id) // Menghasilkan UUID  
	return  
}  
  
// MigrateProduct function untuk membuat tabel Product  
func MigrateProduct(db *gorm.DB) error {  
	return db.AutoMigrate(&Product{}) // Mengembalikan error jika terjadi  
}  
