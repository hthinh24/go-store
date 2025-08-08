package constants

//'ACTIVE'         -- Available for purchase
//'INACTIVE'       -- Temporarily disabled
//'OUT_OF_STOCK'   -- No inventory available
//'DISCONTINUED'

type ProductStatus string

const (
	ProductStatusActive       ProductStatus = "ACTIVE"       // Available for purchase
	ProductStatusInactive     ProductStatus = "INACTIVE"     // Temporarily disabled
	ProductStatusOutOfStock   ProductStatus = "OUT_OF_STOCK" // No inventory available
	ProductStatusDiscontinued ProductStatus = "DISCONTINUED" // No longer available

	SaleTypePercentage = "PERCENTAGE" // Sale type for percentage discount
	SaleTypeFixed      = "FIXED"      // Sale type for fixed amount discount

	DefaultStock = 0
	DefaultPrice = 0.00
)

func IsValidProductStatus(status string) bool {
	switch ProductStatus(status) {
	case ProductStatusActive, ProductStatusInactive, ProductStatusOutOfStock, ProductStatusDiscontinued:
		return true
	default:
		return false
	}
}
