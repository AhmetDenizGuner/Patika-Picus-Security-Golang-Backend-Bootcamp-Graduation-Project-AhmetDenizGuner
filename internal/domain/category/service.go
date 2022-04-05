package category

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/csv"
	"log"
)

type CategoryService struct {
	repository CategoryRepository
}

//NewCategoryService is constructor of CategoryService
func NewCategoryService(r CategoryRepository) *CategoryService {
	return &CategoryService{
		repository: r,
	}
}

func (service *CategoryService) FetchCategoryRootForList() (CategoryModel, error) {
	var rootCategoryModel CategoryModel

	categoryRoot, err := service.repository.FindAll()

	if err != nil {
		return CategoryModel{}, err
	}

	rootCategoryModel = *NewCategoryModel(categoryRoot)

	return rootCategoryModel, nil

}

func (service *CategoryService) AddBulkCategory(fileName string) error {

	data, err := csv.ReadCsv(fileName, 0)

	if err != nil {
		return err
	}

	//TODO
	log.Fatal(data)

	return nil
}

func (service *CategoryService) FetchCategoryById(id int) (Category, error) {

	category, err := service.repository.FindById(id)

	if err != nil {
		return Category{}, err
	}

	return category, nil

}
