package constants

//'ACTIVE'         -- Available for purchase
//'INACTIVE'       -- Temporarily disabled
//'OUT_OF_STOCK'   -- No inventory available
//'DISCONTINUED'

type ProductStatus string

const (
	PRODUCT_STATUS_ACTIVE       ProductStatus = "ACTIVE"       // Available for purchase
	PRODUCT_STATUS_INACTIVE     ProductStatus = "INACTIVE"     // Temporarily disabled
	PRODUCT_STATUS_OUT_OF_STOCK ProductStatus = "OUT_OF_STOCK" // No inventory available
	PRODUCT_STATUS_DISCONTINUED ProductStatus = "DISCONTINUED" // No longer available
)

func IsValidProductStatus(status string) bool {
	switch ProductStatus(status) {
	case PRODUCT_STATUS_ACTIVE, PRODUCT_STATUS_INACTIVE, PRODUCT_STATUS_OUT_OF_STOCK, PRODUCT_STATUS_DISCONTINUED:
		return true
	default:
		return false
	}
}
