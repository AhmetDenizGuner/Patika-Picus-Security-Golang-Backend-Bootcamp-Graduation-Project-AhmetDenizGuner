package category

import (
	"errors"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/csv"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/pagination"
	"gorm.io/gorm"
	"log"
	"strconv"
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

func (service *CategoryService) AddBulkCategory(fileName string) error {
	//read csv with worker pool
	data, err := csv.ReadCSVWithWorkerPool(fileName)

	if err != nil {
		return err
	}

	//add data that read by worker pool
	for _, category_line := range data {
		//check inout count is correct in line
		if len(category_line) != 2 {
			continue
		}
		name := category_line[0]
		parentID, err := strconv.Atoi(category_line[1])

		//is second argument integer
		if err != nil {
			continue
		}

		//check parent is exist
		_, err2 := service.repository.FindById(parentID)

		if errors.Is(err2, gorm.ErrRecordNotFound) {
			continue
		}

		newCategory := NewCategoryWithParent(name, parentID)

		err3 := service.repository.Create(newCategory)

		if err3 != nil {
			log.Println(err3)
		}

	}

	return nil
}

func (service *CategoryService) FetchCategoryById(id int) (Category, error) {

	category, err := service.repository.FindById(id)

	if err != nil {
		return Category{}, err
	}

	return category, nil

}

func (service *CategoryService) fetchCategoriesWithPagination(page pagination.Pages) ([]CategoryModel, error) {
	var categoryModels []CategoryModel

	categories, err := service.repository.FindByPagination(page.Page, page.PageSize)

	if err != nil {
		return nil, err
	}

	categoryModels = make([]CategoryModel, 0)

	for _, product := range categories {
		categoryModel := NewCategoryModel(product)
		categoryModels = append(categoryModels, *categoryModel)
	}

	return categoryModels, nil

}

func (service *CategoryService) InsertSampleData() {

	tableExist := service.repository.db.Migrator().HasTable(&Category{})

	if !tableExist {

		//TABLE MIGRATION
		service.repository.MigrateTable()

		//ADD BASE CATEGORY
		rootCategory := NewCategoryWithParent("CATEGORIES", 1)
		service.repository.Create(rootCategory)

		//read category csv
		categorySlice, err := csv.ReadCsv("../../resources/category.csv", 1)

		if err != nil {
			log.Println(err)
			log.Println("CSV cannot be read!")
			return
		}

		for _, catg_line := range categorySlice {
			name := catg_line[0]
			parentID, _ := strconv.Atoi(catg_line[1])

			category := NewCategoryWithParent(name, parentID)
			service.repository.Create(category)
		}
	}
}
