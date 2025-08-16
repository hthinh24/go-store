package constants

import "time"

// Redis Key Patterns for Cart Service
const (
	// Cart-related keys
	KeyUserCart    = "cart:user:%d"
	KeyCartItem    = "cart:item:%d"
	KeyCartSession = "cart:session:%s"
	KeyCartTemp    = "cart:temp:%s"
	KeyCartSummary = "cart:summary:user:%d"

	// Shopping behavior keys
	KeyAbandonedCart = "cart:abandoned:user:%d"
	KeyCartHistory   = "cart:history:user:%d"
	KeyCartWishlist  = "cart:wishlist:user:%d"
)

// Redis TTL for Cart Service
const (
	TTLUserCart      = 7 * 24 * time.Hour   // 7 days - persist across sessions
	TTLCartItem      = 24 * time.Hour       // 1 day - individual items
	TTLCartSession   = 2 * time.Hour        // 2 hours - guest sessions
	TTLCartTemp      = 30 * time.Minute     // 30 min - temporary operations
	TTLCartSummary   = time.Hour            // 1 hour - calculated data
	TTLAbandonedCart = 30 * 24 * time.Hour  // 30 days - marketing recovery
	TTLCartHistory   = 90 * 24 * time.Hour  // 90 days - user analytics
	TTLCartWishlist  = 365 * 24 * time.Hour // 1 year - long-term storage
)

// GetCartRedisConfig returns Redis key patterns and TTLs for cart service
func GetCartRedisConfig() (map[string]string, map[string]time.Duration) {
	keyPatterns := map[string]string{
		"user_cart":      KeyUserCart,
		"cart_item":      KeyCartItem,
		"cart_session":   KeyCartSession,
		"cart_temp":      KeyCartTemp,
		"cart_summary":   KeyCartSummary,
		"abandoned_cart": KeyAbandonedCart,
		"cart_history":   KeyCartHistory,
		"cart_wishlist":  KeyCartWishlist,
	}

	ttlConfig := map[string]time.Duration{
		"user_cart":      TTLUserCart,
		"cart_item":      TTLCartItem,
		"cart_session":   TTLCartSession,
		"cart_temp":      TTLCartTemp,
		"cart_summary":   TTLCartSummary,
		"abandoned_cart": TTLAbandonedCart,
		"cart_history":   TTLCartHistory,
		"cart_wishlist":  TTLCartWishlist,
	}

	return keyPatterns, ttlConfig
}
