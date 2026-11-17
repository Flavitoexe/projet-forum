package services_api

import (
	"bluedit/forum-backend/dto"
	"bluedit/forum-backend/models"
	repository "bluedit/forum-backend/repositories"
	"fmt"
)

type MessageServices struct {
	messageRepository *repository.MessageRepository
}

func InitMessageService(repository *repository.MessageRepository) *MessageServices {
	return &MessageServices{messageRepository: repository}
}

func (s *MessageServices) Create(message dto.DtoMessageRequest) (int, error) {
	if message.Content == "" || message.FkIdUser == -1 || message.FkIdThread == -1 {
		return -1, fmt.Errorf("Erreur ajout message (service) - Données manquantes/invalides")
	}

	messageId, messageErr := s.messageRepository.Create(message)
	if messageErr != nil {
		return -1, fmt.Errorf("Erreur ajout message (service) - %s", messageErr.Error())
	}
	return messageId, nil
}

func (s *MessageServices) ReadAll() ([]models.Message, error) {
	messagesList, messagesErr := s.messageRepository.ReadAll()
	if messagesErr != nil {
		return nil, fmt.Errorf("Erreur lecture messages (service) - %s", messagesErr.Error())
	}

	return messagesList, nil
}

func (s *MessageServices) ReadById(id int) (models.Message, error) {
	if id <= 0 {
		return models.Message{}, fmt.Errorf("Erreur lecture message par id (service) - Identifiant invalide")
	}

	message, messageErr := s.messageRepository.ReadById(id)
	if messageErr != nil {
		return models.Message{}, fmt.Errorf("Erreur lecture message par id (service) - %s", messageErr.Error())
	}

	return message, nil
}

func (s *MessageServices) UpdateById(message models.Message) error {
	if message.IdMessage <= 0 || message.Content == "" {
		return fmt.Errorf("Erreur mise à jour message (service) - Données manquantes/invalides")
	}

	return s.messageRepository.UpdateById(message)
}

func (s *MessageServices) DeleteById(id int) error {
	if id <= 0 {
		return fmt.Errorf("Erreur suppression message (service) - Identifiant invalide")
	}

	return s.messageRepository.DeleteById(id)
}

func (s *MessageServices) ReadByThread(id int) ([]dto.DtoMessageResponse, error) {
	if id <= 0 {
		return []dto.DtoMessageResponse{}, fmt.Errorf("Erreur lecture message par thread (service) - Identifiant invalide")
	}

	messages, messagesErr := s.messageRepository.ReadByThread(id)
	if messagesErr != nil {
		return []dto.DtoMessageResponse{}, fmt.Errorf("Erreur lecture message par thread (service) - %s", messagesErr.Error())
	}

	var dtoMessages []dto.DtoMessageResponse
	for _, message := range messages {
		dtoMessage := dto.DtoMessageResponse{
			IdMessage: message.IdMessage,
			Content:   message.Content,
			CreatedAt: message.CreatedAt.Format("02.01.2006 15:04:05"),
			LikeCount: message.LikeCount,
			// Author:    &message.Author,
			// FkIdUser: message.FkIdUser,
		}
		if message.UpdatedAt != nil {
			dtoMessage.UpdatedAt = message.UpdatedAt.Format("02.01.2006 15:04:05")
		}
		if message.FkIdUser != nil {
			dtoMessage.FkIdUser = *message.FkIdUser
		}
		if message.Author != nil {
			dtoMessage.Author = message.Author
		}
		dtoMessages = append(dtoMessages, dtoMessage)

	}

	fmt.Println("messages service : ", dtoMessages)
	return dtoMessages, nil
}
