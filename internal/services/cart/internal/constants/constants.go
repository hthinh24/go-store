package constants

const (
	// Cart status constants
	CartStatusActive    = "ACTIVE"
	CartStatusAbandoned = "ABANDONED"
	CartStatusConverted = "CONVERTED"
)

const (
	// CartItem status constants
	CartItemStatusActive       = "ACTIVE"
	CartItemStatusInActive     = "INACTIVE"
	CartItemStatusOutOfStock   = "OUT_OF_STOCK"
	CartItemStatusDiscontinued = "DISCONTINUED"

	ProductStatusActive = "ACTIVE"
)

const (
	// Default cart expiration time (in hours)
	DefaultCartExpirationHours = 24

	// Maximum items per cart
	MaxItemsPerCart = 100

	// Maximum quantity per item
	MaxQuantityPerItem = 999
)
