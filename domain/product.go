package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name		string	`json:"name" form:"name"`
	Price		int		`json:"price" form:"price"`
	Description	string	`json:"description" form:"description"`
	Quantity	int		`json:"quantity" form:"quantity"`
}

type ProductRepository interface {
	ReadAll() ([]Product, error)
	ReadAllOrderByNameASC() ([]Product, error)
	ReadAllOrderByNameDESC() ([]Product, error)
	ReadAllOrderByDateDESC() ([]Product, error)
	ReadAllOrderByPriceASC() ([]Product, error)
	ReadAllOrderByPriceDESC() ([]Product, error)
	ReadByID(id int) (Product, error)
	Create(product Product) (Product, error)
	Update(product Product) (Product, error)
	Delete(id int) error
}

type ProductUseCase interface {
	ReadAll() ([]Product, error)
	ReadAllWithParameter(param string) ([]Product, error)
	ReadByID(id int) (Product, error)
	Create(product Product) (Product, error)
	Update(id int, newProduct Product) (Product, error)
	Delete(id int) error
}