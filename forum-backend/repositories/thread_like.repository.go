package repositories

import (
	"bluedit/forum-backend/models"
	"database/sql"
	"fmt"
)

type ThreadLikeRepository struct {
	db *sql.DB
}

func InitThreadLikeRepositoryRepository(db *sql.DB) *ThreadLikeRepository {
	return &ThreadLikeRepository{db: db}
}

func (r *ThreadLikeRepository) Create(threadLike models.ThreadLike) (int, error) {
	query := "INSERT INTO thread_likes (fk_id_thread, fk_id_user) VALUES (?, ?);"

	sqlResult, sqlErr := r.db.Exec(query,
		threadLike.FkIdThread,
		threadLike.FkIdUser,
	)
	if sqlErr != nil {
		return -1, fmt.Errorf("Erreur création thread_like (repository) - %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf("Erreur création thread_like (repository) - %s", idErr)
	}

	return int(id), nil
}

func (r *ThreadLikeRepository) ReadAll() ([]models.ThreadLike, error) {
	var threadLikesList []models.ThreadLike
	sqlResult, sqlErr := r.db.Query("SELECT id_thread_like, fk_id_thread, fk_id_user FROM thread_likes;")
	if sqlErr != nil {
		return threadLikesList, fmt.Errorf("Erreur lecture thread_like (repository) - %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var threadLike models.ThreadLike
		errScan := sqlResult.Scan(&threadLike.IdThreadLike, &threadLike.FkIdThread, &threadLike.FkIdUser)
		if errScan != nil {
			return nil, fmt.Errorf("Erreur lecture thread_like (repository) - %s", errScan.Error())
		}
		threadLikesList = append(threadLikesList, threadLike)
	}
	return threadLikesList, nil
}

func (r *ThreadLikeRepository) ReadById(id int) (models.ThreadLike, error) {
	var threadLike models.ThreadLike
	sqlErr := r.db.QueryRow("SELECT id_thread_like, fk_id_thread, fk_id_user FROM thread_likes WHERE id_thread_like = ?;", id).Scan(&threadLike.IdThreadLike, &threadLike.FkIdThread, &threadLike.FkIdUser)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.ThreadLike{}, fmt.Errorf("Erreur lecture thread_like par id (repository) - %s", sqlErr.Error())
		}
		return models.ThreadLike{}, fmt.Errorf("Erreur lecture thread_like par id (repository) - %s", sqlErr.Error())
	}

	return threadLike, nil
}

func (r *ThreadLikeRepository) UpdateById(threadLike models.ThreadLike) error {
	query := "UPDATE thread_likes SET fk_id_thread=?, fk_id_user=? WHERE id_thread_like =?;"

	sqlResult, sqlErr := r.db.Exec(query,
		threadLike.FkIdThread,
		threadLike.FkIdUser,
		threadLike.IdThreadLike,
	)
	if sqlErr != nil {
		return fmt.Errorf("Erreur mise à jour thread_like (repository) - %s", sqlErr.Error())
	}

	rowNbr, _ := sqlResult.RowsAffected()
	if rowNbr <= 0 {
		return fmt.Errorf("Erreur mise à jour thread_like (repository) - Aucune ligne modifiée")
	}

	return nil
}

func (r *ThreadLikeRepository) DeleteById(id int) error {
	sqlResult, sqlErr := r.db.Exec("DELETE FROM thread_likes WHERE id_thread_like=?;", id)
	if sqlErr != nil {
		return fmt.Errorf("Erreur suppression thread_like (repository) - %s", sqlErr.Error())
	}

	rowNbr, _ := sqlResult.RowsAffected()
	if rowNbr <= 0 {
		return fmt.Errorf("Erreur suppression thread_like (repository) - Aucune ligne supprimée")
	}

	return nil
}
