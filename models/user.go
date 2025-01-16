package models

import (
    "crypto/rand"
    "encoding/hex"
    "gorm.io/gorm"
)

type User struct {
    ID           uint   `gorm:"primaryKey" json:"id"` // ID sebagai primary key
    UUID         string `gorm:"type:char(36);unique;not null" json:"uuid"` // UUID sebagai kolom unik
    Nama         string `json:"nama"`
    Alamat       string `json:"alamat"`
    NoHp         string `json:"no_hp"`
    Pekerjaan    string `json:"pekerjaan"`
    JenisKelamin string `json:"jenis_kelamin"`
	gorm.Model // Menyertakan gorm.Model untuk otomatisasi timestamp dan soft delete  

}

// BeforeCreate hook untuk mengatur UUID
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    id := make([]byte, 16)
    _, err = rand.Read(id)
    if err != nil {
        return err
    }
    u.UUID = hex.EncodeToString(id) // Menghasilkan UUID
    return
}

// Migrate function untuk membuat tabel User
func Migrate(db *gorm.DB) {
    db.AutoMigrate(&User{})
}
