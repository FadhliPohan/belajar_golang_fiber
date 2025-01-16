package main

import (
    "belajar_crud/database"
    "belajar_crud/models"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // Definisikan rute dasar

	db := database.Connect()

    // Migrasi tabel User
    models.Migrate(db)

    // app.Get("/", func(c *fiber.Ctx) error {
    //     return c.SendString("apakabar semua!")
    // })

    // Jalankan server di port 3000
    app.Listen(":3000")
}