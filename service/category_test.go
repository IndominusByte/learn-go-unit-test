package service

import (
	"learn-go-unit-test/entity"
	"learn-go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: &categoryRepository}

func TestCategoryServiceNotFound(t *testing.T) {
	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryServiceSuccess(t *testing.T) {
	data := entity.Category{
		Id:   "2",
		Name: "Nyoman",
	}

	// program mock
	categoryRepository.Mock.On("FindById", "2").Return(data)

	category, err := categoryService.Get("2")
	assert.NotNil(t, category)
	assert.Nil(t, err)

	assert.Equal(t, data.Id, category.Id)
	assert.Equal(t, data.Name, category.Name)
}
