package types

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignoutRequest struct {
	Email string `json:"email"`
}

type AddProductRequest struct { //Also it used for updated product
	Name          string  `json:"name"`
	StockCode     string  `json:"stock_code"`
	StockQuantity int     `json:"stock_quantity"`
	Price         float64 `json:"price"`
	Description   string  `json:"description"`
	CategoryID    int     `json:"category_id"`
}
