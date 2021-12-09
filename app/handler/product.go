package handler

import (
	"net/http"
	"product-api/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUC domain.ProductUseCase
}

func NewProductHandler(e *echo.Echo, productUC domain.ProductUseCase) {
	handler := &ProductHandler{
		productUC: productUC,
	}

	// Sort by latest products
	// Example: localhost:8000/products?sort_by=desc(created_at)

	// Sort by lowest price
	// Example: localhost:8000/products?sort_by=asc(price)

	// Sort by highest price
	// Example: localhost:8000/products?sort_by=desc(price)

	// Sort by name (A-Z)
	// Example: localhost:8000/products?sort_by=asc(name)

	// Sort by name (Z-A)
	// Example: localhost:8000/products?sort_by=desc(name)

	e.GET("/products", handler.GetAllProducts)
	e.GET("/products/:id", handler.GetProductByID)
	e.POST("/products", handler.AddProduct)
	e.PUT("/products/:id", handler.EditProductByID)
	e.DELETE("/products/:id", handler.DeleteProductByID)
}

func (h *ProductHandler) GetAllProducts(c echo.Context) error {
	var products []domain.Product
	var err error

	param := c.QueryParam("sort_by")

	if param == "" {
		products, err = h.productUC.ReadAll()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}
	} else {
		products, err = h.productUC.ReadAllWithParameter(param)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}
	}

	return c.JSON(http.StatusOK, domain.Response{
		Message: "Success",
		Data: products,
	})
}

func (h *ProductHandler) GetProductByID(c echo.Context) error {	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	brand, err := h.productUC.ReadByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, domain.Response{
		Message: "Success",
		Data: brand,
	})
}

func (h *ProductHandler) AddProduct(c echo.Context) error {	
	var product domain.Product
	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	product, err := h.productUC.Create(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusCreated, domain.Response{
		Message: "Success",
		Data: product,
	})
}

func (h *ProductHandler) EditProductByID(c echo.Context) error {	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	var product domain.Product
	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	product, err = h.productUC.Update(id, product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, domain.Response{
		Message: "Success",
		Data: product,
	})
}

func (h *ProductHandler) DeleteProductByID(c echo.Context) error {	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	
	err = h.productUC.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, domain.ResponseDelete{
		Message: "Success",
	})
}