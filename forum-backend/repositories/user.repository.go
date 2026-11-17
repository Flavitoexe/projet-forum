package repositories

import (
	"bluedit/forum-backend/models"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func InitUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user models.User) (int, error) {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?);"

	sqlResult, sqlErr := r.db.Exec(query,
		user.Name,
		user.Email,
		user.Password,
	)
	if sqlErr != nil {
		return -1, fmt.Errorf("Erreur création utilisateur (repository) - %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf("Erreur création utilisateur (repository) - %s", idErr)
	}

	return int(id), nil
}

func (r *UserRepository) ReadAll() ([]models.User, error) {
	var usersList []models.User
	sqlResult, sqlErr := r.db.Query("SELECT id_user, name, email, password, created_at, role, is_banned FROM users;")
	if sqlErr != nil {
		return usersList, fmt.Errorf("Erreur lecture utilisateurs (repository) - %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var user models.User
		errScan := sqlResult.Scan(&user.IdUser, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.Role, &user.IsBanned)
		if errScan != nil {
			return nil, fmt.Errorf("Erreur lecture utilisateurs (repository) - %s", errScan.Error())
		}
		usersList = append(usersList, user)
	}
	return usersList, nil
}

func (r *UserRepository) ReadById(id int) (models.User, error) {
	var user models.User
	sqlErr := r.db.QueryRow("SELECT id_user, name, email, password, created_at, role, is_banned FROM users WHERE id_user = ?;", id).Scan(&user.IdUser, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.Role, &user.IsBanned)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("Erreur lecture utilisateur par id (repository) - %s", sqlErr.Error())
		}
		return models.User{}, fmt.Errorf("Erreur lecture utilisateur par id (repository) - %s", sqlErr.Error())
	}

	return user, nil
}

func (r *UserRepository) UpdateById(user models.User) error {
	query := "UPDATE users SET name=?, email=?, password=?, role=?, is_banned=? WHERE id_user =?;"

	sqlResult, sqlErr := r.db.Exec(query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		user.IsBanned,
		user.IdUser,
	)
	if sqlErr != nil {
		return fmt.Errorf("Erreur mise à jour utilisateur (repository) - %s", sqlErr.Error())
	}

	rowNbr, _ := sqlResult.RowsAffected()
	if rowNbr <= 0 {
		return fmt.Errorf("Erreur mise à jour utilisateur (repository) - Aucune ligne modifiée")
	}

	return nil
}

func (r *UserRepository) DeleteById(id int) error {
	sqlResult, sqlErr := r.db.Exec("DELETE FROM users WHERE id_user=?;", id)
	if sqlErr != nil {
		return fmt.Errorf("Erreur suppression utilisateur (repository) - %s", sqlErr.Error())
	}

	rowNbr, _ := sqlResult.RowsAffected()
	if rowNbr <= 0 {
		return fmt.Errorf("Erreur suppression utilisateur (repository) - Aucune ligne supprimée")
	}

	return nil
}

func (r *UserRepository) FindByNameOrEmail(identifier string) (*models.User, error) {
	var user models.User

	query := "SELECT id_user, name, email, password, role FROM users WHERE name = ? OR email = ?;"
	err := r.db.QueryRow(query, identifier, identifier).Scan(&user.IdUser, &user.Name, &user.Email, &user.Password, &user.Role)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
