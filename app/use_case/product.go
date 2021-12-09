package use_case

import "product-api/domain"

type ProductUseCase struct {
	productRepo domain.ProductRepository
}

func NewProductUseCase(productRepo domain.ProductRepository) domain.ProductUseCase {
	return &ProductUseCase{
		productRepo: productRepo,
	}
}

func (uc *ProductUseCase) ReadAll() ([]domain.Product, error) {
	products, err := uc.productRepo.ReadAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (uc *ProductUseCase) ReadAllWithParameter(param string) ([]domain.Product, error) {
	var products []domain.Product
	var err error
	
	if param == "asc(name)" {
		products, err = uc.productRepo.ReadAllOrderByNameASC()
		if err != nil {
			return nil, err
		}
	} else if param == "desc(name)" {
		products, err = uc.productRepo.ReadAllOrderByNameDESC()
		if err != nil {
			return nil, err
		}
	} else if param == "desc(created_at)" {
		products, err = uc.productRepo.ReadAllOrderByDateDESC()
		if err != nil {
			return nil, err
		}
	} else if param == "asc(price)" {
		products, err = uc.productRepo.ReadAllOrderByPriceASC()
		if err != nil {
			return nil, err
		}
	} else if param == "desc(price)" {
		products, err = uc.productRepo.ReadAllOrderByPriceDESC()
		if err != nil {
			return nil, err
		}
	}
	
	return products, nil
}

func (uc *ProductUseCase) ReadByID(id int) (domain.Product, error) {
	product, err := uc.productRepo.ReadByID(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (uc *ProductUseCase) Create(product domain.Product) (domain.Product, error) {
	product, err := uc.productRepo.Create(product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (uc *ProductUseCase) Update(id int, newProduct domain.Product) (domain.Product, error) {
	product, err := uc.productRepo.ReadByID(int(id))
	if err != nil {
		return product, err
	}
	
	product.Name = newProduct.Name
	product.Price = newProduct.Price
	product.Description = newProduct.Description
	product.Quantity = newProduct.Quantity
	
	product, err = uc.productRepo.Update(product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (uc *ProductUseCase) Delete(id int) error {
	err := uc.productRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}