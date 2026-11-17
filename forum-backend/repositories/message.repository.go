package repositories

import (
	"bluedit/forum-backend/dto"
	"bluedit/forum-backend/models"
	"database/sql"
	"fmt"
)

type MessageRepository struct {
	db *sql.DB
}

func InitMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(message dto.DtoMessageRequest) (int, error) {
	query := "INSERT INTO messages (content, fk_id_user, fk_id_thread) VALUES (?, ?, ?);"

	sqlResult, sqlErr := r.db.Exec(query,
		message.Content,
		message.FkIdUser,
		message.FkIdThread,
	)
	if sqlErr != nil {
		return -1, fmt.Errorf("Erreur création message (repository) - %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf("Erreur création message (repository) - %s", idErr)
	}

	return int(id), nil
}

func (r *MessageRepository) ReadAll() ([]models.Message, error) {
	var messagesList []models.Message
	sqlResult, sqlErr := r.db.Query(`SELECT 
		id_message, 
		content, 
		messages.created_at, 
		updated_at, 
		fk_id_user, 
		users.name AS author, 
		fk_id_thread,
		(SELECT COUNT(*) FROM message_likes WHERE message_likes.fk_id_message = messages.id_message) AS like_count
	FROM messages
	LEFT JOIN users ON users.id_user = messages.fk_id_user;`)
	if sqlErr != nil {
		return messagesList, fmt.Errorf("Erreur lecture messages (repository) - %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var message models.Message
		errScan := sqlResult.Scan(&message.IdMessage, &message.Content, &message.CreatedAt, &message.UpdatedAt, &message.FkIdUser, &message.Author, &message.FkIdThread, &message.LikeCount)
		if errScan != nil {
			return nil, fmt.Errorf("Erreur lecture messages (repository) - %s", errScan.Error())
		}
		messagesList = append(messagesList, message)
	}
	return messagesList, nil
}

func (r *MessageRepository) ReadById(id int) (models.Message, error) {
	var message models.Message
	sqlErr := r.db.QueryRow(`SELECT 
		id_message, 
		content, 
		messages.created_at, 
		updated_at, 
		(SELECT COUNT(*) FROM message_likes WHERE message_likes.fk_id_message = messages.id_message) AS like_count,
		fk_id_user, 
		users.name AS author, 
		fk_id_thread 
	FROM messages
	LEFT JOIN users ON users.id_user = messages.fk_id_user
	WHERE id_message = ?;`, id).Scan(&message.IdMessage, &message.Content, &message.CreatedAt, &message.UpdatedAt, &message.LikeCount, &message.FkIdUser, &message.Author, &message.FkIdThread)
	if sqlErr != nil {
		return models.Message{}, fmt.Errorf("Erreur lecture message par id (repository) - %s", sqlErr.Error())
	}

	return message, nil
}

func (r *MessageRepository) UpdateById(message models.Message) error {
	query := "UPDATE messages SET content=? WHERE id_message =?;"

	sqlResult, sqlErr := r.db.Exec(query,
		message.Content,

		message.IdMessage,
	)
	if sqlErr != nil {
		return fmt.Errorf("Erreur mise à jour message (repository) - %s", sqlErr.Error())
	}

	rowNbr, _ := sqlResult.RowsAffected()
	if rowNbr <= 0 {
		return fmt.Errorf("Erreur mise à jour message (repository) - Aucune ligne modifiée")
	}

	return nil
}

func (r *MessageRepository) DeleteById(id int) error {
	sqlResult, sqlErr := r.db.Exec("DELETE FROM messages WHERE id_message=?;", id)
	if sqlErr != nil {
		return fmt.Errorf("Erreur suppression message (repository) - %s", sqlErr.Error())
	}

	rowNbr, _ := sqlResult.RowsAffected()
	if rowNbr <= 0 {
		return fmt.Errorf("Erreur suppression message (repository) - Aucune ligne supprimée")
	}

	return nil
}

func (r *MessageRepository) ReadByThread(id int) ([]models.Message, error) {
	var messagesList []models.Message
	sqlResult, sqlErr := r.db.Query(`SELECT 
		id_message, 
		content, 
		messages.created_at, 
		updated_at, 
		fk_id_user, 
		users.name AS author, 
		fk_id_thread 
	FROM messages
	LEFT JOIN users ON users.id_user = messages.fk_id_user
	WHERE fk_id_thread=?;`, id)
	if sqlErr != nil {
		return []models.Message{}, fmt.Errorf("Erreur lecture messages par thread (repository) - %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var message models.Message
		var tempDate sql.NullTime
		errScan := sqlResult.Scan(&message.IdMessage, &message.Content, &message.CreatedAt, &tempDate, &message.FkIdUser, &message.Author, &message.FkIdThread)
		if errScan != nil {
			return []models.Message{}, fmt.Errorf("Erreur lecture message par thread (repository) - %s", errScan.Error())
		}
		if tempDate.Valid {
			message.UpdatedAt = &tempDate.Time
		}
		messagesList = append(messagesList, message)
	}

	fmt.Println("messageslist : ", messagesList)
	return messagesList, nil
}
