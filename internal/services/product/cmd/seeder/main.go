package main

import (
	"flag"
	"github.com/hthinh24/go-store/services/product/internal/seeder"
	"log"
	"time"
)

func main() {
	// Command line flags
	var (
		baseURL  = flag.String("url", "http://localhost:8081/api/v1/products/no-sku", "Base URL of the product service")
		mode     = flag.String("mode", "random", "Seeding mode: random, category, diverse")
		count    = flag.Int("count", 10, "Number of products to create")
		category = flag.String("category", "", "Category name for category mode")
		delay    = flag.Duration("delay", 100*time.Millisecond, "Delay between requests")
	)
	flag.Parse()

	log.Printf("🚀 Starting Product Seeder")
	log.Printf("📍 Target URL: %s", *baseURL)
	log.Printf("🎯 Mode: %s", *mode)
	log.Printf("📦 Count: %d", *count)
	log.Printf("⏱️  Delay: %v", *delay)

	seedingService := seeder.NewSeedingService(*baseURL)
	config := seeder.SeedingConfig{
		BaseURL:              *baseURL,
		BatchSize:            10,
		DelayBetweenRequests: *delay,
	}

	var result *seeder.SeedingResult

	switch *mode {
	case "random":
		result = seedingService.SeedRandomProducts(*count, config)

	case "category":
		if *category == "" {
			log.Fatal("❌ Category mode requires -category flag")
		}
		categories := []string{*category}
		result = seedingService.SeedProductsByCategory(categories, *count, config)

	case "diverse":
		result = seedingService.SeedDiverseProductMix(*count, config)

	case "batch":
		// Predefined batch of categories
		categories := []string{
			"Men's Clothing",
			"Women's Clothing",
			"Computers & Laptops",
			"Mobile Phones",
			"Gaming",
			"Shoes",
		}
		productsPerCategory := *count / len(categories)
		if productsPerCategory < 1 {
			productsPerCategory = 1
		}
		result = seedingService.SeedProductsByCategory(categories, productsPerCategory, config)

	default:
		log.Fatal("❌ Invalid mode. Use: random, category, diverse, or batch")
	}

	// Print final summary
	if result.SuccessfulSeeds > 0 {
		log.Printf("\n🎉 Successfully seeded %d products!", result.SuccessfulSeeds)
	} else {
		log.Printf("\n😞 No products were successfully seeded")
	}
}
