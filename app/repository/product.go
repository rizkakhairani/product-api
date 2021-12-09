package repository

import (
	"errors"
	"product-api/domain"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &ProductRepository{db: db}
}

func (m *ProductRepository) ReadAll() ([]domain.Product, error) {
	var products []domain.Product

	if err := m.db.Where("deleted_at is NULL").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (m *ProductRepository) ReadAllOrderByNameASC() ([]domain.Product, error) {
	var products []domain.Product

	if err := m.db.Where("deleted_at is NULL").Order("name asc").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (m *ProductRepository) ReadAllOrderByNameDESC() ([]domain.Product, error) {
	var products []domain.Product

	if err := m.db.Where("deleted_at is NULL").Order("name desc").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (m *ProductRepository) ReadAllOrderByDateDESC() ([]domain.Product, error) {
	var products []domain.Product

	if err := m.db.Where("deleted_at is NULL").Order("created_at desc").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (m *ProductRepository) ReadAllOrderByPriceASC() ([]domain.Product, error) {
	var products []domain.Product

	if err := m.db.Where("deleted_at is NULL").Order("price asc").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (m *ProductRepository) ReadAllOrderByPriceDESC() ([]domain.Product, error) {
	var products []domain.Product

	if err := m.db.Where("deleted_at is NULL").Order("price desc").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (m *ProductRepository) Create(product domain.Product) (domain.Product, error) {
	if err := m.db.Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (m *ProductRepository) Update(product domain.Product) (domain.Product, error) {
	if err := m.db.Model(&product).Updates(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (m *ProductRepository) ReadByID(id int) (domain.Product, error) {
	var product domain.Product

	if err := m.db.Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (m *ProductRepository) Delete(id int) error {
	var product domain.Product
	if row := m.db.Where("id = ?", id).Delete(&product).RowsAffected; row < 1 {
		err := errors.New("voucher is not found")
		return err
	}

	return nil
}