package models

import (
	"crypto/rand"
	"encoding/hex"

	"gorm.io/gorm"
)

// defenisi struktur kategori produk
type KategoriProduk struct {
	UUID              string `gorm:"type:char(36);unique;not null" json:"uuid"` // UUID sebagai kolom unik
	NamaKategori      string `json:"nama_kategori"`
	DeskripsiKategori string `json:"deskripsi_kategori"`
	gorm.Model
}

func (k *KategoriProduk) BeforeCreate(tx *gorm.DB) (err error) {
	id := make([]byte, 16)
	_, err = rand.Read(id)
	if err != nil {
		return err
	}
	k.UUID = hex.EncodeToString(id)
	return
}

func MigrateKategori(db *gorm.DB) error {
	return db.AutoMigrate(&KategoriProduk{})
}
