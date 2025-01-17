package main

import (
	"belajar_fiber/config"
	"belajar_fiber/database"

	// "belajar_fiber/database/seeder"
	// "log"
	"belajar_fiber/models"
	"belajar_fiber/routes"


	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Initialize Fiber
	app := fiber.New(fiber.Config{BodyLimit: 35 * 1024 * 1024})

	// Load configuration
	cfg := config.LoadConfig()

	// Definisikan rute dasar
	db := database.Connect()

	// Migrasi tabel User
	models.Migrate(db)
	models.MigrateProduct(db)
	models.MigrateKategori(db)

	// if err := seeder.SeedProducts(db); err != nil {
	// 	log.Fatalf("Could not seed the database: %v", err)
	// }

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins
	}))
	app.Use(recover.New())

	// Setup Routes
	routes.SetupRoutes(app, db, cfg)

	// Rute dasar
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("baru awal bosss!")
	})

	app.Get("/halo", func(c *fiber.Ctx) error {
		return c.SendString("halo masyarakat")
	})

	// Jalankan server di port 3000
	app.Listen(":3000")
}
