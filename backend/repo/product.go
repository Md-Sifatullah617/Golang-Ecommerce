package repo

import "fmt"

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}
func (r *productRepo) Get(productID int) (*Product, error) {
	for i := range r.productList {
		if r.productList[i].ID == productID {
			return r.productList[i], nil
		}
	}
	return nil, fmt.Errorf("product with ID %d not found", productID)
}
func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}
func (r *productRepo) Delete(productID int) error {
	for i := range r.productList {
		if r.productList[i].ID == productID {
			r.productList = append(r.productList[:i], r.productList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("product with ID %d not found", productID)
}
func (r *productRepo) Update(p Product) (*Product, error) {
	for i := range r.productList {
		if r.productList[i].ID == p.ID {
			r.productList[i] = &p
			return r.productList[i], nil
		}
	}
	return nil, fmt.Errorf("product with ID %d not found", p.ID)
}

func generateInitialProducts(r *productRepo) {
	allProduct := []*Product{
		{ID: 1, Title: "Wireless Headphones", Description: "High-quality noise-cancelling headphones.", Price: 129.99, ImgURL: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS"},
		{ID: 2, Title: "Smart Watch", Description: "Stylish smart watch with health tracking.", Price: 199.99, ImgURL: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528"},
		{ID: 3, Title: "Running Shoes", Description: "Lightweight shoes for everyday running.", Price: 89.50, ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s"},
		{ID: 4, Title: "Wireless Headphones", Description: "High-quality noise-cancelling headphones.", Price: 129.99, ImgURL: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS"},
		{ID: 5, Title: "Smart Watch", Description: "Stylish smart watch with health tracking.", Price: 199.99, ImgURL: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528"},
		{ID: 6, Title: "Running Shoes", Description: "Lightweight shoes for everyday running.", Price: 89.50, ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s"},
	}

	r.productList = append(r.productList, allProduct...)
}
