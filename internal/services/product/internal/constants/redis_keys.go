package constants

import "time"

// Redis Key Patterns for Product Service
const (
	// Product-related keys
	KeyProductDetail = "product:detail:%d"
	KeyProductSKU    = "product:sku:%d"
	KeyProductList   = "product:list:category:%s"
	KeyProductSearch = "product:search:%s"
	KeyProductPrice  = "product:price:sku:%d"
	KeyProductStock  = "product:stock:sku:%d"
	KeyProductAttrs  = "product:attributes:%d"

	// Category-related keys
	KeyCategoryTree     = "product:category:tree"
	KeyCategoryProducts = "product:category:%d:products"

	// Analytics keys
	KeyTrendingDaily = "product:trending:daily:%s"
)

// Redis TTL for Product Service
const (
	TTLProductDetail    = time.Hour        // 1 hour - stable data
	TTLProductSKU       = 15 * time.Minute // 15 min - price/stock changes
	TTLProductList      = 30 * time.Minute // 30 min - category listings
	TTLProductSearch    = 30 * time.Minute // 30 min - search results
	TTLProductPrice     = 15 * time.Minute // 15 min - pricing sensitive
	TTLProductStock     = 5 * time.Minute  // 5 min - inventory critical
	TTLProductAttrs     = 2 * time.Hour    // 2 hours - rarely changes
	TTLCategoryTree     = 6 * time.Hour    // 6 hours - very stable
	TTLCategoryProducts = 2 * time.Hour    // 2 hours - category content
	TTLTrendingDaily    = time.Hour        // 1 hour - analytics data
)

// GetProductRedisConfig returns Redis key patterns and TTLs for product service
func GetProductRedisConfig() (map[string]string, map[string]time.Duration) {
	keyPatterns := map[string]string{
		"product_detail":    KeyProductDetail,
		"product_sku":       KeyProductSKU,
		"product_list":      KeyProductList,
		"product_search":    KeyProductSearch,
		"product_price":     KeyProductPrice,
		"product_stock":     KeyProductStock,
		"product_attrs":     KeyProductAttrs,
		"category_tree":     KeyCategoryTree,
		"category_products": KeyCategoryProducts,
		"trending_daily":    KeyTrendingDaily,
	}

	ttlConfig := map[string]time.Duration{
		"product_detail":    TTLProductDetail,
		"product_sku":       TTLProductSKU,
		"product_list":      TTLProductList,
		"product_search":    TTLProductSearch,
		"product_price":     TTLProductPrice,
		"product_stock":     TTLProductStock,
		"product_attrs":     TTLProductAttrs,
		"category_tree":     TTLCategoryTree,
		"category_products": TTLCategoryProducts,
		"trending_daily":    TTLTrendingDaily,
	}

	return keyPatterns, ttlConfig
}
