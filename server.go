package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	// Serve static files
	app.Static("/", "./")

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

	// Save to file (simple persistence)
	saveWaitlist()

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

func saveWaitlist() {
	data, err := json.MarshalIndent(waitlist, "", "  ")
	if err != nil {
		log.Printf("Error marshaling waitlist: %v", err)
		return
	}

	if err := os.WriteFile("waitlist.json", data, 0644); err != nil {
		log.Printf("Error saving waitlist: %v", err)
	}
}

func loadWaitlist() {
	data, err := os.ReadFile("waitlist.json")
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("Error loading waitlist: %v", err)
		}
		return
	}

	if err := json.Unmarshal(data, &waitlist); err != nil {
		log.Printf("Error unmarshaling waitlist: %v", err)
	}
}

func init() {
	loadWaitlist()
}