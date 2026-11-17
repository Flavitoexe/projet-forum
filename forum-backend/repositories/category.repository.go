package repositories

import (
	"bluedit/forum-backend/models"
	"database/sql"
	"fmt"
)

type CategoryRepository struct {
	db *sql.DB
}

func InitCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category models.Category) (int, error) {
	query := "INSERT INTO categories (name, description) VALUES (?, ?);"

	sqlResult, sqlErr := r.db.Exec(query,
		category.Name,
		category.Description,
	)
	if sqlErr != nil {
		return -1, fmt.Errorf("Erreur création catégorie (repository) - %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf("Erreur création catégorie (repository) - %s", idErr)
	}

	return int(id), nil
}

func (r *CategoryRepository) ReadAll() ([]models.Category, error) {
	var categoriesList []models.Category
	sqlResult, sqlErr := r.db.Query("SELECT id_category, name, description FROM categories;")
	if sqlErr != nil {
		return categoriesList, fmt.Errorf("Erreur lecture catégories (repository) - %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var category models.Category
		errScan := sqlResult.Scan(&category.IdCategory, &category.Name, &category.Description)
		if errScan != nil {
			return nil, fmt.Errorf("Erreur lecture catégories (repository) - %s", errScan.Error())
		}
		categoriesList = append(categoriesList, category)
	}
	return categoriesList, nil
}

func (r *CategoryRepository) ReadById(id int) (models.Category, error) {
	var category models.Category
	sqlErr := r.db.QueryRow("SELECT id_category, name, description FROM categories WHERE id_category = ?;", id).Scan(&category.IdCategory, &category.Name, &category.Description)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.Category{}, fmt.Errorf("Erreur lecture catégorie par id (repository) - %s", sqlErr.Error())
		}
		return models.Category{}, fmt.Errorf("Erreur lecture catégorie par id (repository) - %s", sqlErr.Error())
	}

	return category, nil
}

func (r *CategoryRepository) UpdateById(category models.Category) error {
	query := "UPDATE categories SET name=?, description=? WHERE id_category =?;"

	sqlResult, sqlErr := r.db.Exec(query,
		category.Name,
		category.Description,
		category.IdCategory,
	)
	if sqlErr != nil {
		return fmt.Errorf("Erreur mise à jour catégorie (repository) - %s", sqlErr.Error())
	}

	rowNbr, _ := sqlResult.RowsAffected()
	if rowNbr <= 0 {
		return fmt.Errorf("Erreur mise à jour catégorie (repository) - Aucune ligne modifiée")
	}

	return nil
}

func (r *CategoryRepository) DeleteById(id int) error {
	sqlResult, sqlErr := r.db.Exec("DELETE FROM categories WHERE id_category=?;", id)
	if sqlErr != nil {
		return fmt.Errorf("Erreur suppression catégorie (repository) - %s", sqlErr.Error())
	}

	rowNbr, _ := sqlResult.RowsAffected()
	if rowNbr <= 0 {
		return fmt.Errorf("Erreur suppression catégorie (repository) - Aucune ligne supprimée")
	}

	return nil
}
