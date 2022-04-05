package product

type ProductModel struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	StockCode     string  `json:"stock_code"`
	StockQuantity int     `json:"stock_quantity"`
	Price         float64 `json:"price"`
	Description   string  `json:"description"`
	CategoryID    int     `json:"category_id"`
	CategoryName  string  `json:"category_name"`
}

func NewProductModel(product Product) *ProductModel {
	var productModel ProductModel

	productModel = ProductModel{
		Id:            int(product.ID),
		Name:          product.Name,
		StockCode:     product.StockCode,
		StockQuantity: product.StockQuantity,
		Price:         product.Price,
		Description:   product.Description,
		CategoryID:    int(product.CategoryID),
		CategoryName:  product.Category.Name,
	}

	return &productModel
}
