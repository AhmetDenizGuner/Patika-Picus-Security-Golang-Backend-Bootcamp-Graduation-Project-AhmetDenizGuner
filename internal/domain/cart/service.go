package cart

import "github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/internal/domain/product"

type CartService struct {
	repository     CartRepository
	productService product.ProductService
}

//NewCartService is constructor of CategoryService
func NewCartService(r CartRepository) *CartService {
	return &CartService{
		repository: r,
	}
}

func (service *CartService) addItem(model product.ProductModel) {

	//TODO check user has cart unless create otherway get
	//create cart struct method with error response and manage follow
	//check product exist return error other way add db cart item and update cart price just

}
