package category

type CategoryModel struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	ParentCategoryName string `json:"parent_category_name"`
	ParentCategoryID   int    `json:"parent_category_id"`
}

func NewCategoryModel(category Category) *CategoryModel {
	categoryModel := &CategoryModel{
		ID:                 int(category.ID),
		Name:               category.Name,
		ParentCategoryName: category.Parent.Name,
		ParentCategoryID:   int(category.ParentID),
	}

	return categoryModel
}
