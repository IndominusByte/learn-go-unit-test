package main

import (
	"errors"
	"fmt"
)

type CategoryEntity struct {
	Id   string
	Name string
}

func (entity CategoryEntity) FindById(id string) *CategoryEntity {
	if id == "1" {
		return &CategoryEntity{
			Id:   "1",
			Name: "Phone",
		}
	} else {
		return nil
	}
}

type CategoryRepository interface {
	FindById(id string) *CategoryEntity
}

type CategoryService struct {
	Repository CategoryRepository
}

func (service CategoryService) Get(id string) (*CategoryEntity, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category not found!")
	} else {
		return category, nil
	}
}

func main() {
	service := CategoryService{Repository: CategoryEntity{}}
	fmt.Println(service.Get("1"))
}
