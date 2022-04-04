package product

type ProductController struct {
	productService *ProductService
}

func NewProductController(service *ProductService) *ProductController {
	return &ProductController{
		productService: service,
	}
}
