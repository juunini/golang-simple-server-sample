package main

import (
	"embed"
	"log"
	"net/http"
	"server/src/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//go:embed html/*
var public embed.FS

func main() {
	db, err := gorm.Open(sqlite.Open("sqlite.db"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Task{})

	app := fiber.New(fiber.Config{
		Views: html.NewFileSystem(http.FS(public), ".html"),
	})

	app.Get("/api/tasks", func(c *fiber.Ctx) error {
		var tasks []models.Task
		db.Find(&tasks)
		return c.JSON(tasks)
	})

	app.Post("/api/tasks", func(c *fiber.Ctx) error {
		task := new(models.Task)

		if err := c.BodyParser(task); err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusBadRequest)
		}

		db.Create(&task)
		return c.JSON(task)
	})

	app.Delete("/api/tasks/:id", func(c *fiber.Ctx) error {
		var task models.Task
		db.Delete(&task, c.Params("id"))
		return c.SendStatus(204)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		var tasks []models.Task
		db.Find(&tasks)

		return c.Render("html/index", fiber.Map{
			"tasks": tasks,
		})
	})

	app.Listen(":3000")
}
