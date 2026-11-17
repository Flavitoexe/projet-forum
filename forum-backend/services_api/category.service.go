package services_api

import (
	"bluedit/forum-backend/models"
	repository "bluedit/forum-backend/repositories"
	"fmt"
)

type CategoryServices struct {
	categoryRepository *repository.CategoryRepository
}

func InitCategoryService(repository *repository.CategoryRepository) *CategoryServices {
	return &CategoryServices{categoryRepository: repository}
}

func (s *CategoryServices) Create(category models.Category) (int, error) {
	if category.Name == "" || category.Description == "" {
		return -1, fmt.Errorf("Erreur ajout catégorie (service) - Données manquantes/invalides")
	}

	categoryId, categoryErr := s.categoryRepository.Create(category)
	if categoryErr != nil {
		return -1, fmt.Errorf("Erreur ajout catégorie (service) - %s", categoryErr.Error())
	}
	return categoryId, nil
}

func (s *CategoryServices) ReadAll() ([]models.Category, error) {
	categoriesList, categoriesErr := s.categoryRepository.ReadAll()
	if categoriesErr != nil {
		return nil, fmt.Errorf("Erreur lecture catégories (service) - %s", categoriesErr.Error())
	}

	return categoriesList, nil
}

func (s *CategoryServices) ReadById(id int) (models.Category, error) {
	if id <= 0 {
		return models.Category{}, fmt.Errorf("Erreur lecture catégorie par id (service) - Identifiant invalide")
	}

	category, categoryErr := s.categoryRepository.ReadById(id)
	if categoryErr != nil {
		return models.Category{}, fmt.Errorf("Erreur lecture catégorie par id (service) - %s", categoryErr.Error())
	}

	return category, nil
}

func (s *CategoryServices) UpdateById(category models.Category) error {
	if category.Name == "" || category.Description == "" {
		return fmt.Errorf("Erreur mise à jour catégorie (service) - Données manquantes/invalides")
	}

	return s.categoryRepository.UpdateById(category)
}

func (s *CategoryServices) DeleteById(id int) error {
	if id <= 0 {
		return fmt.Errorf("Erreur suppression catégorie (service) - Identifiant invalide")
	}

	return s.categoryRepository.DeleteById(id)
}
