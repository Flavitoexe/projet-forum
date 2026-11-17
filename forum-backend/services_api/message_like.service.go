package services_api

import (
	"bluedit/forum-backend/models"
	repository "bluedit/forum-backend/repositories"
	"fmt"
)

type MessageLikeServices struct {
	messageLikeRepository *repository.MessageLikeRepository
}

func InitMessageLikeService(repository *repository.MessageLikeRepository) *MessageLikeServices {
	return &MessageLikeServices{messageLikeRepository: repository}
}

func (s *MessageLikeServices) Create(messageLike models.MessageLike) (int, error) {
	if messageLike.FkIdMessage == 0 || messageLike.FkIdUser == 0 {
		return -1, fmt.Errorf("Erreur ajout message_like (service) - Données manquantes/invalides")
	}

	messageLikeId, messageLikeErr := s.messageLikeRepository.Create(messageLike)
	if messageLikeErr != nil {
		return -1, fmt.Errorf("Erreur ajout message_like (service) - %s", messageLikeErr.Error())
	}
	return messageLikeId, nil
}

func (s *MessageLikeServices) ReadAll() ([]models.MessageLike, error) {
	messageLikesList, messageLikesErr := s.messageLikeRepository.ReadAll()
	if messageLikesErr != nil {
		return nil, fmt.Errorf("Erreur lecture message_likes (service) - %s", messageLikesErr.Error())
	}

	return messageLikesList, nil
}

func (s *MessageLikeServices) ReadById(id int) (models.MessageLike, error) {
	if id <= 0 {
		return models.MessageLike{}, fmt.Errorf("Erreur lecture message_like par id (service) - Identifiant invalide")
	}

	messageLike, messageLikeErr := s.messageLikeRepository.ReadById(id)
	if messageLikeErr != nil {
		return models.MessageLike{}, fmt.Errorf("Erreur lecture message_like par id (service) - %s", messageLikeErr.Error())
	}

	return messageLike, nil
}

func (s *MessageLikeServices) UpdateById(messageLike models.MessageLike) error {
	if messageLike.FkIdMessage == 0 || messageLike.FkIdUser == 0 {
		return fmt.Errorf("Erreur mise à jour message_like (service) - Données manquantes/invalides")
	}

	return s.messageLikeRepository.UpdateById(messageLike)
}

func (s *MessageLikeServices) DeleteById(id int) error {
	if id <= 0 {
		return fmt.Errorf("Erreur suppression message_like (service) - Identifiant invalide")
	}

	return s.messageLikeRepository.DeleteById(id)
}
