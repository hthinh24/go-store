package service

import (
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/services/product"
	"github.com/hthinh24/go-store/services/product/internal/constants"
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

func (p *productService) GetProductSKUByID(skuID int64) (*response.ProductSKUDetailResponse, error) {
	p.logger.Info("Get product SKU with ID: ", skuID)

	productSKUEntity, err := p.productRepository.FindProductSKUByID(skuID)
	if err != nil {
		p.logger.Error("Error retrieving product SKU with ID: ", skuID, "Error: ", err)
		return nil, err
	}

	product, err := p.productRepository.FindProductByID(productSKUEntity.ProductID)
	if err != nil {
		return nil, err
	}

	p.logger.Info("Product SKU retrieved successfully, ID: ", productSKUEntity.ID)
	return p.createProductSKUWithInventoryResponse(product.BasePrice, productSKUEntity), nil
}

func (p *productService) CreateProduct(data *request.CreateProductRequest) (*response.ProductDetailResponse, error) {
	p.logger.Info("Creating product with name", data.Name)

	// Create Product Entity from request data
	productEntity := p.createProductEntity(data)

	// Create transactional repository
	txRepo, err := p.productRepository.WithTransaction()
	if err != nil {
		p.logger.Error("Failed to create transaction:", err)
		return nil, err
	}

	// Ensure rollback on error or panic
	defer func() {
		if r := recover(); r != nil {
			txRepo.Rollback()
			panic(r)
		}
	}()

	// 1. Create & Insert the base product entity
	if err := txRepo.CreateProduct(productEntity); err != nil {
		p.logger.Error("Error creating product:", err)
		txRepo.Rollback()
		return nil, err
	}

	// 2. Create & Insert product attribute info
	if err := p.processCreateProductAttributeInfoWithTx(txRepo, productEntity.ID, data.ProductAttributes); err != nil {
		p.logger.Error("Error creating product attribute info:", err)
		txRepo.Rollback()
		return nil, err
	}

	// 3. Create & Insert product option info
	if err := p.processCreateProductOptionInfoWithTx(txRepo, productEntity.ID, data.OptionValues); err != nil {
		p.logger.Error("Error creating product option info:", err)
		txRepo.Rollback()
		return nil, err
	}

	// 4. Create & Insert product attribute values
	if err := p.processCreateProductAttributesWithTx(txRepo, data.ProductAttributes); err != nil {
		p.logger.Error("Error creating product attributes:", err)
		txRepo.Rollback()
		return nil, err
	}

	// 5. Create & Insert product SKUs
	if err := p.processCreateProductSKUsWithTx(txRepo, productEntity.ID, productEntity.Name, &data.ProductSKUs); err != nil {
		p.logger.Error("Error creating product SKUs:", err)
		txRepo.Rollback()
		return nil, err
	}

	// 6. Create & Insert product option combinations
	if err := p.processCreateProductOptionCombinationsWithTx(txRepo, productEntity.ID, data.OptionValues); err != nil {
		p.logger.Error("Error creating product option combinations:", err)
		txRepo.Rollback()
		return nil, err
	}

	// Commit the transaction
	if err := txRepo.Commit(); err != nil {
		p.logger.Error("Failed to commit transaction:", err)
		return nil, err
	}

	p.logger.Info("Product created successfully, ID: ", productEntity.ID)
	return p.createProductDetailResponse(productEntity), nil
}

func (p *productService) DeleteProduct(id int64) error {
	p.logger.Info("Deleting product with ID: ", id)

	err := p.productRepository.DeleteProduct(id)
	if err != nil {
		return err
	}

	p.logger.Info("Product deleted successfully, ID: ", id)
	return nil
}

func (p *productService) processCreateProductAttributeInfoWithTx(txRepo product.ProductRepository, productID int64, attributeMap map[int64][]string) error {
	var attributeIDs []int64
	var productAttributeInfoEntities []entity.ProductAttributeInfo

	for attributeID, _ := range attributeMap {
		attributeIDs = append(attributeIDs, attributeID)
	}

	// 1. Find product attributes by IDs
	productAttributes, err := txRepo.FindProductAttributesByIDs(attributeIDs)
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
	if err := txRepo.CreateProductAttributeInfo(&productAttributeInfoEntities); err != nil {
		p.logger.Error("Error saving product attribute info to repository, error: ", err)
		return err
	}

	return nil
}

func (p *productService) processCreateProductOptionInfoWithTx(txRepo product.ProductRepository, productID int64, optionMap map[int64][]string) error {
	var productOptionIDs []int64
	var productOptionInfoEntities []entity.ProductOptionInfo

	for productOptionID, _ := range optionMap {
		productOptionIDs = append(productOptionIDs, productOptionID)
	}

	// 1. Find product options by IDs
	productOptions, err := txRepo.FindProductOptionsByIDs(productOptionIDs)
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
	if err := txRepo.CreateProductOptionInfo(&productOptionInfoEntities); err != nil {
		p.logger.Error("Error saving product option info to repository, error: ", err)
		return err
	}

	return nil
}

func (p *productService) processCreateProductAttributesWithTx(txRepo product.ProductRepository, attributeValues map[int64][]string) error {
	// 1. Create product attribute values entities from the attribute values map
	var productAttributeValueEntities []entity.ProductAttributeValue
	for attributeID, values := range attributeValues {
		for _, value := range values {
			productAttributeValueEntity := p.createProductAttributeValueEntity(attributeID, value)
			productAttributeValueEntities = append(productAttributeValueEntities, *productAttributeValueEntity)
		}
	}

	// 2. Save product attribute values to the repository
	err := txRepo.CreateProductAttributeValuesIfNotExist(&productAttributeValueEntities)
	if err != nil {
		return err
	}

	return nil
}

func (p *productService) processCreateProductSKUsWithTx(txRepo product.ProductRepository, productID int64, productName string, productSKUData *[]request.CreateProductSKURequest) error {
	// 1. Create product SKU entities from the product SKU data
	var productSKUEntities []entity.ProductSKU
	for _, sku := range *productSKUData {
		productSKUEntity := p.createProductSKUEntity(productID, productName, &sku)
		productSKUEntities = append(productSKUEntities, *productSKUEntity)
	}

	// 2. Save product SKUs to the repository
	if err := txRepo.CreateProductSKUs(&productSKUEntities); err != nil {
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
	if err := txRepo.CreateProductInventories(&productInventoryEntities); err != nil {
		p.logger.Error("Error saving product inventory to repository, error: ", err)
		return err
	}

	return nil
}

func (p *productService) processCreateProductOptionCombinationsWithTx(txRepo product.ProductRepository, id int64, optionValues map[int64][]string) error {
	var productOptionCombinationEntities []entity.ProductOptionCombination

	// 1. Create product option combination entities from the option values map
	displayOrder := int32(1)
	for option, _ := range optionValues {
		productOptionCombinationEntity := p.createProductOptionCombinationEntity(id, option, int32(displayOrder))
		productOptionCombinationEntities = append(productOptionCombinationEntities, *productOptionCombinationEntity)
		displayOrder++
	}

	// 2. Save product option combinations to the repository
	if err := txRepo.CreateProductOptionCombinations(&productOptionCombinationEntities); err != nil {
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
	if err := txRepo.CreateProductOptionValuesIfNotExist(&productOptionValueEntities); err != nil {
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
		SKU:           data.SKU,
		SKUSignature:  p.generateSKUSignature(productName, data.SKU),
		ExtraPrice:    data.ExtraPrice,
		SaleType:      data.SaleType,
		SaleValue:     data.SaleValue,
		SaleStartDate: data.SaleStartDate,
		SaleEndDate:   data.SaleEndDate,
		Status:        string(constants.ProductStatusActive),
		ProductID:     productID,
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

	var productSKUResponses []*response.ProductSKUDetailResponse
	for _, sku := range *productSKUWithInventories {
		productSKUResponse := p.createProductSKUWithInventoryResponse(product.BasePrice, &sku)
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
	productPrice float64,
	productSKUDetail *repository.ProductSKUDetail,
) *response.ProductSKUDetailResponse {
	productSKUPrice := p.calculateProductSKUPrice(productPrice, productSKUDetail.ExtraPrice)
	return &response.ProductSKUDetailResponse{
		ID:            productSKUDetail.ID,
		SKU:           productSKUDetail.SKU,
		SKUSignature:  productSKUDetail.SKUSignature,
		Price:         productSKUPrice,
		SalePrice:     p.calculateProductSKUSalePrice(productSKUPrice, productSKUDetail.SaleType, productSKUDetail.SaleValue),
		SaleStartDate: productSKUDetail.SaleStartDate,
		SaleEndDate:   productSKUDetail.SaleEndDate,
		Stock:         productSKUDetail.Stock,
		Status:        productSKUDetail.Status,
		ProductID:     productSKUDetail.ProductID,
	}
}

func (p *productService) generateSKUSignature(name string, sku string) string {
	skuSignature := strings.ToLower(name + "-" + sku)
	skuSignature = strings.ReplaceAll(sku, " ", "-")

	return skuSignature
}

func (p *productService) calculateProductSKUPrice(productPrice float64, extraPrice float64) float64 {
	if extraPrice < 0 {
		return productPrice
	}

	finalPrice := productPrice + productPrice*extraPrice
	return finalPrice
}

func (p *productService) calculateProductSKUSalePrice(productSKUPrice float64, saleType *string, saleValue *float64) *float64 {
	if saleType == nil || saleValue == nil {
		return nil
	}

	if *saleType == constants.SaleTypePercentage {
		// Calculate sale price as a percentage discount
		finalPrice := productSKUPrice - (productSKUPrice * *saleValue)
		return &finalPrice
	} else if *saleType == constants.SaleTypeFixed {
		finalPrice := productSKUPrice - *saleValue
		if finalPrice < 0 {
			finalPrice = 0
		}
		return &finalPrice
	}

	return nil
}

// CreateProductWithoutSKU Help to create a product without SKU (for case app only have backend API)
func (p *productService) CreateProductWithoutSKU(data *request.CreateProductWithoutSKURequest) (*response.ProductDetailResponse, error) {
	p.logger.Info("Creating product without SKU with name: ", data.Name)

	// Generate all SKU combinations automatically from option values
	productSKUs, err := p.generateAllSKUCombinations(data.Name, data.OptionValues)
	if err != nil {
		return nil, err
	}

	// Create the full CreateProductRequest with generated SKUs
	createProductRequest := &request.CreateProductRequest{
		Name:              data.Name,
		Description:       data.Description,
		ShortDescription:  data.ShortDescription,
		ImageURL:          data.ImageURL,
		Slug:              data.Slug,
		BasePrice:         data.BasePrice,
		SalePrice:         data.SalePrice,
		IsFeatured:        data.IsFeatured,
		SaleStartDate:     data.SaleStartDate,
		SaleEndDate:       data.SaleEndDate,
		Status:            string(constants.ProductStatusActive),
		BrandID:           data.BrandID,
		CategoryID:        data.CategoryID,
		UserID:            data.UserID,
		ProductAttributes: data.ProductAttributes,
		OptionValues:      data.OptionValues,
		ProductSKUs:       *productSKUs,
	}

	// Call the existing CreateProduct function
	return p.CreateProduct(createProductRequest)
}

// generateAllSKUCombinations generates all possible SKU combinations from option values
func (p *productService) generateAllSKUCombinations(productName string, optionValues map[int64][]string) (*[]request.CreateProductSKURequest, error) {
	// Clean up option values - remove empty options
	cleanedOptions := make(map[int64][]string)
	optionIDs := make([]int64, 0)

	for optionID, values := range optionValues {
		if len(values) > 0 {
			cleanedOptions[optionID] = values
			optionIDs = append(optionIDs, optionID)
		}
	}

	// If no options, create a single default SKU
	if len(cleanedOptions) == 0 {
		defaultSKU := []request.CreateProductSKURequest{
			{
				SKU:        productName + "_default",
				ExtraPrice: 0,
				Stock:      100,
			},
		}
		return &defaultSKU, nil
	}

	// Generate all combinations using cartesian product
	combinations := p.generateCartesianProduct(cleanedOptions, optionIDs)

	// Create SKU requests from combinations
	var productSKUs []request.CreateProductSKURequest
	for _, combination := range combinations {
		sku := p.buildSKUFromCombination(productName, combination, optionIDs)
		price := constants.DefaultPrice
		stock := constants.DefaultStock

		productSKUs = append(productSKUs, request.CreateProductSKURequest{
			SKU:        sku,
			ExtraPrice: price,
			Stock:      int32(stock),
		})
	}

	return &productSKUs, nil
}

// generateCartesianProduct generates all possible combinations of option values
func (p *productService) generateCartesianProduct(optionValues map[int64][]string, optionIDs []int64) []map[int64]string {
	if len(optionIDs) == 0 {
		return []map[int64]string{}
	}

	// Start with the first option
	var result []map[int64]string
	firstOptionID := optionIDs[0]
	firstValues := optionValues[firstOptionID]

	for _, value := range firstValues {
		combination := make(map[int64]string)
		combination[firstOptionID] = value
		result = append(result, combination)
	}

	// Add remaining options one by one
	for i := 1; i < len(optionIDs); i++ {
		optionID := optionIDs[i]
		values := optionValues[optionID]

		var newResult []map[int64]string
		for _, existingCombination := range result {
			for _, value := range values {
				newCombination := make(map[int64]string)
				// Copy existing combination
				for k, v := range existingCombination {
					newCombination[k] = v
				}
				// Add new option value
				newCombination[optionID] = value
				newResult = append(newResult, newCombination)
			}
		}
		result = newResult
	}

	return result
}

// buildSKUFromCombination builds SKU string from option combination
func (p *productService) buildSKUFromCombination(productName string, combination map[int64]string, optionIDs []int64) string {
	skuParts := []string{productName}

	// Add option values in consistent order
	for _, optionID := range optionIDs {
		if value, exists := combination[optionID]; exists {
			// Clean up value for SKU (remove spaces, special chars)
			cleanValue := strings.ReplaceAll(value, " ", "")
			skuParts = append(skuParts, cleanValue)
		}
	}

	return strings.Join(skuParts, "_")
}
