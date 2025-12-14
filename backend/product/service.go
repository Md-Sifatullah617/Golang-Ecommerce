package product

import "ecommerce/domain"

type service struct {
	prdcRepo ProductRepo
}

func NewService(prdRepo ProductRepo) Service {
	return &service{
		prdcRepo: prdRepo,
	}
}

func (svc *service) Create(p domain.Product) (*domain.Product, error) {
	return svc.prdcRepo.Create(p)
}

func (svc *service) Get(id int) (*domain.Product, error) {
	return svc.prdcRepo.Get(id)
}

func (svc *service) List(page, limit int64) ([]*domain.Product, error) {
	return svc.prdcRepo.List(page, limit)
}

func (svc *service) Count() (int64, error) {
	return svc.prdcRepo.Count()
}

func (svc *service) Delete(id int) error {
	return svc.prdcRepo.Delete(id)
}

func (svc *service) Update(p domain.Product) (*domain.Product, error) {
	return svc.prdcRepo.Update(p)
}
