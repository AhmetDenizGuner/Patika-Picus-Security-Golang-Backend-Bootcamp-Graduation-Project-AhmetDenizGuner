package product

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/api/types"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/category"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/pagination"
)

type ProductService struct {
	repository      ProductRepository
	categoryService category.CategoryService
}

//NewProductService is constructor of CategoryService
func NewProductService(r ProductRepository, categoryService category.CategoryService) *ProductService {
	return &ProductService{
		repository:      r,
		categoryService: categoryService,
	}
}

func (service *ProductService) fetchProductsWithPagination(page pagination.Pages) ([]ProductModel, error) {
	var productModels []ProductModel

	products, err := service.repository.FindByPagination(page.Page, page.PageSize)

	if err != nil {
		return nil, err
	}

	productModels = make([]ProductModel, 0)

	for _, product := range products {
		productModel := NewProductModel(product)
		productModels = append(productModels, *productModel)
	}

	return productModels, nil

}

func (service *ProductService) searchProductsWithPagination(page pagination.Pages, searchItem string) ([]ProductModel, error) {
	var productModels []ProductModel

	if len(searchItem) < 3 {
		return nil, ErrProductShortSearchKeyword
	}

	products, err := service.repository.FindByPaginationAndKey(page.Page, page.PageSize, searchItem)

	if err != nil {
		return nil, err
	}

	productModels = make([]ProductModel, 0)

	for _, product := range products {
		productModel := NewProductModel(product)
		productModels = append(productModels, *productModel)
	}

	return productModels, nil

}

func (service *ProductService) addNewProduct(newProduct types.AddProductRequest) error {

	if newProduct.Price <= 0 || newProduct.StockQuantity <= 0 {
		return ErrProductFieldsMustBePositive
	}

	_, err := service.repository.FindByStockCode(newProduct.StockCode)

	if err == nil {
		return ErrProductStockCodeMustBeUnique
	}

	_, err1 := service.categoryService.FetchCategoryById(newProduct.CategoryID)

	if err1 == nil {
		return err1
	}

	product := NewProduct(newProduct.Name, newProduct.StockCode, newProduct.StockQuantity, newProduct.Price, newProduct.Description, uint(newProduct.CategoryID))

	err2 := service.repository.Create(*product)

	if err2 == nil {
		return err2
	}

	return nil
}

func (service *ProductService) deleteProduct(id int) error {

	product, err := service.repository.FindByID(id)

	if err != nil {
		return err
	}

	//TODO check active orders

	err1 := service.repository.DeleteById(int(product.ID))

	if err1 != nil {
		return err1
	}

	return nil
}

func (service *ProductService) updateProduct(model ProductModel) error {

	if model.Price <= 0 || model.StockQuantity <= 0 {
		return ErrProductFieldsMustBePositive
	}

	_, err := service.repository.FindByStockCode(model.StockCode)

	if err == nil {
		return ErrProductStockCodeMustBeUnique
	}

	product, err1 := service.repository.FindByID(model.Id)

	if err1 != nil {
		return err1
	}

	product.Update(model.Name, model.StockCode, model.StockQuantity, model.Price, model.Description, uint(model.CategoryID))

	err2 := service.repository.Update(product)

	if err2 != nil {
		return err2
	}

	return nil
}
