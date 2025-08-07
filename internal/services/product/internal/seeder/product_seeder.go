package seeder

import (
	"fmt"
	"github.com/hthinh24/go-store/services/product/internal/dto/request"
	"math/rand"
	"strings"
	"time"
)

type ProductSeeder struct {
	templates []ProductTemplate
}

type ProductTemplate struct {
	Category    string             `json:"category"`
	BrandID     int64              `json:"brand_id"`
	BaseProduct ProductInfo        `json:"base_product"`
	Options     map[int64][]string `json:"options"` // Changed to int64 keys
	Attributes  map[int64][]string `json:"attributes"`
	PriceRange  PriceRange         `json:"price_range"`
}

type ProductInfo struct {
	NameTemplate     string `json:"name_template"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
	Tags             string `json:"tags"`
	Status           string `json:"status"`
}

type PriceRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

func NewProductSeeder() *ProductSeeder {
	return &ProductSeeder{
		templates: getProductTemplates(),
	}
}

func (ps *ProductSeeder) GenerateRandomProduct() *request.CreateProductWithoutSKURequest {
	template := ps.templates[rand.Intn(len(ps.templates))]

	// Generate random price within range
	price := template.PriceRange.Min + rand.Float64()*(template.PriceRange.Max-template.PriceRange.Min)
	price = float64(int(price*100)) / 100 // Round to 2 decimal places

	// Generate unique product name with timestamp
	timestamp := time.Now().Unix()
	productName := fmt.Sprintf(template.BaseProduct.NameTemplate, timestamp)

	// Generate slug from product name
	slug := generateSlug(productName)

	// Add some randomness to description
	descriptions := []string{
		template.BaseProduct.Description,
		template.BaseProduct.Description + " - Premium Quality",
		template.BaseProduct.Description + " - Best Seller",
		template.BaseProduct.Description + " - Limited Edition",
	}

	// Generate image URL
	imageURL := generateImageURL(template.Category, productName)

	// Generate sale price (30% chance to have sale price)
	var salePrice *float64
	if rand.Float32() < 0.3 {
		sale := price * (0.8 + rand.Float64()*0.15) // 80-95% of base price
		sale = float64(int(sale*100)) / 100
		salePrice = &sale
	}

	return &request.CreateProductWithoutSKURequest{
		Name:              productName,
		Description:       descriptions[rand.Intn(len(descriptions))],
		ShortDescription:  template.BaseProduct.ShortDescription,
		ImageURL:          imageURL,
		Slug:              slug,
		BasePrice:         price,
		SalePrice:         salePrice,
		IsFeatured:        rand.Float32() < 0.2, // 20% chance to be featured
		BrandID:           template.BrandID,
		CategoryID:        getCategoryIDByName(template.Category),
		UserID:            1, // Default admin user ID
		ProductAttributes: template.Attributes,
		OptionValues:      template.Options,
	}
}

func (ps *ProductSeeder) GenerateProductsByCategory(category string, count int) []*request.CreateProductWithoutSKURequest {
	var products []*request.CreateProductWithoutSKURequest
	var categoryTemplates []ProductTemplate

	// Filter templates by category
	for _, template := range ps.templates {
		if template.Category == category {
			categoryTemplates = append(categoryTemplates, template)
		}
	}

	if len(categoryTemplates) == 0 {
		return products
	}

	for i := 0; i < count; i++ {
		template := categoryTemplates[rand.Intn(len(categoryTemplates))]

		// Generate random price within range
		price := template.PriceRange.Min + rand.Float64()*(template.PriceRange.Max-template.PriceRange.Min)
		price = float64(int(price*100)) / 100

		// Generate unique product name with timestamp
		timestamp := time.Now().Unix() + int64(i)
		productName := fmt.Sprintf(template.BaseProduct.NameTemplate, timestamp)

		// Generate slug and image URL
		slug := generateSlug(productName)
		imageURL := generateImageURL(template.Category, productName)

		// Generate sale price (30% chance)
		var salePrice *float64
		if rand.Float32() < 0.3 {
			sale := price * (0.8 + rand.Float64()*0.15)
			sale = float64(int(sale*100)) / 100
			salePrice = &sale
		}

		product := &request.CreateProductWithoutSKURequest{
			Name:              productName,
			Description:       template.BaseProduct.Description,
			ShortDescription:  template.BaseProduct.ShortDescription,
			ImageURL:          imageURL,
			Slug:              slug,
			BasePrice:         price,
			SalePrice:         salePrice,
			IsFeatured:        rand.Float32() < 0.2,
			BrandID:           template.BrandID,
			CategoryID:        getCategoryIDByName(template.Category),
			UserID:            1, // Default admin user ID
			ProductAttributes: template.Attributes,
			OptionValues:      template.Options,
		}

		products = append(products, product)
		time.Sleep(time.Millisecond * 10)
	}

	return products
}

func getProductTemplates() []ProductTemplate {
	return []ProductTemplate{
		// Fashion - T-Shirts
		{
			Category: "Men's Clothing",
			BrandID:  11, // Nike
			BaseProduct: ProductInfo{
				NameTemplate:     "Nike Classic T-Shirt %d",
				Description:      "Comfortable cotton t-shirt perfect for casual wear and sports activities. Made from high-quality materials for durability and comfort.",
				ShortDescription: "Comfortable cotton t-shirt perfect for casual wear",
				Tags:             "casual,sports,comfortable",
				Status:           "active",
			},
			Options: map[int64][]string{
				2: {"S", "M", "L", "XL"},      // Size (option ID 2)
				1: {"Black", "White", "Navy"}, // Color (option ID 1)
			},
			Attributes: map[int64][]string{
				11: {"S", "M", "L", "XL"},      // Size
				12: {"Cotton", "Cotton Blend"}, // Material
				17: {"Black", "White", "Navy"}, // Color
			},
			PriceRange: PriceRange{Min: 25.99, Max: 49.99},
		},

		// Fashion - Jeans
		{
			Category: "Men's Clothing",
			BrandID:  16, // Levi's
			BaseProduct: ProductInfo{
				NameTemplate:     "Levi's 501 Original Jeans %d",
				Description:      "Classic straight-fit jeans made from premium denim. Timeless style that never goes out of fashion.",
				ShortDescription: "Classic straight-fit jeans made from premium denim",
				Tags:             "denim,classic,everyday",
				Status:           "active",
			},
			Options: map[int64][]string{
				2: {"30", "32", "34", "36"},  // Size (option ID 2)
				1: {"Blue", "Black", "Gray"}, // Color (option ID 1)
			},
			Attributes: map[int64][]string{
				11: {"30", "32", "34", "36"},  // Size
				12: {"Denim"},                 // Material
				17: {"Blue", "Black", "Gray"}, // Color
			},
			PriceRange: PriceRange{Min: 79.99, Max: 129.99},
		},

		// Electronics - Laptops
		{
			Category: "Computers & Laptops",
			BrandID:  1, // Apple
			BaseProduct: ProductInfo{
				NameTemplate:     "MacBook Air %d",
				Description:      "Lightweight and powerful laptop with M2 chip for productivity and creativity. Perfect for professionals and students.",
				ShortDescription: "Lightweight laptop with M2 chip for productivity",
				Tags:             "laptop,productivity,portable",
				Status:           "active",
			},
			Options: map[int64][]string{
				3: {"256GB", "512GB", "1TB"},            // Storage (option ID 3)
				4: {"8GB", "16GB"},                      // Memory (option ID 4)
				1: {"Silver", "Space Gray", "Midnight"}, // Color (option ID 1)
			},
			Attributes: map[int64][]string{
				1:  {"Apple M2"},                          // Processor
				2:  {"8GB", "16GB"},                       // RAM
				3:  {"256GB SSD", "512GB SSD", "1TB SSD"}, // Storage
				5:  {"13.3 inch"},                         // Screen Size
				6:  {"macOS"},                             // OS
				17: {"Silver", "Space Gray", "Midnight"},  // Color
			},
			PriceRange: PriceRange{Min: 999.99, Max: 1899.99},
		},

		// Electronics - Smartphones
		{
			Category: "Mobile Phones",
			BrandID:  2, // Samsung
			BaseProduct: ProductInfo{
				NameTemplate:     "Samsung Galaxy S24 %d",
				Description:      "Premium smartphone with advanced camera system and 5G connectivity. Experience the future of mobile technology.",
				ShortDescription: "Premium smartphone with advanced camera and 5G",
				Tags:             "smartphone,5G,camera",
				Status:           "active",
			},
			Options: map[int64][]string{
				3: {"128GB", "256GB", "512GB"},           // Storage (option ID 3)
				1: {"Black", "White", "Purple", "Green"}, // Color (option ID 1)
			},
			Attributes: map[int64][]string{
				2:  {"8GB"},                               // RAM
				3:  {"128GB", "256GB", "512GB"},           // Storage
				5:  {"6.1 inch"},                          // Screen Size
				6:  {"Android 14"},                        // OS
				8:  {"50MP"},                              // Camera
				17: {"Black", "White", "Purple", "Green"}, // Color
			},
			PriceRange: PriceRange{Min: 799.99, Max: 1199.99},
		},

		// Home & Garden - Furniture
		{
			Category: "Furniture",
			BrandID:  17, // IKEA
			BaseProduct: ProductInfo{
				NameTemplate:     "IKEA Modern Chair %d",
				Description:      "Stylish and comfortable chair perfect for home or office use. Ergonomic design meets modern aesthetics.",
				ShortDescription: "Stylish and comfortable chair for home or office",
				Tags:             "furniture,chair,modern",
				Status:           "active",
			},
			Options: map[int64][]string{
				1: {"Black", "White", "Gray", "Brown"}, // Color (option ID 1)
				5: {"Classic", "Modern"},               // Style (option ID 5)
			},
			Attributes: map[int64][]string{
				12: {"Wood", "Metal"},                   // Material
				17: {"Black", "White", "Gray", "Brown"}, // Color
				19: {"60x60x80cm"},                      // Dimensions
			},
			PriceRange: PriceRange{Min: 89.99, Max: 299.99},
		},

		// Sports & Outdoors - Shoes
		{
			Category: "Shoes",
			BrandID:  11, // Nike
			BaseProduct: ProductInfo{
				NameTemplate:     "Nike Air Max Running Shoes %d",
				Description:      "High-performance running shoes with Air Max technology. Designed for comfort and performance.",
				ShortDescription: "High-performance running shoes with Air Max technology",
				Tags:             "running,sports,comfortable",
				Status:           "active",
			},
			Options: map[int64][]string{
				2: {"7", "8", "9", "10", "11", "12"}, // Size (option ID 2)
				1: {"Black", "White", "Red", "Blue"}, // Color (option ID 1)
			},
			Attributes: map[int64][]string{
				11: {"7", "8", "9", "10", "11", "12"}, // Size
				12: {"Synthetic", "Leather"},          // Material
				17: {"Black", "White", "Red", "Blue"}, // Color
			},
			PriceRange: PriceRange{Min: 119.99, Max: 179.99},
		},

		// Electronics - Gaming
		{
			Category: "Gaming",
			BrandID:  4, // Microsoft
			BaseProduct: ProductInfo{
				NameTemplate:     "Xbox Wireless Controller %d",
				Description:      "Precision gaming controller with enhanced D-pad and textured grips. Take your gaming to the next level.",
				ShortDescription: "Precision gaming controller with enhanced D-pad",
				Tags:             "gaming,controller,xbox",
				Status:           "active",
			},
			Options: map[int64][]string{
				1: {"Black", "White", "Red", "Blue"}, // Color (option ID 1)
				7: {"Standard", "Elite"},             // Edition (option ID 7)
			},
			Attributes: map[int64][]string{
				9:  {"Wireless", "USB-C"},             // Connectivity
				17: {"Black", "White", "Red", "Blue"}, // Color
				20: {"Battery"},                       // Power Source
			},
			PriceRange: PriceRange{Min: 59.99, Max: 179.99},
		},

		// Women's Clothing - Dresses
		{
			Category: "Women's Clothing",
			BrandID:  13, // Zara
			BaseProduct: ProductInfo{
				NameTemplate:     "Zara Summer Dress %d",
				Description:      "Elegant summer dress perfect for casual and formal occasions. Made from premium fabrics for all-day comfort.",
				ShortDescription: "Elegant summer dress for casual and formal occasions",
				Tags:             "dress,summer,elegant",
				Status:           "active",
			},
			Options: map[int64][]string{
				2: {"XS", "S", "M", "L", "XL"},               // Size (option ID 2)
				1: {"Black", "White", "Red", "Blue", "Pink"}, // Color (option ID 1)
			},
			Attributes: map[int64][]string{
				11: {"XS", "S", "M", "L", "XL"},               // Size
				12: {"Cotton", "Polyester", "Silk"},           // Material
				17: {"Black", "White", "Red", "Blue", "Pink"}, // Color
			},
			PriceRange: PriceRange{Min: 49.99, Max: 129.99},
		},
	}
}

// Helper functions
func generateSlug(name string) string {
	// Convert to lowercase and replace spaces with hyphens
	slug := strings.TrimSpace(name)
	slug = strings.ToLower(slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "'", "")
	slug = strings.ReplaceAll(slug, "&", "and")
	slug = strings.ReplaceAll(slug, ".", "")
	slug = strings.ReplaceAll(slug, ",", "")
	slug = strings.ReplaceAll(slug, "(", "")
	slug = strings.ReplaceAll(slug, ")", "")

	// Remove multiple consecutive hyphens
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}

	// Remove leading/trailing hyphens
	slug = strings.Trim(slug, "-")

	// Use 8-character UUID for uniqueness (recommended approach)
	uuid := generateShortUUID()
	slug = fmt.Sprintf("%s-%s", slug, uuid)

	return slug
}

func generateShortUUID() string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+-"
	result := make([]byte, 8)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func generateImageURL(category, productName string) string {
	// Generate a placeholder image URL based on category
	categoryMap := map[string]string{
		"Men's Clothing":      "mens-clothing",
		"Women's Clothing":    "womens-clothing",
		"Shoes":               "shoes",
		"Computers & Laptops": "laptops",
		"Mobile Phones":       "smartphones",
		"Gaming":              "gaming",
		"Furniture":           "furniture",
	}

	categorySlug := categoryMap[category]
	if categorySlug == "" {
		categorySlug = "products"
	}

	return fmt.Sprintf("https://images.example.com/%s/%s.jpg", categorySlug, generateSlug(productName))
}

func getCategoryIDByName(categoryName string) int64 {
	// This is a mapping based on your seed data
	categoryMap := map[string]int64{
		"Electronics":           1,
		"Clothing & Fashion":    2,
		"Home & Garden":         3,
		"Sports & Outdoors":     4,
		"Books & Media":         5,
		"Health & Beauty":       6,
		"Automotive":            7,
		"Toys & Games":          8,
		"Computers & Laptops":   9,
		"Mobile Phones":         10,
		"Audio & Video":         11,
		"Gaming":                12,
		"Cameras & Photography": 13,
		"Smart Home":            14,
		"Men's Clothing":        15,
		"Women's Clothing":      16,
		"Children's Clothing":   17,
		"Shoes":                 18,
		"Accessories":           19,
		"Furniture":             20,
		"Kitchen & Dining":      21,
		"Home Decor":            22,
		"Garden & Outdoor":      23,
		"Fitness & Exercise":    24,
		"Outdoor Recreation":    25,
		"Team Sports":           26,
		"Water Sports":          27,
	}

	if id, exists := categoryMap[categoryName]; exists {
		return id
	}
	return 1 // Default to Electronics
}
