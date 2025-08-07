package seeder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hthinh24/go-store/services/product/internal/dto/request"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type SeedingService struct {
	baseURL       string
	httpClient    *http.Client
	productSeeder *ProductSeeder
}

type SeedingConfig struct {
	BaseURL              string        `json:"base_url"`
	BatchSize            int           `json:"batch_size"`
	DelayBetweenRequests time.Duration `json:"delay_between_requests"`
}

type SeedingResult struct {
	TotalRequests   int           `json:"total_requests"`
	SuccessfulSeeds int           `json:"successful_seeds"`
	FailedSeeds     int           `json:"failed_seeds"`
	Errors          []string      `json:"errors"`
	Duration        time.Duration `json:"duration"`
}

func NewSeedingService(baseURL string) *SeedingService {
	return &SeedingService{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		productSeeder: NewProductSeeder(),
	}
}

// SeedRandomProducts generates and posts random products across all categories
func (ss *SeedingService) SeedRandomProducts(count int, config SeedingConfig) *SeedingResult {
	startTime := time.Now()
	result := &SeedingResult{
		TotalRequests: count,
		Errors:        make([]string, 0),
	}

	log.Printf("Starting to seed %d random products...", count)

	for i := 0; i < count; i++ {
		// Generate random product
		productData := ss.productSeeder.GenerateRandomProduct()

		// Post to API
		success, err := ss.postProduct(productData)
		if success {
			result.SuccessfulSeeds++
			log.Printf("✅ Successfully created product %d: %s", i+1, productData.Name)
		} else {
			result.FailedSeeds++
			errorMsg := fmt.Sprintf("Failed to create product %d: %v", i+1, err)
			result.Errors = append(result.Errors, errorMsg)
			log.Printf("❌ %s", errorMsg)
		}

		// Delay between requests to avoid overwhelming the server
		if i < count-1 {
			time.Sleep(config.DelayBetweenRequests)
		}
	}

	result.Duration = time.Since(startTime)
	ss.logSeedingResults(result)
	return result
}

// SeedProductsByCategory generates and posts products for specific categories
func (ss *SeedingService) SeedProductsByCategory(categories []string, countPerCategory int, config SeedingConfig) *SeedingResult {
	startTime := time.Now()
	totalCount := len(categories) * countPerCategory
	result := &SeedingResult{
		TotalRequests: totalCount,
		Errors:        make([]string, 0),
	}

	log.Printf("Starting to seed %d products per category for %d categories...", countPerCategory, len(categories))

	for _, category := range categories {
		log.Printf("🔄 Seeding products for category: %s", category)

		products := ss.productSeeder.GenerateProductsByCategory(category, countPerCategory)

		for i, productData := range products {
			success, err := ss.postProduct(productData)
			if success {
				result.SuccessfulSeeds++
				log.Printf("✅ Successfully created %s product %d: %s", category, i+1, productData.Name)
			} else {
				result.FailedSeeds++
				errorMsg := fmt.Sprintf("Failed to create %s product %d: %v", category, i+1, err)
				result.Errors = append(result.Errors, errorMsg)
				log.Printf("❌ %s", errorMsg)
			}

			time.Sleep(config.DelayBetweenRequests)
		}
	}

	result.Duration = time.Since(startTime)
	ss.logSeedingResults(result)
	return result
}

// SeedDiverseProductMix creates a balanced mix of products across different categories
func (ss *SeedingService) SeedDiverseProductMix(totalCount int, config SeedingConfig) *SeedingResult {
	categories := []string{
		"Men's Clothing",
		"Women's Clothing",
		"Shoes",
		"Computers & Laptops",
		"Mobile Phones",
		"Gaming",
		"Furniture",
		"Sports & Outdoors",
	}

	productsPerCategory := totalCount / len(categories)
	remainder := totalCount % len(categories)

	startTime := time.Now()
	result := &SeedingResult{
		TotalRequests: totalCount,
		Errors:        make([]string, 0),
	}

	log.Printf("Creating diverse product mix: %d total products across %d categories", totalCount, len(categories))

	for i, category := range categories {
		countForCategory := productsPerCategory
		// Distribute remainder among first few categories
		if i < remainder {
			countForCategory++
		}

		if countForCategory == 0 {
			continue
		}

		log.Printf("🔄 Creating %d products for %s", countForCategory, category)

		products := ss.productSeeder.GenerateProductsByCategory(category, countForCategory)

		for j, productData := range products {
			success, err := ss.postProduct(productData)
			if success {
				result.SuccessfulSeeds++
				log.Printf("✅ [%s %d/%d] %s", category, j+1, countForCategory, productData.Name)
			} else {
				result.FailedSeeds++
				errorMsg := fmt.Sprintf("Failed to create %s product %d: %v", category, j+1, err)
				result.Errors = append(result.Errors, errorMsg)
				log.Printf("❌ %s", errorMsg)
			}

			time.Sleep(config.DelayBetweenRequests)
		}
	}

	result.Duration = time.Since(startTime)
	ss.logSeedingResults(result)
	return result
}

func (ss *SeedingService) postProduct(productData *request.CreateProductWithoutSKURequest) (bool, error) {
	jsonData, err := json.Marshal(productData)
	if err != nil {
		return false, fmt.Errorf("failed to marshal product data: %w", err)
	}

	req, err := http.NewRequest("POST", ss.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := ss.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true, nil
	}

	return false, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
}

func (ss *SeedingService) logSeedingResults(result *SeedingResult) {
	log.Printf("\n" + strings.Repeat("=", 50))
	log.Printf("🎯 SEEDING COMPLETED")
	log.Printf(strings.Repeat("=", 50))
	log.Printf("📊 Total Requests: %d", result.TotalRequests)
	log.Printf("✅ Successful Seeds: %d", result.SuccessfulSeeds)
	log.Printf("❌ Failed Seeds: %d", result.FailedSeeds)
	log.Printf("⏱️  Duration: %v", result.Duration)
	log.Printf("📈 Success Rate: %.2f%%", float64(result.SuccessfulSeeds)/float64(result.TotalRequests)*100)

	if len(result.Errors) > 0 {
		log.Printf("\n🔍 ERRORS:")
		for i, err := range result.Errors {
			if i < 5 { // Show only first 5 errors
				log.Printf("   %d. %s", i+1, err)
			}
		}
		if len(result.Errors) > 5 {
			log.Printf("   ... and %d more errors", len(result.Errors)-5)
		}
	}
	log.Printf(strings.Repeat("=", 50))
}
