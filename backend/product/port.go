package product

import (
	"ecommerce/domain"
	prdcHandler "ecommerce/rest/handlers/products"
)

type Service interface {
	prdcHandler.Service
}

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(productID int) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Count() (int64, error)
	Delete(productID int) error
	Update(p domain.Product) (*domain.Product, error)
}
