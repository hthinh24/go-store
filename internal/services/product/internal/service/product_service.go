package service

import (
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/services/product"
	"github.com/hthinh24/go-store/services/product/internal/dto/repository"
	"github.com/hthinh24/go-store/services/product/internal/dto/request"
	"github.com/hthinh24/go-store/services/product/internal/dto/response"
	"github.com/hthinh24/go-store/services/product/internal/entity"
	"strings"
)

type productService struct {
	logger            logger.Logger
	productRepository product.ProductRepository
}

// NewProductService creates a new instance of ProductService
func NewProductService(logger logger.Logger, productRepository product.ProductRepository) product.ProductService {
	return &productService{
		logger:            logger,
		productRepository: productRepository,
	}
}

func (p *productService) GetProductByID(id int64) (*response.ProductResponse, error) {
	p.logger.Info("Get product with ID: ", id)

	productEntity, err := p.productRepository.FindProductByID(id)
	if err != nil {
		return nil, err
	}

	p.logger.Info("Product retrieved successfully, ID: ", productEntity.ID)
	return p.createProductResponse(productEntity), nil
}

func (p *productService) GetProductDetailByID(id int64) (*response.ProductDetailResponse, error) {
	p.logger.Info("Get product with ID: ", id)

	productEntity, err := p.productRepository.FindProductByID(id)
	if err != nil {
		return nil, err
	}

	p.logger.Info("Product retrieved successfully, ID: ", productEntity.ID)
	return p.createProductDetailResponse(productEntity), nil
}

func (p *productService) CreateProduct(data *request.CreateProductRequest) (*response.ProductDetailResponse, error) {
	p.logger.Info("Creating product with name", data.Name)

	// Create Product Entity from request data
	productEntity := p.createProductEntity(data)

	// 1. Create & Insert the base product entity
	if err := p.processCreateProduct(productEntity); err != nil {
		return nil, err
	}

	// 2. Create & Insert product attribute info
	if err := p.processCreateProductAttributeInfo(productEntity.ID, data.ProductAttributes); err != nil {
		return nil, err
	}

	// 3. Create & Insert product option info
	if err := p.processCreateProductOptionInfo(productEntity.ID, data.OptionValues); err != nil {
		return nil, err
	}

	// 4. Create & Insert product attribute values
	if err := p.processCreateProductAttributes(data.ProductAttributes); err != nil {
		return nil, err
	}

	// 5. Create & Insert product SKUs
	if err := p.processCreateProductSKUs(productEntity.ID, productEntity.Name, &data.ProductSKUs); err != nil {
		return nil, err
	}

	// 6. Create & Insert product option combinations
	if err := p.processCreateProductOptionCombinations(productEntity.ID, data.OptionValues); err != nil {
		return nil, err
	}

	p.logger.Info("Product created successfully, ID: ", productEntity.ID)
	return p.createProductDetailResponse(productEntity), nil
}

func (p *productService) DeleteProduct(id int64) error {
	p.logger.Info("Deleting product with ID: ", id)

	err := p.productRepository.DeleteProduct(id)
	if err != nil {
		p.logger.Error("Error deleting product", err)
		return err
	}

	p.logger.Info("Product deleted successfully, ID: ", id)
	return nil
}

func (p *productService) processCreateProduct(product *entity.Product) error {
	if err := p.productRepository.CreateProduct(product); err != nil {
		p.logger.Error("Error saving product to repository, error: ", err)
		return err
	}

	return nil
}

func (p *productService) processCreateProductAttributeInfo(productID int64, attributeMap map[int64][]string) error {
	var attributeIDs []int64
	var productAttributeInfoEntities []entity.ProductAttributeInfo
	for attributeID, _ := range attributeMap {
		attributeIDs = append(attributeIDs, attributeID)
	}

	// 1. Find product attributes by IDs
	productAttributes, err := p.productRepository.FindProductAttributesByIDs(attributeIDs)
	if err != nil {
		p.logger.Error("Error finding product attributes by IDs, error: ", err)
		return err
	}

	// 2. Create product attribute info entities from the attribute map
	for _, attribute := range *productAttributes {
		if values, ok := attributeMap[attribute.ID]; ok {
			for _, value := range values {
				productAttributeInfoEntity := p.createProductAttributeInfoEntity(productID, attribute.Name, value)
				productAttributeInfoEntities = append(productAttributeInfoEntities, *productAttributeInfoEntity)
			}
		}
	}

	// 3. Save product attribute info entities to the repository
	if err := p.productRepository.CreateProductAttributeInfo(&productAttributeInfoEntities); err != nil {
		p.logger.Error("Error saving product attribute info to repository, error: ", err)
		return err
	}

	return nil
}

func (p *productService) processCreateProductOptionInfo(productID int64, optionMap map[int64][]string) error {
	var productOptionIDs []int64
	var productOptionInfoEntities []entity.ProductOptionInfo

	for productOptionID, _ := range optionMap {
		productOptionIDs = append(productOptionIDs, productOptionID)
	}

	// 1. Find product options by IDs
	productOptions, err := p.productRepository.FindProductOptionsByIDs(productOptionIDs)
	if err != nil {
		p.logger.Error("Error finding product options by IDs, error: ", err)
		return err
	}

	// 2. Create product option info entities from the option map
	for _, option := range *productOptions {
		if values, ok := optionMap[option.ID]; ok {
			for _, value := range values {
				productOptionInfoEntity := p.createProductOptionInfoEntity(productID, option.Name, value)
				productOptionInfoEntities = append(productOptionInfoEntities, *productOptionInfoEntity)
			}
		}
	}

	// 3. Save product option info entities to the repository
	if err := p.productRepository.CreateProductOptionInfo(&productOptionInfoEntities); err != nil {
		p.logger.Error("Error saving product option info to repository, error: ", err)
		return err
	}

	return nil
}

func (p *productService) processCreateProductAttributes(attributeValues map[int64][]string) error {
	// 1. Create product attribute values entities from the attribute values map
	var productAttributeValueEntities []entity.ProductAttributeValue
	for attributeID, values := range attributeValues {
		for _, value := range values {
			productAttributeValueEntity := p.createProductAttributeValueEntity(attributeID, value)
			productAttributeValueEntities = append(productAttributeValueEntities, *productAttributeValueEntity)
		}
	}

	// 2. Save product attribute values to the repository
	err := p.productRepository.CreateProductAttributeValuesIfNotExist(&productAttributeValueEntities)
	if err != nil {
		return err
	}

	return nil
}

func (p *productService) processCreateProductSKUs(productID int64,
	productName string,
	productSKUData *[]request.CreateProductSKURequest) error {
	// 1. Create product SKU entities from the product SKU data
	var productSKUEntities []entity.ProductSKU
	for _, sku := range *productSKUData {
		productSKUEntity := p.createProductSKUEntity(productID, productName, &sku)
		productSKUEntities = append(productSKUEntities, *productSKUEntity)
	}

	// 2. Save product SKUs to the repository
	if err := p.productRepository.CreateProductSKUs(&productSKUEntities); err != nil {
		p.logger.Error("Error saving product SKUs to repository, error: ", err)
		return err
	}

	// 3. Create product inventory entities based on product SKUs and stock data
	var productInventoryEntities []entity.ProductInventory
	for i, productSKUEntity := range productSKUEntities {
		productInventory := p.createProductInventoryEntity(&productSKUEntity, (*productSKUData)[i].Stock)
		productInventoryEntities = append(productInventoryEntities, *productInventory)
	}

	// 4. Save product inventory entities to repository
	if err := p.productRepository.CreateProductInventories(&productInventoryEntities); err != nil {
		p.logger.Error("Error saving product inventory to repository, error: ", err)
		return err
	}

	return nil
}

func (p *productService) processCreateProductOptionCombinations(id int64, optionValues map[int64][]string) error {
	var productOptionCombinationEntities []entity.ProductOptionCombination

	// 1. Create product option combination entities from the option values map
	displayOrder := int32(1)
	for option, _ := range optionValues {
		productOptionCombinationEntity := p.createProductOptionCombinationEntity(id, option, int32(displayOrder))
		productOptionCombinationEntities = append(productOptionCombinationEntities, *productOptionCombinationEntity)
		displayOrder++
	}
	// 2. Save product option combinations to the repository
	if err := p.productRepository.CreateProductOptionCombinations(&productOptionCombinationEntities); err != nil {
		p.logger.Error("Error saving product option combinations to repository, error: ", err)
		return err
	}

	// 3. Create product option value entities from the option values map
	var productOptionValueEntities []entity.ProductOptionValue
	for optionID, values := range optionValues {
		for _, value := range values {
			productOptionValueEntity := p.createProductOptionValueEntity(optionID, value)
			productOptionValueEntities = append(productOptionValueEntities, *productOptionValueEntity)
		}
	}

	// 4. Create product option values in the repository if they do not already exist
	if err := p.productRepository.CreateProductOptionValuesIfNotExist(&productOptionValueEntities); err != nil {
		p.logger.Error("Error saving product option values to repository, error: ", err)
		return err
	}

	return nil
}

func (p *productService) createProductEntity(data *request.CreateProductRequest) *entity.Product {
	return &entity.Product{
		Name:             data.Name,
		Description:      data.Description,
		ShortDescription: data.ShortDescription,
		ImageURL:         data.ImageURL,
		Slug:             data.Slug,
		BasePrice:        data.BasePrice,
		SalePrice:        data.SalePrice,
		IsFeatured:       data.IsFeatured,
		SaleStartDate:    data.SaleStartDate,
		SaleEndDate:      data.SaleEndDate,
		Status:           data.Status,
		BrandID:          data.BrandID,
		CategoryID:       data.CategoryID,
		UserID:           data.UserID,
	}
}

func (p *productService) createProductAttributeInfoEntity(productID int64, attributeName string, attributeValue string) *entity.ProductAttributeInfo {
	return &entity.ProductAttributeInfo{
		AttributeName:  attributeName,
		AttributeValue: attributeValue,
		ProductID:      productID,
	}
}

func (p *productService) createProductOptionInfoEntity(productID int64, optionName string, optionValue string) *entity.ProductOptionInfo {
	return &entity.ProductOptionInfo{
		OptionName:  optionName,
		OptionValue: optionValue,
		ProductID:   productID,
	}
}

func (p *productService) createProductAttributeValueEntity(productAttributeID int64, value string) *entity.ProductAttributeValue {
	return &entity.ProductAttributeValue{
		ProductAttributeID: productAttributeID,
		Value:              value,
	}
}

func (p *productService) createProductSKUEntity(productID int64, productName string, data *request.CreateProductSKURequest) *entity.ProductSKU {
	return &entity.ProductSKU{
		SKU:          data.SKU,
		SKUSignature: p.generateSKUSignature(productName, data.SKU),
		Price:        data.Price,
		ProductID:    productID,
	}
}

func (p *productService) createProductInventoryEntity(productSKU *entity.ProductSKU, stock int32) *entity.ProductInventory {
	return &entity.ProductInventory{
		ProductID:      productSKU.ProductID,
		ProductSKUID:   productSKU.ID,
		AvailableStock: stock,
		ReservedStock:  0,
		DamagedStock:   0,
	}
}

func (p *productService) createProductOptionCombinationEntity(productID int64, productOptionID int64, displayOrder int32) *entity.ProductOptionCombination {
	return &entity.ProductOptionCombination{
		ProductID:       productID,
		ProductOptionID: productOptionID,
		DisplayOrder:    displayOrder,
	}
}

func (p *productService) createProductOptionValueEntity(id int64, value string) *entity.ProductOptionValue {
	return &entity.ProductOptionValue{
		ProductOptionID: id,
		Value:           value,
	}
}

func (p *productService) createProductResponse(product *entity.Product) *response.ProductResponse {
	return &response.ProductResponse{
		ID:               product.ID,
		Name:             product.Name,
		ShortDescription: product.ShortDescription,
		ImageURL:         product.ImageURL,
		BasePrice:        product.BasePrice,
		SalePrice:        product.SalePrice,
		IsFeatured:       product.IsFeatured,
		SaleStartDate:    product.SaleStartDate,
		SaleEndDate:      product.SaleEndDate,
		Status:           product.Status,
		BrandID:          product.BrandID,
		CategoryID:       product.CategoryID,
		UserID:           product.UserID,
	}
}

func (p *productService) createProductDetailResponse(product *entity.Product) *response.ProductDetailResponse {
	// 1. Fetch product attributes and options
	productAttributes, err := p.productRepository.FindProductAttributesInfoByProductID(product.ID)
	if err != nil {
		p.logger.Error("Error fetching product attributes for product ID", product.ID, "Error:", err)
		return nil
	}
	productOptions, err := p.productRepository.FindProductOptionsInfoByProductID(product.ID)
	if err != nil {
		p.logger.Error("Error fetching product options for product ID", product.ID, "Error:", err)
		return nil
	}

	// 3. Fetch product SKUs
	productSKUWithInventories, err := p.productRepository.FindProductSKUsByProductID(product.ID)
	if productSKUWithInventories == nil {
		p.logger.Error("No product SKUs found for product ID", product.ID)
		return nil
	}

	// 2. Create response objects for attributes and options
	var attributeValues []*response.ProductWithAttributeValuesResponse
	var optionValues []*response.ProductWithOptionValuesResponse
	for _, attribute := range *productAttributes {
		attributeValues = append(attributeValues, p.createProductWithAttributeValuesResponse(&attribute))
	}
	for _, option := range *productOptions {
		optionValues = append(optionValues, p.createProductWithOptionValuesResponse(&option))
	}

	var productSKUResponses []*response.ProductSKUWithInventoryResponse
	for _, sku := range *productSKUWithInventories {
		productSKUResponse := p.createProductSKUWithInventoryResponse(&sku)
		productSKUResponses = append(productSKUResponses, productSKUResponse)
	}

	return &response.ProductDetailResponse{
		ID:               product.ID,
		Name:             product.Name,
		Description:      product.Description,
		ShortDescription: product.ShortDescription,
		ImageURL:         product.ImageURL,
		Slug:             product.Slug,
		BasePrice:        product.BasePrice,
		SalePrice:        product.SalePrice,
		IsFeatured:       product.IsFeatured,
		SaleStartDate:    product.SaleStartDate,
		SaleEndDate:      product.SaleEndDate,
		Status:           product.Status,
		BrandID:          product.BrandID,
		CategoryID:       product.CategoryID,
		UserID:           product.UserID,
		Version:          product.Version,
		AttributeValues:  &attributeValues,
		ProductSKUs:      &productSKUResponses,
		OptionValues:     &optionValues,
	}

}

func (p *productService) createProductWithAttributeValuesResponse(attribute *entity.ProductAttributeInfo) *response.ProductWithAttributeValuesResponse {
	return &response.ProductWithAttributeValuesResponse{
		ID:              attribute.ID,
		AttributeName:   attribute.AttributeName,
		AttributeValues: attribute.AttributeValue,
	}
}

func (p *productService) createProductWithOptionValuesResponse(option *entity.ProductOptionInfo) *response.ProductWithOptionValuesResponse {
	return &response.ProductWithOptionValuesResponse{
		ID:           option.ID,
		OptionNames:  option.OptionName,
		OptionValues: option.OptionValue,
	}
}

func (p *productService) createProductSKUWithInventoryResponse(
	productSKUWithInventory *repository.ProductSKUWithInventory,
) *response.ProductSKUWithInventoryResponse {
	return &response.ProductSKUWithInventoryResponse{
		ID:           productSKUWithInventory.ID,
		SKU:          productSKUWithInventory.SKU,
		SKUSignature: productSKUWithInventory.SKUSignature,
		Price:        productSKUWithInventory.Price,
		Stock:        productSKUWithInventory.Stock,
		Status:       productSKUWithInventory.Status,
		ProductID:    productSKUWithInventory.ProductID,
	}
}

func (p *productService) generateSKUSignature(name string, sku string) string {
	skuSignature := strings.ToLower(name + "-" + sku)
	skuSignature = strings.ReplaceAll(sku, " ", "-")

	return skuSignature
}
