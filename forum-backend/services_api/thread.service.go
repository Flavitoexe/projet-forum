package services_api

import (
	"bluedit/forum-backend/dto"
	"bluedit/forum-backend/models"
	repository "bluedit/forum-backend/repositories"
	"fmt"
)

type ThreadServices struct {
	threadRepository *repository.ThreadRepository
}

func InitThreadService(repository *repository.ThreadRepository) *ThreadServices {
	return &ThreadServices{threadRepository: repository}
}

func (s *ThreadServices) Create(thread dto.DtoThreadRequest) (int, error) {
	if thread.Name == "" || thread.Content == "" || thread.FkIdCategory == -1 || thread.FkIdUser == -1 {
		return -1, fmt.Errorf("Erreur ajout thread (service) - Données manquantes/invalides")
	}

	threadId, threadErr := s.threadRepository.Create(thread)
	if threadErr != nil {
		return -1, fmt.Errorf("Erreur ajout thread (service) - %s", threadErr.Error())
	}
	return threadId, nil
}

func (s *ThreadServices) ReadAll() ([]models.Thread, error) {
	threadsList, threadsErr := s.threadRepository.ReadAll()
	if threadsErr != nil {
		return nil, fmt.Errorf("Erreur lecture threads (service) - %s", threadsErr.Error())
	}

	return threadsList, nil
}

func (s *ThreadServices) ReadById(id int) (models.Thread, error) {
	if id <= 0 {
		return models.Thread{}, fmt.Errorf("Erreur lecture thread par id (service) - Identifiant invalide")
	}

	thread, threadErr := s.threadRepository.ReadById(id)
	if threadErr != nil {
		return models.Thread{}, fmt.Errorf("Erreur lecture thread par id (service) - %s", threadErr.Error())
	}

	return thread, nil
}

func (s *ThreadServices) UpdateById(thread models.Thread) error {
	if thread.IdThread <= 0 || thread.Name == "" || thread.Content == "" || thread.FkIdCategory == -1 {
		return fmt.Errorf("Erreur mise à jour thread (service) - Données manquantes/invalides")
	}

	return s.threadRepository.UpdateById(thread)
}

func (s *ThreadServices) DeleteById(id int) error {
	if id <= 0 {
		return fmt.Errorf("Erreur suppression thread (service) - Identifiant invalide")
	}

	return s.threadRepository.DeleteById(id)
}
