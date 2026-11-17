package repositories

import (
	"bluedit/forum-backend/models"
	"database/sql"
	"fmt"
)

type MessageLikeRepository struct {
	db *sql.DB
}

func InitMessageLikeRepositoryRepository(db *sql.DB) *MessageLikeRepository {
	return &MessageLikeRepository{db: db}
}

func (r *MessageLikeRepository) Create(messageLike models.MessageLike) (int, error) {
	query := "INSERT INTO message_likes (fk_id_message, fk_id_user) VALUES (?, ?);"

	sqlResult, sqlErr := r.db.Exec(query,
		messageLike.FkIdMessage,
		messageLike.FkIdUser,
	)
	if sqlErr != nil {
		return -1, fmt.Errorf("Erreur création message_like (repository) - %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf("Erreur création message_like (repository) - %s", idErr)
	}

	return int(id), nil
}

func (r *MessageLikeRepository) ReadAll() ([]models.MessageLike, error) {
	var messageLikesList []models.MessageLike
	sqlResult, sqlErr := r.db.Query("SELECT id_message_like, fk_id_message, fk_id_user FROM message_likes;")
	if sqlErr != nil {
		return messageLikesList, fmt.Errorf("Erreur lecture message_like (repository) - %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var messageLike models.MessageLike
		errScan := sqlResult.Scan(&messageLike.IdMessageLike, &messageLike.FkIdMessage, &messageLike.FkIdUser)
		if errScan != nil {
			return nil, fmt.Errorf("Erreur lecture message_like (repository) - %s", errScan.Error())
		}
		messageLikesList = append(messageLikesList, messageLike)
	}
	return messageLikesList, nil
}

func (r *MessageLikeRepository) ReadById(id int) (models.MessageLike, error) {
	var messageLike models.MessageLike
	sqlErr := r.db.QueryRow("SELECT id_message_like, fk_id_message, fk_id_user FROM message_likes WHERE id_message_like = ?;", id).Scan(&messageLike.IdMessageLike, &messageLike.FkIdMessage, &messageLike.FkIdUser)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.MessageLike{}, fmt.Errorf("Erreur lecture message_like par id (repository) - %s", sqlErr.Error())
		}
		return models.MessageLike{}, fmt.Errorf("Erreur lecture message_like par id (repository) - %s", sqlErr.Error())
	}

	return messageLike, nil
}

func (r *MessageLikeRepository) UpdateById(messageLike models.MessageLike) error {
	query := "UPDATE message_likes SET fk_id_message=?, fk_id_user=? WHERE id_message_like =?;"

	sqlResult, sqlErr := r.db.Exec(query,
		messageLike.FkIdMessage,
		messageLike.FkIdUser,
		messageLike.IdMessageLike,
	)
	if sqlErr != nil {
		return fmt.Errorf("Erreur mise à jour message_like (repository) - %s", sqlErr.Error())
	}

	rowNbr, _ := sqlResult.RowsAffected()
	if rowNbr <= 0 {
		return fmt.Errorf("Erreur mise à jour message_like (repository) - Aucune ligne modifiée")
	}

	return nil
}

func (r *MessageLikeRepository) DeleteById(id int) error {
	sqlResult, sqlErr := r.db.Exec("DELETE FROM message_likes WHERE id_message_like=?;", id)
	if sqlErr != nil {
		return fmt.Errorf("Erreur suppression message_like (repository) - %s", sqlErr.Error())
	}

	rowNbr, _ := sqlResult.RowsAffected()
	if rowNbr <= 0 {
		return fmt.Errorf("Erreur suppression message_like (repository) - Aucune ligne supprimée")
	}

	return nil
}
