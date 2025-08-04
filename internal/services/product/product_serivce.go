package product

import (
	"github.com/hthinh24/go-store/services/product/internal/dto/request"
	"github.com/hthinh24/go-store/services/product/internal/dto/response"
	"github.com/hthinh24/go-store/services/product/internal/entity"
)

type ProductService interface {
	GetProduct(id int64) (*entity.Product, error)
	CreateProduct(data *request.CreateProductRequest) (*response.ProductResponse, error)
	DeleteProduct(id int64) error
}
