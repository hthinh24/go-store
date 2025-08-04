package product

import (
	"github.com/hthinh24/go-store/services/product/internal/entity"
)

type ProductRepository interface {
	FindProductByID(id int64) (*entity.Product, error)
	CreateProduct(product *entity.Product) error
	DeleteProduct(id int64) error
}
