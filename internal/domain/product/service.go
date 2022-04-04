package product

type ProductService struct {
	repository ProductRepository
}

//NewProductService is constructor of CategoryService
func NewProductService(r ProductRepository) *ProductService {
	return &ProductService{
		repository: r,
	}
}
