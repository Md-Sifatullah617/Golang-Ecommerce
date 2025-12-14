package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/product"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	dbCon *sqlx.DB
}

func NewProductRepo(dbCon *sqlx.DB) product.ProductRepo {
	return &productRepo{
		dbCon: dbCon,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
	INSERT INTO products (
		title,
		description,
		price,
		image_url
	) VALUES (
		$1,
		$2,
		$3,
		$4 
	)
		RETURNING id
	`
	err := r.dbCon.QueryRow(query, p.Title, p.Description, p.Price, p.ImgURL).Scan(&p.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create product %w", err)
	}
	return &p, nil
}

func (r *productRepo) Get(id int) (*domain.Product, error) {
	var prod domain.Product

	query := `
		SELECT 
			id,
			title,
			description,
			price,
			image_url
		FROM products
		WHERE id = $1
		`

	err := r.dbCon.Get(&prod, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with ID %d not found", id)
		}
		return nil, fmt.Errorf("database error %w", err)
	}

	return &prod, nil
}

func (r *productRepo) List(page, limit int64) ([]*domain.Product, error) {
	offset := ((page - 1) * limit) + 1
	var prdList []*domain.Product

	query := `
		SELECT 
			id,
			title,
			description,
			price,
			image_url
		FROM products
		LIMIT $1
		OFFSET $2;
		`

	err := r.dbCon.Select(&prdList, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list product %w", err)
	}

	return prdList, nil
}

func (r *productRepo) Count() (int64, error) {
	var count int
	query := `
		SELECT COUNT (*)
		FROM products
		`

	err := r.dbCon.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to list product %w", err)
	}

	return int64(count), nil
}

func (r *productRepo) Delete(id int) error {
	query := `
		DELETE FROM products WHERE id = $1
	`
	_, err := r.dbCon.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product %w", err)
	}
	return nil
}

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products
		SET 
			title=$1,
			description=$2, 
			price=$3, 
			image_url=$4
		WHERE id = $5
	`
	_, err := r.dbCon.Exec(query, p.Title, p.Description, p.Price, p.ImgURL, p.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update product %w", err)
	}
	return &p, nil
}
