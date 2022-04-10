package product

import (
	"errors"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/api/types"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/category"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/csv"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/pagination"
	"gorm.io/gorm"
	"log"
	"strconv"
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

	if newProduct.Price <= 0 || newProduct.StockQuantity <= 0 || newProduct.CategoryID < 1 {
		return ErrProductFieldsMustBePositive
	}

	_, err := service.repository.FindByStockCode(newProduct.StockCode)

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrProductStockCodeMustBeUnique
	}

	_, err1 := service.categoryService.FetchCategoryById(newProduct.CategoryID)

	if err1 != nil {
		return err1
	}

	product := NewProduct(newProduct.Name, newProduct.StockCode, newProduct.StockQuantity, newProduct.Price, newProduct.Description, uint(newProduct.CategoryID))

	err2 := service.repository.Create(product)

	if err2 != nil {
		return err2
	}

	return nil
}

func (service *ProductService) deleteProduct(stockCode string) error {

	product, err := service.repository.FindByStockCode(stockCode)

	if err != nil {
		log.Println(err)
		return err
	}

	//TODO check active orders

	err1 := service.repository.DeleteById(product.ID)

	if err1 != nil {
		log.Println(err1)
		return err1
	}

	return nil
}

func (service *ProductService) updateProduct(model types.AddProductRequest) error {

	if model.Price <= 0 || model.StockQuantity <= 0 || model.CategoryID < 1 {
		return ErrProductFieldsMustBePositive
	}

	product, err := service.repository.FindByStockCode(model.StockCode)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err.Error())
		return err
	}

	product.Update(model.Name, model.StockCode, model.StockQuantity, model.Price, model.Description, uint(model.CategoryID))

	err2 := service.repository.Update(product)

	if err2 != nil {
		return err2
	}

	return nil
}

func (service *ProductService) FetchBySKU(sku string) (Product, error) {

	product, err := service.repository.FindByStockCode(sku)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Product{}, err
	}

	return product, nil

}

func (service *ProductService) UpdateProductQuantityForOrder(itemList []Product, amount []int) error {

	for index, item := range itemList {
		product, err := service.repository.FindByStockCode(item.StockCode)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		err1 := product.UpdateQuantity(amount[index])
		if err1 != nil {
			return err1
		}
	}

	for index, item := range itemList {
		product, _ := service.repository.FindByStockCode(item.StockCode)
		product.UpdateQuantity(amount[index])
		service.repository.Update(product)
	}

	return nil
}

func (service *ProductService) InsertSampleData() {

	tableExist := service.repository.db.Migrator().HasTable(&Product{})

	if !tableExist {

		service.repository.MigrateTable()

		//read category csv
		productSlice, err := csv.ReadCsv("../../resources/products.csv", 1)

		//Name,StockCode,StockQuantity,Price,Description,CategoryID

		if err != nil {
			log.Println(err)
			log.Println("CSV cannot be read!")
			return
		}

		for _, product_line := range productSlice {
			name := product_line[0]
			stockCode := product_line[1]
			stockQuantity, _ := strconv.Atoi(product_line[2])
			price, _ := strconv.ParseFloat(product_line[3], 64)
			description := product_line[4]
			categoryID, _ := strconv.Atoi(product_line[5])

			product := NewProduct(name, stockCode, stockQuantity, price, description, uint(categoryID))
			service.repository.Create(product)
		}
	}
}
