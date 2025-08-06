package product

import (
	"github.com/hthinh24/go-store/services/product/internal/dto/request"
	"github.com/hthinh24/go-store/services/product/internal/dto/response"
)

type ProductService interface {
	GetProductByID(id int64) (*response.ProductResponse, error)
	GetProductDetailByID(id int64) (*response.ProductDetailResponse, error)
	CreateProduct(data *request.CreateProductRequest) (*response.ProductDetailResponse, error)
	DeleteProduct(id int64) error
}
