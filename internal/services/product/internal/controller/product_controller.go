package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hthinh24/go-store/internal/pkg/logger"
	"github.com/hthinh24/go-store/internal/pkg/rest"
	"github.com/hthinh24/go-store/services/product"
	"github.com/hthinh24/go-store/services/product/internal/dto/request"
)

type ProductController struct {
	logger         logger.Logger
	productService product.ProductService
}

func NewProductController(logger logger.Logger, productService product.ProductService) *ProductController {
	return &ProductController{
		logger:         logger,
		productService: productService,
	}
}

func (pc *ProductController) GetProductByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			pc.logger.Error("Invalid product ID: %v", err)
			response := rest.NewErrorResponse(rest.BadRequestError, "Invalid product ID")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		product, err := pc.productService.GetProductByID(id)
		if err != nil {
			pc.ErrorHandler(c, err, "Failed to get product")
			return
		}

		response := rest.NewAPIResponse(http.StatusOK, "Product retrieved successfully", product)
		c.JSON(http.StatusOK, response)
	}
}

func (pc *ProductController) GetProductDetailByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			pc.logger.Error("Invalid product ID: %v", err)
			response := rest.NewErrorResponse(rest.BadRequestError, "Invalid product ID")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		productDetail, err := pc.productService.GetProductDetailByID(id)
		if err != nil {
			pc.ErrorHandler(c, err, "Failed to get product detail")
			return
		}

		response := rest.NewAPIResponse(http.StatusOK, "Product detail retrieved successfully", productDetail)
		c.JSON(http.StatusOK, response)
	}
}

func (pc *ProductController) GetProductSKUByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			pc.logger.Error("Invalid SKU ID: %v", err)
			response := rest.NewErrorResponse(rest.BadRequestError, "Invalid SKU ID")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		productSKU, err := pc.productService.GetProductSKUByID(id)
		if err != nil {
			pc.ErrorHandler(c, err, "Failed to get product SKU")
			return
		}

		c.JSON(http.StatusOK, productSKU)
	}
}

func (pc *ProductController) CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.CreateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			pc.logger.Error("Invalid request body: %v", err)
			response := rest.NewErrorResponse(rest.BadRequestError, "Invalid request body")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		product, err := pc.productService.CreateProduct(&req)
		if err != nil {
			pc.ErrorHandler(c, err, "Failed to create product")
			return
		}

		response := rest.NewAPIResponse(http.StatusCreated, "Product created successfully", product)
		c.JSON(http.StatusCreated, response)
	}
}

func (pc *ProductController) DeleteProductByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			pc.logger.Error("Invalid product ID: %v", err)
			response := rest.NewErrorResponse(rest.BadRequestError, "Invalid product ID")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		err = pc.productService.DeleteProduct(id)
		if err != nil {
			pc.ErrorHandler(c, err, "Failed to delete product")
			return
		}

		response := rest.NewAPIResponse(http.StatusOK, "Product deleted successfully", nil)
		c.JSON(http.StatusOK, response)
	}
}

func (pc *ProductController) CreateProductWithoutSKU() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.CreateProductWithoutSKURequest
		if err := c.ShouldBindJSON(&req); err != nil {
			pc.logger.Error("Invalid request body: %v", err)
			response := rest.NewErrorResponse(rest.BadRequestError, "Invalid request body")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		product, err := pc.productService.CreateProductWithoutSKU(&req)
		if err != nil {
			pc.ErrorHandler(c, err, "Failed to create product")
			return
		}

		response := rest.NewAPIResponse(http.StatusCreated, "Product created successfully", product)
		c.JSON(http.StatusCreated, response)
	}
}
