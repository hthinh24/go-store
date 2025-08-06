package product

import (
	"github.com/hthinh24/go-store/services/product/internal/dto/repository"
	"github.com/hthinh24/go-store/services/product/internal/entity"
)

type ProductRepository interface {
	FindProductByID(id int64) (*entity.Product, error)
	FindProductAttributesInfoByProductID(productID int64) (*[]entity.ProductAttributeInfo, error)
	FindProductOptionsInfoByProductID(productID int64) (*[]entity.ProductOptionInfo, error)
	FindProductSKUsByProductID(id int64) (*[]repository.ProductSKUWithInventory, error)

	FindProductAttributesByIDs(productAttributeIDs []int64) (*[]entity.ProductAttribute, error)
	FindProductOptionsByIDs(productOptionIDs []int64) (*[]entity.ProductOption, error)

	CreateProduct(product *entity.Product) error
	CreateProductAttributeInfo(productAttributeInfos *[]entity.ProductAttributeInfo) error
	CreateProductOptionInfo(productOptionInfos *[]entity.ProductOptionInfo) error
	CreateProductAttributeValuesIfNotExist(*[]entity.ProductAttributeValue) error
	CreateProductProductAttributeValues(*[]entity.ProductProductAttributeValue) error
	CreateProductSKUs(productSKUs *[]entity.ProductSKU) error
	CreateProductInventories(inventories *[]entity.ProductInventory) error
	CreateProductOptionCombinations(productOptionCombinations *[]entity.ProductOptionCombination) error
	CreateProductOptionValuesIfNotExist(productOptionValues *[]entity.ProductOptionValue) error

	DeleteProduct(id int64) error
}
