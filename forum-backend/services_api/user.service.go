package services_api

import (
	"bluedit/forum-backend/models"
	repository "bluedit/forum-backend/repositories"
	"fmt"
)

type UserServices struct {
	userRepository *repository.UserRepository
}

func InitUserService(repository *repository.UserRepository) *UserServices {
	return &UserServices{userRepository: repository}
}

func (s *UserServices) Create(user models.User) (int, error) {
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Role == "" {
		return -1, fmt.Errorf("Erreur ajout utilisateur (service) - Données manquantes/invalides")
	}

	userId, userErr := s.userRepository.Create(user)
	if userErr != nil {
		return -1, fmt.Errorf("Erreur ajout utilisateur (service) - %s", userErr.Error())
	}
	return userId, nil
}

func (s *UserServices) ReadAll() ([]models.User, error) {
	usersList, usersErr := s.userRepository.ReadAll()
	if usersErr != nil {
		return nil, fmt.Errorf("Erreur lecture utilisateurs (service) - %s", usersErr.Error())
	}

	return usersList, nil
}

func (s *UserServices) ReadById(id int) (models.User, error) {
	if id <= 0 {
		return models.User{}, fmt.Errorf("Erreur lecture utilisateur par id (service) - Identifiant invalide")
	}

	user, userErr := s.userRepository.ReadById(id)
	if userErr != nil {
		return models.User{}, fmt.Errorf("Erreur lecture utilisateur par id (service) - %s", userErr.Error())
	}

	return user, nil
}

func (s *UserServices) UpdateById(user models.User) error {
	if user.IdUser <= 0 || user.Name == "" || user.Email == "" || user.Password == "" || user.Role == "" {
		return fmt.Errorf("Erreur mise à jour utilisateur (service) - Données manquantes/invalides")
	}

	return s.userRepository.UpdateById(user)
}

func (s *UserServices) DeleteById(id int) error {
	if id <= 0 {
		return fmt.Errorf("Erreur suppression utilisateur (service) - Identifiant invalide")
	}

	return s.userRepository.DeleteById(id)
}
