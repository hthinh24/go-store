package postgres

import (
	"strings"

	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/services/product"
	"github.com/hthinh24/go-store/services/product/internal/dto/repository"
	"github.com/hthinh24/go-store/services/product/internal/entity"
	productErrors "github.com/hthinh24/go-store/services/product/internal/errors"
	"gorm.io/gorm"
)

type productRepository struct {
	logger logger.Logger
	db     *gorm.DB
}

func NewProductRepository(logger logger.Logger, db *gorm.DB) *productRepository {
	return &productRepository{
		logger: logger,
		db:     db,
	}
}

// Transaction methods
func (p *productRepository) WithTransaction() (product.ProductRepository, error) {
	p.logger.Info("Creating transactional repository")

	tx := p.db.Begin()
	if tx.Error != nil {
		p.logger.Error("Failed to begin transaction:", tx.Error)
		return nil, tx.Error
	}

	return &productRepository{
		logger: p.logger,
		db:     tx,
	}, nil
}

func (p *productRepository) Commit() error {
	p.logger.Info("Committing transaction")

	if err := p.db.Commit().Error; err != nil {
		p.logger.Error("Failed to commit transaction:", err)
		return err
	}

	p.logger.Info("Transaction committed successfully")
	return nil
}

func (p *productRepository) Rollback() error {
	p.logger.Info("Rolling back transaction")

	if err := p.db.Rollback().Error; err != nil {
		p.logger.Error("Failed to rollback transaction:", err)
		return err
	}

	p.logger.Info("Transaction rolled back successfully")
	return nil
}

func (p *productRepository) FindProductByID(id int64) (*entity.Product, error) {
	p.logger.Info("Finding product by ID:", id)

	var product entity.Product
	if err := p.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, productErrors.ErrProductNotFound{}
	}

	return &product, nil
}

func (p *productRepository) FindProductAttributesInfoByProductID(productID int64) (*[]entity.ProductAttributeInfo, error) {
	p.logger.Info("Finding product attributes info by product ID:", productID)

	var productAttributesInfo []entity.ProductAttributeInfo
	if err := p.db.Where("product_id = ?", productID).Find(&productAttributesInfo).Error; err != nil {
		p.logger.Error("Failed to find product attributes info by product ID:", productID, "Error:", err)
		return nil, err
	}

	p.logger.Info("Found product attributes info:", productAttributesInfo)
	return &productAttributesInfo, nil
}

func (p *productRepository) FindProductOptionsInfoByProductID(productID int64) (*[]entity.ProductOptionInfo, error) {
	p.logger.Info("Finding product options info by product ID:", productID)

	var productOptionsInfo []entity.ProductOptionInfo
	if err := p.db.Where("product_id = ?", productID).Find(&productOptionsInfo).Error; err != nil {
		p.logger.Error("Failed to find product options info by product ID:", productID, "Error:", err)
		return nil, err
	}

	p.logger.Info("Found product options info:", productOptionsInfo)
	return &productOptionsInfo, nil
}

func (p *productRepository) FindProductSKUsByProductID(id int64) (*[]repository.ProductSKUDetail, error) {
	p.logger.Info("Finding product SKUs by product ID:", id)

	var productSKUsWithInventory []repository.ProductSKUDetail
	if err := p.db.
		Table(entity.ProductSKU{}.TableName()+" AS ps").
		Select("ps.id, ps.sku, ps.sku_signature, ps.extra_price,"+
			"ps.sale_type", "ps.sale_value", "ps.sale_start_date", "ps.sale_end_date").
		Joins("JOIN product_inventory AS pi ON ps.id = pi.product_sku_id").
		Where("ps.product_id = ?", id).
		Find(&productSKUsWithInventory).Error; err != nil {
		p.logger.Error("Failed to find product SKUs by product ID:", id, "Error:", err)
		return nil, err
	}

	return &productSKUsWithInventory, nil
}

func (p *productRepository) FindProductAttributesByIDs(productAttributeIDs []int64) (*[]entity.ProductAttribute, error) {
	p.logger.Info("Finding product attributes by IDs:", productAttributeIDs)

	var productAttributes []entity.ProductAttribute
	if err := p.db.Where("id IN ?", productAttributeIDs).Find(&productAttributes).Error; err != nil {
		p.logger.Error("Failed to find product attributes by IDs:", productAttributeIDs, "Error:", err)
		return nil, err
	}

	p.logger.Info("Found product attributes:", productAttributes)
	return &productAttributes, nil
}

func (p *productRepository) FindProductOptionsByIDs(productOptionIDs []int64) (*[]entity.ProductOption, error) {
	p.logger.Info("Finding product options by IDs:", productOptionIDs)

	var productOptions []entity.ProductOption
	if err := p.db.Where("id IN ?", productOptionIDs).Find(&productOptions).Error; err != nil {
		p.logger.Error("Failed to find product options by IDs:", productOptionIDs, "Error:", err)
		return nil, err
	}

	p.logger.Info("Found product options:", productOptions)
	return &productOptions, nil
}

func (p *productRepository) CreateProduct(product *entity.Product) error {
	p.logger.Info("Creating product:", product)

	if err := p.db.Create(product).Error; err != nil {
		p.logger.Error("Failed to create product:", product, "Error:", err)

		// Check for specific database constraint violations
		errMsg := strings.ToLower(err.Error())

		// Check for duplicate slug constraint
		if strings.Contains(errMsg, "duplicate") && strings.Contains(errMsg, "slug") {
			return productErrors.ErrProductAlreadyExists{Slug: product.Slug}
		}

		// Check for duplicate name constraint
		if strings.Contains(errMsg, "duplicate") && strings.Contains(errMsg, "name") {
			return productErrors.ErrProductAlreadyExists{Name: product.Name}
		}

		// Check for foreign key violations
		if strings.Contains(errMsg, "foreign key") {
			if strings.Contains(errMsg, "category") {
				return productErrors.ErrCategoryNotFound{ID: product.CategoryID}
			}
			if strings.Contains(errMsg, "brand") {
				return productErrors.ErrBrandNotFound{ID: product.BrandID}
			}
			if strings.Contains(errMsg, "user") {
				return productErrors.ErrUserNotFound{ID: product.UserID}
			}
		}

		// Check for check constraint violations
		if strings.Contains(errMsg, "check") || strings.Contains(errMsg, "constraint") {
			if strings.Contains(errMsg, "price") {
				return productErrors.ErrInvalidProductData{Field: "price", Message: "price must be greater than 0"}
			}
			if strings.Contains(errMsg, "status") {
				return productErrors.ErrInvalidProductData{Field: "status", Message: "invalid status value"}
			}
		}

		// Generic database transaction error
		return productErrors.ErrDatabaseTransaction{Operation: "create product"}
	}

	p.logger.Info("Product created successfully:", product.ID)
	return nil
}

func (p *productRepository) CreateProductAttributeInfo(productAttributeInfos *[]entity.ProductAttributeInfo) error {
	p.logger.Info("Creating product attribute infos:", productAttributeInfos)

	if err := p.db.Create(productAttributeInfos).Error; err != nil {
		p.logger.Error("Failed to create product attribute infos:", productAttributeInfos, "Error:", err)

		errMsg := strings.ToLower(err.Error())

		// Check for foreign key violations
		if strings.Contains(errMsg, "foreign key") {
			if strings.Contains(errMsg, "product_attribute") {
				return productErrors.ErrAttributeNotFound{ID: 0} // We'd need to parse which specific ID failed
			}
		}

		return productErrors.ErrDatabaseTransaction{Operation: "create product attribute info"}
	}

	p.logger.Info("Product attribute infos created successfully")
	return nil
}

func (p *productRepository) CreateProductOptionInfo(productOptionInfos *[]entity.ProductOptionInfo) error {
	p.logger.Info("Creating product option infos:", productOptionInfos)

	if err := p.db.Create(productOptionInfos).Error; err != nil {
		p.logger.Error("Failed to create product option infos:", productOptionInfos, "Error:", err)

		errMsg := strings.ToLower(err.Error())

		// Check for foreign key violations
		if strings.Contains(errMsg, "foreign key") {
			if strings.Contains(errMsg, "product_option") {
				return productErrors.ErrOptionNotFound{ID: 0} // We'd need to parse which specific ID failed
			}
		}

		return productErrors.ErrDatabaseTransaction{Operation: "create product option info"}
	}

	p.logger.Info("Product option infos created successfully")
	return nil
}

func (p *productRepository) CreateProductAttributeValuesIfNotExist(productAttributeValues *[]entity.ProductAttributeValue) error {
	p.logger.Info("Creating product attribute values")

	for _, value := range *productAttributeValues {
		if err := p.db.Where("value = ?", value.Value).FirstOrCreate(&value).Error; err != nil {
			p.logger.Error("Failed to create or find product attribute value:", value, "Error:", err)
			return err
		}
	}

	p.logger.Info("Product attribute values created successfully")
	return nil
}

func (p *productRepository) CreateProductProductAttributeValues(i *[]entity.ProductProductAttributeValue) error {
	p.logger.Info("Creating product product attribute values:", i)

	if err := p.db.Create(i).Error; err != nil {
		p.logger.Error("Failed to create product product attribute values:", i, "Error:", err)
		return err
	}

	p.logger.Info("Product product attribute values created successfully")
	return nil
}

func (p *productRepository) CreateProductSKUs(productSKUs *[]entity.ProductSKU) error {
	p.logger.Info("Creating product SKUs:", productSKUs)

	if err := p.db.Create(productSKUs).Error; err != nil {
		p.logger.Error("Failed to create product SKUs:", productSKUs, "Error:", err)

		errMsg := strings.ToLower(err.Error())

		// Check for duplicate SKU constraint
		if strings.Contains(errMsg, "duplicate") && strings.Contains(errMsg, "sku") {
			// Extract SKU from productSKUs if possible
			if len(*productSKUs) > 0 {
				return productErrors.ErrSKUAlreadyExists{SKU: (*productSKUs)[0].SKU}
			}
			return productErrors.ErrSKUAlreadyExists{SKU: "unknown"}
		}

		// Check for invalid price
		if strings.Contains(errMsg, "check") && strings.Contains(errMsg, "price") {
			return productErrors.ErrInvalidSKUData{SKU: "unknown", Message: "price must be greater than or equal to 0"}
		}

		return productErrors.ErrDatabaseTransaction{Operation: "create product SKUs"}
	}

	p.logger.Info("Product SKUs created successfully")
	return nil
}

func (p *productRepository) CreateProductInventories(inventories *[]entity.ProductInventory) error {
	p.logger.Info("Creating product inventories:", inventories)

	if err := p.db.Create(inventories).Error; err != nil {
		p.logger.Error("Failed to create product inventories:", inventories, "Error:", err)
		return err
	}

	p.logger.Info("Product inventories created successfully")
	return nil
}

func (p *productRepository) CreateProductOptionCombinations(productOptionCombinations *[]entity.ProductOptionCombination) error {
	p.logger.Info("Creating product option combinations:", productOptionCombinations)

	if err := p.db.Create(productOptionCombinations).Error; err != nil {
		p.logger.Error("Failed to create product option combinations:", productOptionCombinations, "Error:", err)
		return err
	}

	p.logger.Info("Product option combinations created successfully")
	return nil
}

func (p *productRepository) CreateProductOptionValuesIfNotExist(productOptionValues *[]entity.ProductOptionValue) error {
	p.logger.Info("Creating product option values if not exist:", productOptionValues)

	for _, value := range *productOptionValues {
		if err := p.db.Where("value = ?", value.Value).FirstOrCreate(&value).Error; err != nil {
			p.logger.Error("Failed to create or find product option value:", value, "Error:", err)
			return err
		}
	}

	p.logger.Info("Product option values created or found successfully")
	return nil
}

func (p *productRepository) DeleteProduct(id int64) error {
	p.logger.Info("Deleting product with ID:", id)

	if err := p.db.Where("id = ?", id).Delete(&entity.Product{}).Error; err != nil {
		p.logger.Error("Failed to delete product with ID:", id, "Error:", err)
		return err
	}

	p.logger.Info("Product deleted successfully:", id)
	return nil
}
