package category

type CategoryModel struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	ChildCategories []CategoryModel `json:"child_categories"`
}

func NewCategoryModel(category Category) *CategoryModel {
	categoryModel := &CategoryModel{
		ID:              int(category.ID),
		Name:            category.Name,
		ChildCategories: []CategoryModel{},
	}

	for _, subCategory := range category.SubCategories {
		subCategoryModel := NewCategoryModel(subCategory)
		categoryModel.ChildCategories = append(categoryModel.ChildCategories, *subCategoryModel)
	}

	return categoryModel
}
