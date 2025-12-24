package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type WaitlistEntry struct {
	Email     string `json:"email"`
	Timestamp string `json:"timestamp"`
}

var waitlist []WaitlistEntry

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Beeline Waitlist",
	})

	// Middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// API endpoints
	app.Post("/api/waitlist", addToWaitlist)
	app.Get("/api/waitlist/count", getWaitlistCount)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
			"waitlist_count": len(waitlist),
			"launch_date": "2025-12-31T23:59:59Z",
		})
	})

	// Handle all other routes (SPA fallback)
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Beeline Waitlist server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func addToWaitlist(c *fiber.Ctx) error {
	var entry WaitlistEntry
	if err := c.BodyParser(&entry); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if entry.Email == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Email is required",
		})
	}

	entry.Timestamp = time.Now().Format(time.RFC3339)

	// Check if email already exists
	for _, existing := range waitlist {
		if existing.Email == entry.Email {
			return c.Status(409).JSON(fiber.Map{
				"error": "Email already registered",
			})
		}
	}

	waitlist = append(waitlist, entry)

	return c.JSON(fiber.Map{
		"message": "Successfully added to waitlist",
		"position": len(waitlist),
	})
}

func getWaitlistCount(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"count": len(waitlist),
	})
}