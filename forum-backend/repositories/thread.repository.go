package repositories

import (
	"bluedit/forum-backend/dto"
	"bluedit/forum-backend/models"
	"database/sql"
	"fmt"
)

type ThreadRepository struct {
	db *sql.DB
}

func InitThreadRepository(db *sql.DB) *ThreadRepository {
	return &ThreadRepository{db: db}
}

func (r *ThreadRepository) Create(thread dto.DtoThreadRequest) (int, error) {
	query := "INSERT INTO threads (name, content, fk_id_category, fk_id_user) VALUES (?, ?, ?, ?);"

	sqlResult, sqlErr := r.db.Exec(query,
		thread.Name,
		thread.Content,
		thread.FkIdCategory,
		thread.FkIdUser,
	)
	if sqlErr != nil {
		return -1, fmt.Errorf("Erreur création thread (repository) - %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf("Erreur création thread (repository) - %s", idErr)
	}

	return int(id), nil
}

func (r *ThreadRepository) ReadAll() ([]models.Thread, error) {
	var threadsList []models.Thread
	sqlResult, sqlErr := r.db.Query(`
		SELECT 
			id_thread, 
			threads.name, 
			content, 
			threads.created_at, 
			(select count(*) from messages where messages.fk_id_thread = threads.id_thread) as message_count, 
			(select count(*) from thread_likes where thread_likes.fk_id_thread = threads.id_thread) as like_count,
			fk_id_category, 
			categories.name AS category_name, 
			fk_id_user,
			users.name AS author, 
			status
		FROM threads 
		INNER JOIN categories ON categories.id_category = threads.fk_id_category
		LEFT JOIN users ON users.id_user = threads.fk_id_user;`)

	if sqlErr != nil {
		return threadsList, fmt.Errorf("Erreur lecture threads (repository) - %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var thread models.Thread
		errScan := sqlResult.Scan(&thread.IdThread, &thread.Name, &thread.Content, &thread.CreatedAt, &thread.MessageCount, &thread.LikeCount, &thread.FkIdCategory, &thread.CategoryName, &thread.FkIdUser, &thread.Author, &thread.Status)
		if errScan != nil {
			return nil, fmt.Errorf("Erreur lecture threads (repository) - %s", errScan.Error())
		}
		threadsList = append(threadsList, thread)
	}
	return threadsList, nil
}

func (r *ThreadRepository) ReadById(id int) (models.Thread, error) {
	var thread models.Thread
	sqlErr := r.db.QueryRow(`SELECT 
			id_thread, 
			threads.name, 
			content, 
			threads.created_at, 
			(SELECT COUNT(*) FROM messages WHERE messages.fk_id_thread = threads.id_thread) AS message_count, 
			(SELECT COUNT(*) FROM thread_likes WHERE thread_likes.fk_id_thread = threads.id_thread) AS like_count,
			fk_id_category, 
			categories.name AS category_name, 
			fk_id_user, 
			users.name AS author,
			status
		FROM threads 
		INNER JOIN categories ON categories.id_category = threads.fk_id_category
		LEFT JOIN users ON users.id_user = threads.fk_id_user
		WHERE id_thread = ?;`, id).Scan(&thread.IdThread, &thread.Name, &thread.Content, &thread.CreatedAt, &thread.MessageCount, &thread.LikeCount, &thread.FkIdCategory, &thread.CategoryName, &thread.FkIdUser, &thread.Author, &thread.Status)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.Thread{}, fmt.Errorf("Erreur lecture thread par id (repository) - %s", sqlErr.Error())
		}
		return models.Thread{}, fmt.Errorf("Erreur lecture thread par id (repository) - %s", sqlErr.Error())
	}

	return thread, nil
}

func (r *ThreadRepository) UpdateById(thread models.Thread) error {
	query := "UPDATE threads SET name=?, content=?, fk_id_category=? WHERE id_thread =?;"

	sqlResult, sqlErr := r.db.Exec(query,
		thread.Name,
		thread.Content,
		thread.FkIdCategory,
		thread.IdThread,
	)
	if sqlErr != nil {
		return fmt.Errorf("Erreur mise à jour thread (repository) - %s", sqlErr.Error())
	}

	rowNbr, _ := sqlResult.RowsAffected()
	if rowNbr <= 0 {
		return fmt.Errorf("Erreur mise à jour thread (repository) - Aucune ligne modifiée")
	}

	return nil
}

func (r *ThreadRepository) DeleteById(id int) error {
	sqlResult, sqlErr := r.db.Exec("DELETE FROM threads WHERE id_thread=?;", id)
	if sqlErr != nil {
		return fmt.Errorf("Erreur suppression thread (repository) - %s", sqlErr.Error())
	}

	rowNbr, _ := sqlResult.RowsAffected()
	if rowNbr <= 0 {
		return fmt.Errorf("Erreur suppression thread (repository) - Aucune ligne supprimée")
	}

	return nil
}
