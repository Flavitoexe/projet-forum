package services_api

import (
	"bluedit/forum-backend/models"
	repository "bluedit/forum-backend/repositories"
	"fmt"
)

type ThreadLikeServices struct {
	threadLikeRepository *repository.ThreadLikeRepository
}

func InitThreadLikeServices(repository *repository.ThreadLikeRepository) *ThreadLikeServices {
	return &ThreadLikeServices{threadLikeRepository: repository}
}

func (s *ThreadLikeServices) Create(threadLike models.ThreadLike) (int, error) {
	if threadLike.FkIdThread == 0 || threadLike.FkIdUser == 0 {
		return -1, fmt.Errorf("Erreur ajout thread_like (service) - Données manquantes/invalides")
	}

	threadLikeId, threadLikeErr := s.threadLikeRepository.Create(threadLike)
	if threadLikeErr != nil {
		return -1, fmt.Errorf("Erreur ajout thread_like (service) - %s", threadLikeErr.Error())
	}
	return threadLikeId, nil
}

func (s *ThreadLikeServices) ReadAll() ([]models.ThreadLike, error) {
	threadLikesList, threadLikesErr := s.threadLikeRepository.ReadAll()
	if threadLikesErr != nil {
		return nil, fmt.Errorf("Erreur lecture message_likes (service) - %s", threadLikesErr.Error())
	}

	return threadLikesList, nil
}

func (s *ThreadLikeServices) ReadById(id int) (models.ThreadLike, error) {
	if id <= 0 {
		return models.ThreadLike{}, fmt.Errorf("Erreur lecture thread_like par id (service) - Identifiant invalide")
	}

	threadLike, threadLikeErr := s.threadLikeRepository.ReadById(id)
	if threadLikeErr != nil {
		return models.ThreadLike{}, fmt.Errorf("Erreur lecture thread_like par id (service) - %s", threadLikeErr.Error())
	}

	return threadLike, nil
}

func (s *ThreadLikeServices) UpdateById(threadLike models.ThreadLike) error {
	if threadLike.FkIdThread == 0 || threadLike.FkIdUser == 0 {
		return fmt.Errorf("Erreur mise à jour thread_like (service) - Données manquantes/invalides")
	}

	return s.threadLikeRepository.UpdateById(threadLike)
}

func (s *ThreadLikeServices) DeleteById(id int) error {
	if id <= 0 {
		return fmt.Errorf("Erreur suppression thread_like (service) - Identifiant invalide")
	}

	return s.threadLikeRepository.DeleteById(id)
}
